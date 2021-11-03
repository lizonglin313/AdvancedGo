package models

import (
	"gorm.io/gorm"
	"time"
)

type Article struct {
	Model // 里面已经有 created on modified on

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// ExistArticleByID
// @Desc: 	通过文章 id 判断文章是否存在
// @Param:	id
// @Return:	bool
// @Notice:
func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)

	if article.ID > 0 {
		return true
	}
	return false
}

// GetArticleTotal
// @Desc: 	获取所有文章数量
// @Param:	maps
// @Return:	count
// @Notice:
func GetArticleTotal(maps interface{}) (count int64) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

// GetArticles
// @Desc: 	获取所有文章
// @Param:	pageNum
// @Param:	pageSize
// @Param:	maps
// @Return:	articles
// @Notice: Preload就是一个预加载器，它会执行两条 SQL分别是
//			SELECT * FROM blog_articles
//			SELECT * FROM blog_tag WHERE id IN (1,2,3,4)
//			那么在查询出结构后，gorm内部处理对应的映射逻辑，
//			将其填充到Article的Tag中会特别方便，并且避免了循环查询
func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

// GetArticle
// @Desc: 	通过 id 获取文章
// @Param:	id
// @Return:	article
// @Notice: Article有一个结构体成员是TagID，就是外键
//			gorm会通过类名+ID 的方式去找到这两个类之间的关联关系
// 			先把主结构体查出来，再把嵌套结构体查出来
func GetArticle(id int) (article Article) {
	//user 层级大于 language
	//codes := []string{"zh-CN", "en-US", "ja-JP"}
	//db.Model(&user).Where("code IN ?", codes).Association("Languages").Find(&languages)
	//db.Model(&user).Where("code IN ?", codes).Order("code desc").Association("Languages").Find(&languages)

	db.Where("id = ?", id).First(&article)
	db.Model(&article).Association("Tag").Find(&article.Tag) // 存疑 -> 已解决

	return
}

func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(data)
	return true
}

func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})
	return true
}

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})
	return true
}

func (article *Article) BeforeCreate(db *gorm.DB) error {
	db.Statement.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (article *Article) BeforeUpdate(db *gorm.DB) error {
	db.Statement.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

// 后面再补充
//func (article *Article) BeforeDelete(db *gorm.DB) error {
//	db.Statement.SetColumn("")
//}
