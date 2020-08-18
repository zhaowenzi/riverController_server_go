package models

import "go_server/pkg/util"

type User struct {
	ID         int    `gorm:"primary_key;column:id;type:int;not null" json:"id"`
	Phone      string `gorm:"column:phone;type:varchar(20)" json:"phone"`                // 手机号码
	NickName   string `gorm:"column:nickName;type:varchar(30);not null" json:"nickName"` // 微信名
	Realname   string `gorm:"column:realname;type:varchar(20)" json:"realname"`          // 真实姓名
	RiverName  interface{}    `gorm:"column:river_name;type:int" json:"river_name"`              // 权限河流
	RiverName_ string `json:"riverName"`
	OpenID     string `gorm:"column:openId;type:varchar(100);not null" json:"open_id"`                       // 微信openid
	LevelID    string `gorm:"column:levelId;type:enum('0','1','2','3','4','5','6');not null" json:"levelId"` // 角色权限id
	IsAdmin    string `gorm:"column:is_admin;type:enum('0','1');not null" json:"is_admin"`                   // 是否为管理员
	CreateTime int    `gorm:"column:create_time;type:int;not null" json:"create_time"`                       // 创建时间
	UpdateTime int    `gorm:"column:update_time;type:int;not null" json:"update_time"`                       // 更新时间
	DeleteTime int    `gorm:"column:delete_time;type:int" json:"delete_time"`
	AvatarURL  string `gorm:"column:avatarUrl;type:varchar(255)" json:"avatar_url"` // 头像
}

type PatrolRecord struct {
	ID        int    `gorm:"primary_key;column:id;type:int;not null" json:"id"`
	Title     string `gorm:"column:title;type:varchar(255)" json:"title"`  // 标题
	BeginTime int    `gorm:"column:begin_time;type:int" json:"begin_time"` // 开始时间
	Meters    string `gorm:"column:meters;type:varchar(255)" json:"meter"` // 单位（米）
	Time      string `gorm:"column:time;type:varchar(255)" json:"time"`    // 结束时间
	Month     int    `gorm:"column:month;type:int" json:"month"`           // 月
	Year      int    `gorm:"column:year;type:int" json:"year"`             // 年
}

type ApplyRiver struct {
	ID        int    `gorm:"primary_key;column:id;type:int;not null"`
	UId       int    `gorm:"column:uid;type:int;not null"`        // 用户id
	RiverName int    `gorm:"column:river_name;type:int;not null"` // 申请的河流ID
	Status    int    `gorm:"column:status;type:int;not null"`     // 1-申请 2-通过 3-不通过
	Time      string `gorm:"column:time;type:varchar(255)"`
}

func GetUserAndCount(pageNum int, pageSize int, realName string, riverName string, LevelId int, uid int) (users []User, count int) {
	if LevelId == 1 {
		if riverName == "" {
			db.Table("user").Where("realname LIKE ?", "%"+realName+"%").Offset(pageNum).Limit(pageSize).Find(&users)
			db.Table("user").Where("realname LIKE ?", "%"+realName+"%").Count(&count)
		} else {
			db.Table("user").Where("realname LIKE ?", "%"+realName+"%").Where("river_name = ?", riverName).Offset(pageNum).Limit(pageSize).Find(&users)
			db.Table("user").Where("realname LIKE ?", "%"+realName+"%").Where("river_name = ?", riverName).Count(&count)
		}
		for index := range users {
			var river RiverName
			db.Table("river_name").Where("id = ?", users[index].RiverName).First(&river)
			users[index].RiverName_ = river.Name
		}
	} else {
		var riverIdList []int
		db.Table("admin_river").Where("adminId = ?", uid).Model(&AdminRiver{}).Pluck("riverId", &riverIdList)
		if riverName == "" {
			db.Table("user").Where("realname LIKE ?", "%"+realName+"%").Where("river_name in (?)", riverIdList).Offset(pageNum).Limit(pageSize).Find(&users)
			db.Table("user").Where("realname LIKE ?", "%"+realName+"%").Where("river_name in (?)", riverIdList).Count(&count)
		} else {
			db.Table("user").Where("realname LIKE ?", "%"+realName+"%").Where("river_name in (?)", riverIdList).Where("river_name = ?", riverName).Offset(pageNum).Limit(pageSize).Find(&users)
			db.Table("user").Where("realname LIKE ?", "%"+realName+"%").Where("river_name in (?)", riverIdList).Where("river_name = ?", riverName).Count(&count)
		}
		for index := range users {
			var river RiverName
			db.Table("river_name").Where("id = ?", users[index].RiverName).First(&river)
			users[index].RiverName_ = river.Name
		}
	}
	return
}

func DelUser(id int) error {
	err := db.Table("user").Where("id = ?", id).Delete(User{}).Error
	return err
}

func PatrolList(id int) ([]PatrolRecord, error) {
	var user User
	err := db.Table("user").Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	var patrols []PatrolRecord
	err = db.Table("patrol_record").Where("openid = ?", user.OpenID).Find(&patrols).Error
	if err != nil {
		return nil, err
	}
	return patrols, nil
}

type AuthUserResult struct {
	ID        int    `json:"id"`
	Status    int    `json:"status"`
	Time      string `json:"time"`
	Phone     string `json:"phone"`
	NickName  string `gorm:"column:nickName" json:"nickName"`
	RealName  string `gorm:"column:realname"json:"realName"`
	RiverName string `gorm:"column:name" json:"riverName"`
	Type      string `json:"type"`
}

func GetAuthUser(pageNum int, pageSize int, status string, LevelId int, uid int) (authUserResult []AuthUserResult, count int) {
	if LevelId == 1 {
		if status != "" {
			db.Table("apply_river").Select("apply_river.id, apply_river.status, apply_river.time, user.phone, user.nickName, user.realname, river_name.name, river_name.type").Where("apply_river.status = ?", status).Joins("join user on user.id = apply_river.uid").Joins("join river_name on river_name.id = apply_river.river_name").
				Offset(pageNum).Limit(pageSize).Scan(&authUserResult)
			db.Table("apply_river").Select("apply_river.id, apply_river.status, apply_river.time, user.phone, user.nickName, user.realname, river_name.name, river_name.type").Where("apply_river.status = ?", status).Joins("join user on user.id = apply_river.uid").Joins("join river_name on river_name.id = apply_river.river_name").
				Count(&count)
		} else {
			db.Table("apply_river").Select("apply_river.id, apply_river.status, apply_river.time, user.phone, user.nickName, user.realname, river_name.name, river_name.type").Joins("join user on user.id = apply_river.uid").Joins("join river_name on river_name.id = apply_river.river_name").
				Offset(pageNum).Limit(pageSize).Scan(&authUserResult)
			db.Table("apply_river").Select("apply_river.id, apply_river.status, apply_river.time, user.phone, user.nickName, user.realname, river_name.name, river_name.type").Joins("join user on user.id = apply_river.uid").Joins("join river_name on river_name.id = apply_river.river_name").
				Count(&count)
		}
	} else {
		var riverIdList []int
		db.Table("admin_river").Where("adminId = ?", uid).Model(&AdminRiver{}).Pluck("riverId", &riverIdList)
		if status != "" {
			db.Table("apply_river").Select("apply_river.id, apply_river.status, apply_river.time, user.phone, user.nickName, user.realname, river_name.name, river_name.type").Where("apply_river.status = ?", status).Where("apply_river.river_name in (?)", riverIdList).Joins("join user on user.id = apply_river.uid").Joins("join river_name on river_name.id = apply_river.river_name").
				Offset(pageNum).Limit(pageSize).Scan(&authUserResult)
			db.Table("apply_river").Select("apply_river.id, apply_river.status, apply_river.time, user.phone, user.nickName, user.realname, river_name.name, river_name.type").Where("apply_river.status = ?", status).Where("apply_river.river_name in (?)", riverIdList).Joins("join user on user.id = apply_river.uid").Joins("join river_name on river_name.id = apply_river.river_name").
				Count(&count)
		} else {
			db.Table("apply_river").Select("apply_river.id, apply_river.status, apply_river.time, user.phone, user.nickName, user.realname, river_name.name, river_name.type").Where("apply_river.river_name in (?)", riverIdList).Joins("join user on user.id = apply_river.uid").Joins("join river_name on river_name.id = apply_river.river_name").
				Offset(pageNum).Limit(pageSize).Scan(&authUserResult)
			db.Table("apply_river").Select("apply_river.id, apply_river.status, apply_river.time, user.phone, user.nickName, user.realname, river_name.name, river_name.type").Where("apply_river.river_name in (?)", riverIdList).Joins("join user on user.id = apply_river.uid").Joins("join river_name on river_name.id = apply_river.river_name").
				Count(&count)
		}
	}
	for index := range authUserResult {
		authUserResult[index].Type = util.GetRiverTypeFromId(authUserResult[index].Type)
	}
	return
}

func AuthUserRiver(id int, status int) error {
	var err error
	if status == 3 {
		err = db.Table("apply_river").Model(&ApplyRiver{}).Where("id = ?", id).Update("status", status).Error
	}
	if status == 2 {
		var applyRiver ApplyRiver
		db.Table("apply_river").Where("id = ?", id).First(&applyRiver)
		err = db.Table("apply_river").Model(&ApplyRiver{}).Where("id = ?", id).Update("status", status).Error
		if err == nil {
			err = db.Table("user").Model(&User{}).Where("id = ?", applyRiver.UId).Update("river_name", applyRiver.RiverName).Error
		}
	}
	if status == 4 {
		var applyRiver ApplyRiver
		db.Table("apply_river").Where("id = ?", id).First(&applyRiver)
		err = db.Table("apply_river").Model(&ApplyRiver{}).Where("id = ?", id).Update("status", 3).Error
		if err == nil {
			err = db.Table("user").Model(&User{}).Where("id = ?", applyRiver.UId).Update("river_name", nil).Error
		}
	}
	return err
}
