package article

import (
	"encoding/json"
	"jarvan/src/models"
	"jarvan/src/pkg/gredis"
	"jarvan/src/pkg/logging"
	cacheService "jarvan/src/service/cache"
)

type Article struct {
	ID            int
	TagID         int
	CategoryId    int
	Title         string
	Desc          string
	Content       string
	CoverImageUrl string
	State         int
	CreatedAt     string
	UpdatedAt     string

	PageNum  int
	PageSize int
}

const CacheTime = 3600

func (a *Article) Add() error {
	article := map[string]interface{}{
		"tag_id":          a.TagID,
		"category_id":     a.CategoryId,
		"title":           a.Title,
		"desc":            a.Desc,
		"content":         a.Content,
		"created_at":      a.CreatedAt,
		"cover_image_url": a.CoverImageUrl,
		"state":           a.State,
	}

	if err := models.AddArticle(article); err != nil {
		return err
	}

	return nil
}

func (a *Article) ExistByID() (bool, error) {
	return models.ExistArticleById(a.ID)
}

func (a *Article) Get() (*models.Article, error) {
	var cacheArticle *models.Article

	cache := cacheService.Article{ID: a.ID}
	key := cache.GetArticleKey()
	logging.Info(key)
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheArticle)
			return cacheArticle, nil
		}
	}

	article, err := models.GetArticle(a.ID)
	if err != nil {
		return nil, err
	}

	gredis.Set(key, article, CacheTime)
	return article, nil
}

func (a *Article) GetList(pageNum, pageSize int) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	return data, nil
}
