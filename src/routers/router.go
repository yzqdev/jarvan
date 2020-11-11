package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"jarvan/src/middleware/jwt"
	"jarvan/src/pkg/setting"
	"jarvan/src/pkg/upload"
	"jarvan/src/routers/api"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunModel)

	r.POST("/api/auth/login", api.Login)
	r.POST("/api/auth/logout", api.Logout)
	// upload
	r.POST("/api/upload", api.UploadImage)
	// swagger docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router := r.Group("/api")
	router.Use(jwt.JWT())
	{
		// statistics
		router.GET("/statistics", api.Count)
		// user
		router.GET("/user/list", api.Users)
		router.POST("/user/add", api.UserAdd)
		router.POST("/user/delete", api.UserDelete)
		// article
		router.GET("/article/detail/:id", api.GetArticle)
		router.GET("/article/list", api.GetArticleList)
		router.POST("/article/save", api.SaveArticle)
		// tag

		// category
		router.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	}

	return r
}
