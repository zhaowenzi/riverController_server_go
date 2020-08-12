package models

import "go_server/pkg/util"

type Announcement struct {
	ID          int    `gorm:"primary_key" json:"id"`
	Uid         int    `gorm:"uid" json:"uid"`
	Title       string `gorm:"title" json:"title"`
	Source      string `gorm:"source" json:"source"`
	Content     string `gorm:"content" json:"content"`
	CreateTime  int64  `gorm: "create_time" json:"create_time"`
	UpdateTime  int64  `gorm: "update_time" json:"update_time"`
	ReleaseTime int64  `gorm: "release_time" json:"release_time"`
	Status      int    `gorm: "status" json:"status"`
	Browse      int    `gorm: "browse" json:"browse"`
}

func GetAnnouncementAndCount(pageNum int, pageSize int, title string) (announcements []Announcement, count int) {
	db.Table("announcement").Where("title LIKE ?", "%"+title+"%").
		Offset(pageNum).Limit(pageSize).Find(&announcements)
	db.Table("announcement").Where("title LIKE ?", "%"+title+"%").
		Count(&count)
	return
}

func EditAnnouncementStatus(id int, status int) error {
	var err error
	if status == 1 {
		err = db.Table("announcement").Model(&Announcement{}).Where("id = ?", id).
			Update(map[string]interface{}{"status": status, "release_time": util.GenerateCurrentTimeStamp()}).Error
	} else {
		err = db.Table("announcement").Model(&Announcement{}).Where("id = ?", id).
			Update(map[string]interface{}{"status": status}).Error
	}

	return err
}

func DelAnnouncement(id int) error {
	err := db.Table("announcement").Where("id = ?", id).Delete(Announcement{}).Error
	return err
}

func GetAnnouncement(id int) (announcement Announcement) {
	db.Table("announcement").Where("id = ?", id).Find(&announcement)
	return
}

func EditAnnouncement(id int, data interface{}) error {
	err := db.Table("announcement").Model(&Announcement{}).Where("id = ?", id).Update(data).Error
	return err
}

func AddAnnouncement(data map[string]interface{}) bool {
	db.Table("announcement").Create(&Announcement{
		Uid: data["uid"].(int),
		Title: data["title"].(string),
		Source: data["source"].(string),
		Content: data["content"].(string),
		CreateTime: data["create_time"].(int64),
		UpdateTime: data["update_time"].(int64),
		ReleaseTime: data["release_time"].(int64),
		Status: data["status"].(int),
	})
	return true
}