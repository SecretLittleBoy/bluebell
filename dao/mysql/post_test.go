package mysql

import (
	"bluebell/models"
	"bluebell/settings"
	"testing"
)

func init() {
	settings.Config = &settings.ConfigStruct{
		AppConfig:  &settings.AppConfig{},
		AuthConfig: &settings.AuthConfig{},
		LogConfig:  &settings.LogConfig{},
		MySQLConfig: &settings.MySQLConfig{
			User:     "root",
			Password: "88888888",
			Host:     "127.0.0.1",
			Port:     3306,
			DbName:   "bluebell",
		},
		RedisConfig:     &settings.RedisConfig{},
		SnowflakeConfig: &settings.SnowflakeConfig{},
	}
	if err := Init(); err != nil {//测试前需要先把init()中initUserTable()注释掉
		panic(err)
	}
}

func TestCreatePost(t *testing.T) {
	post := &models.Post{
		AuthorID:    1,
		CommunityID: 1,
		Title:       "test",
		Content:     "just a test",
	}
	err := CreatePost(post)
	if err != nil {
		t.Errorf("CreatePost failed, err:%v\n", err)
	}
	t.Logf("post:%#v\n", post)
}
