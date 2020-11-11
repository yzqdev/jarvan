package models

import "github.com/jinzhu/gorm"

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	CategoryId    int    `json:"category_id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	State         int    `json:"state"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

func ExistArticleById(id int) (bool, error) {
	var article Article
	err := db.Select("id").Where("id = ?", id).First(&article).Error
	if err != nil {
		return false, err
	}

	if article.ID > 0 {
		return true, nil
	}

	return false, nil
}

func GetArticleTotal(maps interface{}) (count int) {

	db.Model(&Article{}).Where(maps).Count(&count)

	return
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

func GetArticle(id int) (*Article, error) {
	var article Article
	err := db.Where("id = ? AND deleted_on = ? ", id, 0).First(&article).Related(&article.Tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &article, nil
}

func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Update(data)

	return true
}

func AddArticle(data map[string]interface{}) error {
	article := Article{
		TagID:         data["tag_id"].(int),
		Title:         data["title"].(string),
		Desc:          data["desc"].(string),
		Content:       data["content"].(string),
		CreatedAt:     data["created_at"].(string),
		State:         data["state"].(int),
		CoverImageUrl: data["cover_image_url"].(string),
	}

	if err := db.Create(&article).Error; err != nil {
		return err
	}

	return nil
}

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})

	return true
}

func CleanAllArticle() bool {
	db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Article{})

	return true
}
