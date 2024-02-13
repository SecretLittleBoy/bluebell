package logic

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/models"
	"bluebell/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	id := snowflake.GenID()
	p.ID = id
	err =  mysql.CreatePost(p)
	if err!=nil{
		return
	}
	return redis.CreatePost(p.ID)
}

func GetPostById(postId int64) (postDetail *models.ApiPostDetail, err error) {
	post, err := mysql.GetPostById(postId)
	if err != nil {
		return
	}

	author, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		return
	}
	communityDetail, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		return
	}
	postDetail = &models.ApiPostDetail{
		AuthorName:      author.Username,
		Post:            post,
		CommunityDetail: communityDetail,
	}
	return
}

func GetPostList(pageNum int64, pageSize int64) (data []*models.ApiPostDetail, err error) {

	posts, err := mysql.GetPostList(pageNum, pageSize)
	if err != nil {
		return
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))
	for _, post := range posts {
		var author *models.User
		author, err = mysql.GetUserByID(post.AuthorID)
		if err != nil {
			return
		}
		var communityDetail *models.CommunityDetail
		communityDetail, err = mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			return
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      author.Username,
			Post:            post,
			CommunityDetail: communityDetail,
		}
		data = append(data, postDetail)
	}
	return
}
