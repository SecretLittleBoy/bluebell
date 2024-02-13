package redis

import (
	"bluebell/models"
	"time"

	"github.com/go-redis/redis"
)

func CreatePost(postID int64) error {
	pipeline := rdb.TxPipeline()
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{Score: float64(time.Now().Unix()), Member: postID})
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{Score: float64(time.Now().Unix()), Member: postID})
	_, err := pipeline.Exec()
	return err
}

func GetPostListInOrder(p *models.ParamPostList) ([]string, error){
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderByScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	// ZRevRange 按分数从大到小的顺序取出指定数量的元素postID
	return rdb.ZRevRange(key, start, end).Result()
}