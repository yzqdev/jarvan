package api

import (
	"github.com/gin-gonic/gin"
	"jarvan/src/models"
	"jarvan/src/pkg/app"
	e "jarvan/src/pkg/error"
	"net/http"
	"strconv"
)

var user models.User

// @Summary Get User List
// @Produce json
// @Param username param string false "username"
// @Param email param string false "email"
// @Success 200 {string} json "{"code":200,"data":{"id":1,"name":"Mantis", "email":"tangchunlinit@gmail.com"}]"
// @Router /api/user/list [get]
func Users(c *gin.Context) {
	appG := app.Gin{c}

	user.UserName = c.Param("username")
	user.Email = c.Param("email")

	result, err := user.Users()

	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_GET_USER_LIST_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, result)
}

// @Summary Add User
// @Produce json
// @Param username param string true "username"
// @Param email param string true "email"
// @Param password param string true "password"
// @Success 200 {string} json "{"code":200,"data":1}"
// @Router /api/user/add [post]
func UserAdd(c *gin.Context) {
	appG := app.Gin{c}

	user.UserName = c.Param("username")
	user.Email = c.Param("email")
	user.Password = c.Param("password")

	id, err := user.Insert()

	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_ADD_USER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, id)
}

// @Summary Delete User
// @Produce json
// @Param uid param int true "uid"
// @Success 200 {string} json "{"code":200, "msg": "操作成功", "data":""}"
// @Router /api/user/delete [post]
func UserDelete(c *gin.Context) {
	appG := app.Gin{c}

	id, err := strconv.ParseInt(c.Param("uid"), 10, 64)

	result, err := user.Destroy(id)

	if err != nil || result.Id == 0 {
		appG.Response(http.StatusOK, e.ERROR_DELETE_USER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
