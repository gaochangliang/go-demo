package article_service

import (
	"02_go-gin-example/gin-blog/models"
	"02_go-gin-example/gin-blog/pkg/gredis"
	"02_go-gin-example/gin-blog/pkg/logging"
	"02_go-gin-example/gin-blog/service/cache_service"
	"encoding/json"
)

type Article struct {
	ID            int
	TagID         int
	Title         string
	Desc          string
	Content       string
	CoverImageUrl string
	State         int
	CreateBy      string
	ModifiedBy    string

	PageNum  int
	PageSize int
}

func (a *Article) ExistByID() (bool, error) {
	return models.ExistArticleById(a.ID)
}

func (a *Article) Count() (int, error) {
	return models.GetArticleTotal(a.getMaps())
}

func (a *Article) Get() (*models.Article, error) {
	var cacheArticle *models.Article
	cache := cache_service.Article{ID: a.ID}
	key := cache.GetArticleKey()

	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheArticle)
		}
	}
	article, err := models.GetArticle(a.ID)
	if err != nil {
		return nil, err
	}
	_ = gredis.Set(key, article, 3600)
	return article, nil
}

func (a *Article) GetAll() ([]*models.Article, error) {
	var (
		articles, cacheArticles []*models.Article
	)

	cache := cache_service.Article{
		TagID:    a.TagID,
		State:    a.State,
		PageNum:  0,
		PageSize: 0,
	}

	key := cache.GetArticlesKey()

	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheArticles)
			return cacheArticles, nil
		}
	}

	articles, err := models.GetArticles(a.PageSize, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}

	gredis.Set(key, articles, 3600)
	return articles, nil
}

/*ID            int
TagID         int
Title         string
Desc          string
Content       string
CoverImageUrl string
State         int
CreateBy      string


*/

func (a *Article) Add() error {
	article := map[string]interface{}{
		"tag_id":          a.TagID,
		"title":           a.Title,
		"desc":            a.Desc,
		"content":         a.Content,
		"cover_image_url": a.CoverImageUrl,
		"state":           a.State,
		"create_by":       a.CreateBy,
	}

	if err := models.AddArticle(article); err != nil {
		return err
	}
	return nil
}

func (a *Article) Edit() error {
	//注意这个地方map的使用方式
	return models.EditArticle(a.ID, map[string]interface{}{
		"tag_id":          a.TagID,
		"title":           a.Title,
		"desc":            a.Desc,
		"content":         a.Content,
		"cover_image_url": a.CoverImageUrl,
		"state":           a.State,
		"modified_by":     a.ModifiedBy,
	})
}

func (a *Article) Delete() error {
	//注意这个地方map的使用方式
	return models.DeleteArticle(a.ID)
}

func (a *Article) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["delete_on"] = 0

	if a.State != -1 {
		maps["state"] = a.State
	}

	if a.TagID != -1 {
		maps["tag_id"] = a.TagID
	}

	return maps
}
