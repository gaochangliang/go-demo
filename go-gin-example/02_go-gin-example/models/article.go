package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

//因为需要被外面访问，所以一般字段都是大写，
//CreatedBy 这种形式才会被json成created_by
type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Content    string `json:"content"`
	Desc       string `json:"desc"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

//获取文章总数
func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

//根据条件获取文章列表
func GetArticles(pageNum int, pageSize int, maps interface{}) ([]*Article, error) {
	var articles []*Article
	err := db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articles, nil
}

//根据某个id获取文章
func GetArticle(id int) (*Article, error) {
	var article Article
	err := db.Where("id = ? AND deleted_on = ?", id, 0).First(&article).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	err = db.Model(&article).Related(&article.Tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &article, nil
}

//新增文章
func AddArticle(data map[string]interface{}) bool {
	//观察到没这里的key都是小写，最好和json后的字段保持一致
	//类型断言用于判断  非接口类型是否实现了某个接口  或者 接口类型是否为某个类型
	db.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Content:   data["content"].(string),
		Desc:      data["desc"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})
	return true
}

//编辑文章
func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ? ", id).Update(data)
	return true
}

//删除文章
func DeleteArticle(id int) bool {
	db.Where("id = ? ", id).Delete(&Article{})
	return true
}

//判断文章id是否存在
//判断name是否存在
func ExistArticleById(id int) (bool, error) {
	var article Article
	err := db.Select("id").Where("id =?", id).First(&article).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if article.ID > 0 {
		return true, nil
	}
	return false, nil
}

/*
这属于gorm的Callbacks，可以将回调方法定义为模型结构的指针，在创建、更新、查询、删除时将被调用，如果任何回调返回错误，gorm 将停止未来操作并回滚所有更改。

gorm所支持的回调方法：

创建：BeforeSave、BeforeCreate、AfterCreate、AfterSave
更新：BeforeSave、BeforeUpdate、AfterUpdate、AfterSave
删除：BeforeDelete、AfterDelete
查询：AfterFind

*/

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
