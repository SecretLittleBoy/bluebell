package redis

import ()

const (
	KeyPrefix              = "bluebell:"
	KeyPostTimeZSet        = "post:time"   // zset. post and post time
	KeyPostScoreZSet       = "post:score"  // zset. post and post score
	KeyPostVotedZSetPrefix = "post:voted:" // zset. post and user voted. prefix + postID
)
