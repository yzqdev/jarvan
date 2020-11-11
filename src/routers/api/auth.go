package api

import (
	"encoding/json"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"jarvan/src/models"
	e "jarvan/src/pkg/error"
	"jarvan/src/pkg/logging"
	"jarvan/src/pkg/util"
	"net/http"
	"time"
)

type auth struct {
	Email    string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func Login(c *gin.Context) {
	user := map[string]interface{}{}
	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &user)

	email := user["email"].(string)
	password := user["password"].(string)

	valid := validation.Validation{}
	a := auth{Email: email, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExist, uname := models.CheckAuth(email, password)
		if isExist {
			token, err := util.GenerateToken(email, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				data["username"] = uname
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func Logout(c *gin.Context) {
	//appG := app.Gin{c}
	user := map[string]interface{}{}
	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &user)

	refreshToken := user["token"].(string)
	expireTime, _ := time.ParseDuration("-1m") // 10分钟前

	code := e.SUCCESS

	data := make(map[string]interface{})

	_, err := util.RefreshToken(refreshToken, expireTime);
	if err != nil {
		code = e.ERROR
	}

	//appG.Response(http.StatusOK, code, data)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
