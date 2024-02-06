package mysql

import (
	"fmt"
	"bluebell/settings"
	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local",
		settings.Config.MySQLConfig.User,
		settings.Config.MySQLConfig.Password,
		settings.Config.MySQLConfig.Host,
		settings.Config.MySQLConfig.Port,
		settings.Config.MySQLConfig.DbName,
	)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect mysql failed", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(settings.Config.MySQLConfig.MaxOpenConns)
	db.SetMaxIdleConns(settings.Config.MySQLConfig.MaxIdleConns)
	zap.L().Debug("connect mysql success")
	return
}

func Close() {
	_ = db.Close()
}