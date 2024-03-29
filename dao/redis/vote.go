package redis

import (
	"errors"
	"math"
	"time"

	"github.com/go-redis/redis"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432
)

var (
	errVoteTimeExpire = errors.New("投票时间已过")
	errVoteRepeated   = errors.New("不允许重复投票")
)

func VoteForPost(userID, postID string, direction float64) error {
	postTime := rdb.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return errVoteTimeExpire
	}
	oldVote := rdb.ZScore(getRedisKey(KeyPostVotedZSetPrefix+postID), userID).Val()
	if direction == oldVote {
		return errVoteRepeated
	}
	var op float64
	if direction > oldVote {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(oldVote - direction)

	pipeline := rdb.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postID)
	if direction == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZSetPrefix+postID), userID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZSetPrefix+postID), redis.Z{Score: direction, Member: userID})
	}
	_, err := pipeline.Exec()
	return err
}

func GetPostVoteAgreeNums(postIDs []string) (data []int64, err error) {
	// data = make([]int64, 0, len(postIDs))
	// for _, postID := range postIDs {
	// 	key := getRedisKey(KeyPostVotedZSetPrefix + postID)
	// 	v := rdb.ZCount(key, "1", "1").Val()//统计赞成票的数量
	// 	data = append(data, v)
	// }
	pipeline := rdb.Pipeline()
	for _, id := range postIDs {
		key := getRedisKey(KeyPostVotedZSetPrefix + id)
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}
