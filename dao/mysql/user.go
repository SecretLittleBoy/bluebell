package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"encoding/hex"
)

const secret = "salt"


func QueryUserByUserName() {

}

func InsertUser(newUser *models.User) (err error) {
	encryptedPassword := encryptPassword(newUser.Password)
	sqlStr := `insert into user(user_id, username, password) values(?, ?, ?)`
	_, err = db.Exec(sqlStr, newUser.UserID, newUser.Username, encryptedPassword)
	return
}

func IsUserExist(username string) (bool, error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int64
	if err := db.Get(&count, sqlStr, username); err != nil {
		return false, err
	}
	return count > 0, nil
}

func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}

func Login(p *models.User) (err error) {
	var user models.User
	sqlStr := `select user_id, username, password from user where username = ?`
	err = db.Get(&user, sqlStr, p.Username)
	if err != nil || encryptPassword(p.Password) != user.Password {
		err = ErrorInvalidPassword
		return
	}
	return
}
