package models

type AdminRiver struct {
	ID      int `gorm:"primary_key;column:id;type:int;not null"`
	AdminID int `gorm:"column:adminId;type:int;not null"`
	RiverID int `gorm:"column:riverId;type:int;not null"`
}