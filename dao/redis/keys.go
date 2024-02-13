package redis

import ()

const (
	KeyPrefix              = "bluebell:"
	KeyPostTimeZSet        = "post:time"   // zset. post and post time
	KeyPostScoreZSet       = "post:score"  // zset. post and post score
	KeyPostVotedZSetPrefix = "post:voted:" // zset. user and user vote. usage: prefix + postID

	KeyCommunitySetPrefix = "community:" // set. community and post
)

func getRedisKey(key string) string {
	return KeyPrefix + key
}
