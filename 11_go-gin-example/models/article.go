package models

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
func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

//根据某个id获取文章
func GetArticle(id int) (article Article) {
	db.Where("id=?", id).First(&article)
	db.Model(&article).Related(&article.Tag)
	return
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
func ExistArticleById(id int) bool {
	var article Article
	db.Select("id").Where("id =?", id).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}
