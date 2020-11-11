package jwt

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	e "jarvan/src/pkg/error"
	"jarvan/src/pkg/util"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		token := ""
		code = e.SUCCESS
		// TODO 待优化，不同请求方式获取请求参数方式不同
		if c.Request.Method == "GET" {
			token = c.Query("token")
		} else {
			user := map[string]interface{}{}
			body, _ := ioutil.ReadAll(c.Request.Body)
			json.Unmarshal(body, &user)
			token = user["token"].(string)
		}
		if token == "" {
			code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}
		c.Next()
	}
}
