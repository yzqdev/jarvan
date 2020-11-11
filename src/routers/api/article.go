package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"jarvan/src/pkg/app"
	e "jarvan/src/pkg/error"
	"jarvan/src/service/article"
	"net/http"
	"strconv"
)

// @Summary Get article
// @Produce json
// @Param id param int true "id"
// @Success 200 {string} json "{"code":200,"data":{"id":1,"name":""}]"
// @Router /api/articles/{id} [get]
func GetArticle(c *gin.Context) {
	appG := app.Gin{c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("has to be greater than 0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	articleService := article.Article{ID: id}
	exists, err := articleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	resArticle, err := articleService.Get()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, resArticle)
}

// @Summary Get Article List
// @Produce json
// @Param page_num param int true "page_num"
// @Param page_size param int true "page_size"
// @Success 200 {string} json "{"code":200,"data":[{"id":1,"title":"abc"}]}"
// @Router /api/articles/list [get]
func GetArticleList(c *gin.Context) {
	appG := app.Gin{c}
	pageNum := c.DefaultQuery("page_num", "1")
	pageSize := c.DefaultQuery("page_size", "20")

	articleService := article.Article{}

	page_num, _ := strconv.Atoi(pageNum)
	page_size, _ := strconv.Atoi(pageSize)

	list, err := articleService.GetList(page_num, page_size)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_GET_ARTICLES_FAIL, nil)
	}

	appG.Response(http.StatusOK, e.SUCCESS, list)
}

func SaveArticle(c *gin.Context)  {

}
