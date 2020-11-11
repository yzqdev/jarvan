package api

import (
	"github.com/gin-gonic/gin"
	"jarvan/src/pkg/app"
	e "jarvan/src/pkg/error"
	"net/http"
)

func Count(c *gin.Context) {
	appG := app.Gin{c}
	data := make(map[string]interface{})
	data["articleCount"] = 0
	data["userCount"] = 0
	data["tagCount"] = 0
	data["categoryCount"] = 0
	appG.Response(http.StatusOK, e.SUCCESS, data)
}
