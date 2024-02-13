package redis

import (
	"bluebell/models"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

func CreatePost(postID, communityID int64) error {
	pipeline := rdb.TxPipeline()
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{Score: float64(time.Now().Unix()), Member: postID})
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{Score: float64(time.Now().Unix()), Member: postID})
	communityKey := getRedisKey(KeyCommunitySetPrefix + strconv.Itoa(int(communityID)))
	pipeline.SAdd(communityKey, postID)
	_, err := pipeline.Exec()
	return err
}

func getIDsFormKey(key string, pageNum, pageSize int64) ([]string, error) {
	start := (pageNum - 1) * pageSize
	end := start + pageSize - 1
	// ZREVRANGE 按分数从大到小的顺序查询指定数量的元素
	return rdb.ZRevRange(key, start, end).Result()
}

func GetPostListInOrder(p *models.ParamPostList) ([]string, error) {
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderByScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	return getIDsFormKey(key, p.Page, p.Size)
}

func GetCommunityPostIDsInOrder(p *models.ParamCommunityPostList) ([]string, error) {
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderByScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}

	// 使用 zinterstore 把分区的帖子set与帖子分数的 zset 生成一个新的zset
	communityKey := getRedisKey(KeyCommunitySetPrefix + strconv.Itoa(int(p.CommunityID)))
	// 利用缓存key减少zinterstore执行的次数
	destinationKey := orderKey + strconv.Itoa(int(p.CommunityID))
	if rdb.Exists(destinationKey).Val() < 1 {
		// 不存在，需要计算
		pipeline := rdb.Pipeline()
		pipeline.ZInterStore(destinationKey, redis.ZStore{
			Aggregate: "MAX",
		}, communityKey, orderKey) // zinterstore 计算
		pipeline.Expire(destinationKey, 60*time.Second) // 设置超时时间
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}
	return getIDsFormKey(destinationKey, p.Page, p.Size)
}
