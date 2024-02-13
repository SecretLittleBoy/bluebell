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
	if direction ==0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZSetPrefix+postID), userID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZSetPrefix+postID), redis.Z{Score: direction, Member: userID})
	}
	_, err := pipeline.Exec()
	return err
}

func CreatePost(postID int64) error {
	pipeline := rdb.TxPipeline()
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{Score: float64(time.Now().Unix()), Member: postID})
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{Score: float64(time.Now().Unix()), Member: postID})
	_, err := pipeline.Exec()
	return err
}