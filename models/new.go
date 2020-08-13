package models

import "go_server/pkg/util"

type New struct {
	ID          int    `gorm:"primary_key" json:"id"`
	Uid         int    `gorm:"uid" json:"uid"`
	Title       string `gorm:"title" json:"title"`
	Author      string `gorm:"author" json:"author"`
	Source      string `gorm:"source" json:"source"`
	Content     string `gorm:"content" json:"content"`
	CreateTime  int64  `gorm: "create_time" json:"create_time"`
	UpdateTime  int64  `gorm: "update_time" json:"update_time"`
	ReleaseTime int64  `gorm: "release_time" json:"release_time"`
	Status      int    `gorm: "status" json:"status"`
	Browse      int    `gorm: "browse" json:"browse"`
}

func GetNewAndCount(pageNum int, pageSize int, title string) (news []New, count int) {
	db.Table("news").Where("title LIKE ?", "%"+title+"%").
		Offset(pageNum).Limit(pageSize).Find(&news)
	db.Table("news").Where("title LIKE ?", "%"+title+"%").
		Count(&count)
	return
}

func EditNewStatus(id int, status int) error {
	var err error
	if status == 1 {
		err = db.Table("news").Model(&New{}).Where("id = ?", id).
			Update(map[string]interface{}{"status": status, "release_time": util.GenerateCurrentTimeStamp()}).Error
	} else {
		err = db.Table("news").Model(&New{}).Where("id = ?", id).
			Update(map[string]interface{}{"status": status}).Error
	}
	return err
}

func DelNew(id int) error {
	err := db.Table("news").Where("id = ?", id).Delete(New{}).Error
	return err
}

func GetNew(id int) (new New) {
	db.Table("news").Where("id = ?", id).Find(&new)
	return
}

func EditNew(id int, data interface{}) error {
	err := db.Table("news").Model(&New{}).Where("id = ?", id).Update(data).Error
	return err
}

func AddNew(data map[string]interface{}) bool {
	db.Table("news").Create(&New{
		Uid: data["uid"].(int),
		Title: data["title"].(string),
		Source: data["source"].(string),
		Content: data["content"].(string),
		Author: data["author"].(string),
		CreateTime: data["create_time"].(int64),
		UpdateTime: data["update_time"].(int64),
		ReleaseTime: data["release_time"].(int64),
		Status: data["status"].(int),
	})
	return true
}