package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type AdminList struct {
	ID         int    `gorm:"primary_key" json:"id"`
	Name       string `json:"name"`
	NickName   string `gorm:"column:nickname" json:"nickname"`
	Level_ID   int    `gorm:"column:level_ID" json:"level_ID"`
	Phone      string `json:"phone"`
	CreateTime string `gorm:"column:create_time" json:"create_time"`
}

func GetAdminListsAndCount(pageNum int, pageSize int, name string, nickname string) (adminLists []AdminList, count int) {
	db.Table("admin").Where("name LIKE ? AND nickname LIKE ?", "%"+name+"%", "%"+nickname+"%").
		Offset(pageNum).Limit(pageSize).Find(&adminLists)
	db.Table("admin").Where("name LIKE ? AND nickname LIKE ?", "%"+name+"%", "%"+nickname+"%").
		Count(&count)
	return
}

type addAdmin struct {
	ID         int    `gorm:"primary_key" json:"id"`
	Name       string `json:"name"`
	NickName   string `gorm:"column:nickname" json:"nickname"`
	Level      string `gorm:"column:level_ID" json:"level_ID"`
	Phone      string `gorm:"column:phone" json:"phone"`
	Password   string `gorm:"column:password" json:"password"`
	CreateTime int    `gorm:"create_time" json:create_time`
	UpdateTime int    `gorm:"update_time" json:update_time`
}

func AddAdmin(name string, nickname string, level_ID string, phone string, password string) error {
	admin := addAdmin{Name: name, NickName: nickname, Level: level_ID, Phone: phone, Password: password}
	err := db.Table("admin").Create(&admin).Error
	return err
}

func (addAdmin *addAdmin) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreateTime", time.Now().Unix())
	scope.SetColumn("UpdateTime", time.Now().Unix())
	return nil
}

func DelAdmin(id int) error {
	err := db.Table("admin").Where("id = ?", id).Delete(addAdmin{}).Error
	return err
}

func EditAdmin(id int, data interface{}) error {
	err := db.Table("admin").Model(&addAdmin{}).Where("id = ?", id).Update(data).Error
	return err
}
