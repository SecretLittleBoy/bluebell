package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	id := snowflake.GenID()
	p.ID = id
	return mysql.CreatePost(p)
}
