package logic

import (
	"bluebell/dao/redis"
	"bluebell/models"
	"strconv"
)

//推荐阅读：
//基于用户投票的排序算法 http://www.ruanyifeng.com/blog/algorithm

func VoteForPost(userID int64, p *models.ParamVoteData) error {
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
