package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

func GetCommunityList() (data []*models.Community, err error) {
	data, err = mysql.GetCommunityList()
	return
}

func GetCommunityDetail(id int64) (data *models.CommunityDetail, err error) {
	data, err = mysql.GetCommunityDetailByID(id)
	return
}
