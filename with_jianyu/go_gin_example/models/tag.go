package models

import (
	"gorm.io/gorm"
	"time"
)

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int64) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func ExistTagById(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func AddTag(name string, state int, createBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createBy,
	})
	return true
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}

func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)
	return true
}

// BeforeCreate
// @Desc: 	hook 用于回调函数
//			创建对象前自动生成时间
//			https://gorm.io/zh_CN/docs/hooks.html
// @Rece:	tag
// @Param:	db
// @Return:	error
// @Notice:
func (tag *Tag) BeforeCreate(db *gorm.DB) error {
	db.Statement.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

// BeforeUpdate
// @Desc: 	修改对象时自动生成修改时间
// @Rece:	tag
// @Param:	db
// @Return:	error
// @Notice:
func (tag *Tag) BeforeUpdate(db *gorm.DB) error {
	db.Statement.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}
