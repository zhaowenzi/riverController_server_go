package models

import (
	"log"
	"strconv"
)

type Report struct {
	ID           int    `gorm:"primary_key;column:id;type:int;not null" json:"id"`
	UId          int    `gorm:"column:uid;type:int" json:"uid"`
	Openid       string `gorm:"column:openid;type:varchar(255)" json:"openid"`
	RiverName    string `gorm:"column:river_name;type:varchar(255)" json:"river_name"` // 巡河名称
	Address      string `gorm:"column:address;type:varchar(255)" json:"address"`       // 举报地址
	Scene        string `gorm:"column:scene;type:varchar(255)" json:"scene"`           // 场景
	Label        string `gorm:"column:label;type:varchar(255)" json:"label"`           // 标签
	Describe     string `gorm:"column:describe;type:varchar(255)" json:"describe"`     // 描述
	Suggest      string `gorm:"column:suggest;type:varchar(255)" json:"suggest"`       // 建议
	Phone        string `gorm:"column:phone;type:varchar(20)" json:"phone"`            // 手机
	Status       int    `gorm:"column:status;type:int" json:"status"`                  // 状态
	Audit        int    `gorm:"column:audit;type:int" json:"audit"`                    // 审核
	Delegate     string `gorm:"column:delegate;type:varchar(255)" json:"delegate"`     // 授权
	Map          string `gorm:"column:map;type:varchar(255)" json:"map"`               // 地点
	EndTime      int    `gorm:"column:end_time;type:int" json:"end_time"`              // 结束时间
	CreateTime   int    `gorm:"column:create_time;type:int" json:"create_time"`
	UpdateTime   int    `gorm:"column:update_time;type:int" json:"update_time"` // 更新时间
	DeleteTime   int    `gorm:"column:delete_time;type:int" json:"delete_time"` // 软删除时间
	Time         string `gorm:"column:time;type:varchar(255)" json:"time"`
	DelegateName string `json:"delegatename"`
}

type ReportImage struct {
	ID         int    `gorm:"primary_key;column:id;type:int;not null" json:"id"`
	Openid     string `gorm:"column:openid;type:varchar(255)" json:"openid"` // 微信openid
	Image      string `gorm:"column:image;type:varchar(255)" json:"image"`  // 图片
	CreateTime int    `gorm:"column:create_time;type:int" json:"create_time"`
	UpdateTime int    `gorm:"column:update_time;type:int" json:"update_time"`
	Time       string `gorm:"column:time;type:varchar(255)" json:"time"` // 时间
}

func GetReportAndCount(pageNum int, pageSize int, scene string, status string) (reports []Report, count int) {
	if scene != "" && status != "" {
		db.Table("report").Where("scene = ?", scene).Where("status = ?", status).
			Offset(pageNum).Limit(pageSize).Find(&reports)
		db.Table("report").Where("scene = ?", scene).Where("status = ?", status).
			Count(&count)
	} else if scene != "" && status == "" {
		db.Table("report").Where("scene = ?", scene).
			Offset(pageNum).Limit(pageSize).Find(&reports)
		db.Table("report").Where("scene = ?", scene).
			Count(&count)
	} else if scene == "" && status != "" {
		db.Table("report").Where("status = ?", status).
			Offset(pageNum).Limit(pageSize).Find(&reports)
		db.Table("report").Where("status = ?", status).
			Count(&count)
	} else {
		db.Table("report").
			Offset(pageNum).Limit(pageSize).Find(&reports)
		db.Table("report").
			Count(&count)
	}
	for index := range reports {
		if reports[index].Openid != "" {
			var user User
			db.Table("user").Where("openId = ?", reports[index].Openid).First(&user)
			if user.Realname != "" {
				reports[index].Openid = user.Realname
			} else {
				reports[index].Openid = user.NickName
			}
		}
		if reports[index].Delegate != "0" {
			var user User
			db.Table("user").Where("id = ?", reports[index].Delegate).First(&user)
			if user.ID != 0 {
				if user.Realname != "" {
					reports[index].DelegateName = user.Realname
				} else {
					reports[index].DelegateName = user.NickName
				}
			} else {
				reports[index].DelegateName = "用户已删除，重新分配给其他人"
			}
		} else {
			reports[index].DelegateName = "未分配给处理人"
		}
	}
	return
}

func GetReport(id int) (photo []string)  {
	var reportImages []ReportImage
	var report Report
	db.Table("report").Where("id = ?", id).First(&report)
	db.Table("report_image").Where("openid = ?", report.Openid).Where("time = ?", report.Time).Find(&reportImages)
	for index := range reportImages {
		photo = append(photo, reportImages[index].Image)
	}
	return
}

func GetDelegate() []interface{} {
	var users []User
	db.Table("user").Find(&users)
	var data []interface{}
	for index := range users {
		if users[index].RiverName != nil {
			item := make(map[string]interface{})
			item["id"] = users[index].ID
			item["realName"] = users[index].Realname
			data = append(data, item)
		}
	}
	return data
}

func DelegateUser(id int, delegate int) error {
	data := make(map[string]interface{})
	data["status"] = 0
	data["delegate"] = strconv.Itoa(delegate)
	log.Println(data)
	err := db.Table("report").Model(&Report{}).Where("id = ?", id).Update(data).Error
	if err != nil {
		return err
	}
	return nil
}

func Audit(id int, type_ int) error {
	var err error
	if type_ == 1 {
		err = db.Table("report").Where("id = ?", id).Update("status", 3).Error
	} else if type_ == 2 {
		err = db.Table("report").Where("id = ?", id).Update("status", 4).Error
	}
	return err
}