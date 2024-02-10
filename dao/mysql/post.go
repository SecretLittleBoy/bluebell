package mysql

import (
	"bluebell/models"
)

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post (post_id, author_id, community_id, title, content) values (?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, p.ID, p.AuthorID, p.CommunityID, p.Title, p.Content)
	return
}