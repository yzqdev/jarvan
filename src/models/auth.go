package models

import (
	"jarvan/src/pkg/setting"
	"jarvan/src/pkg/util"
	"strings"
)

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CheckAuth(email, password string) (bool, string) {
	var auth Auth
	password = util.Sha1Encode(strings.Join([]string{setting.AppSetting.JwtSecret, password}, ""))
	db.Select("id, username").Where(Auth{Email: email, Password: password}).First(&auth)
	if auth.ID > 0 {
		return true, auth.Username
	}

	return false, ""
}

func GetUsers(page, pageSize int) {

}
