package models

type MonitoringPoint struct {
	ID               int    `gorm:"primary_key;column:id;type:int;not null" json:"id"`
	Longitude        string `gorm:"column:longitude;type:varchar(50)" json:"longitude"` // 经度
	Latitude         string `gorm:"column:latitude;type:varchar(50)" json:"latitude"`  // 纬度
	Scene            string `gorm:"column:scene;type:varchar(50)" json:"scene"`     // 场景 （目前无用）
	Name             string `gorm:"column:name;type:varchar(255)" json:"name"`     // 检测点名字
	RiverName        int    `gorm:"column:river_name;type:int" json:"river_name"`
	MNNumber         string `gorm:"column:MN_number;type:varchar(27)" json:"MN_number"` // 硬件设备号
	CreateTime       int    `gorm:"column:create_time;type:int" json:"create_time"`
	UpdateTime       int    `gorm:"column:update_time;type:int" json:"update_time"`
	Address          string `gorm:"column:address;type:varchar(100)" json:"address"`           // 地址
	OpenID           string `gorm:"column:openId;type:varchar(100)" json:"openid"`            // openId
	AdminID          int    `gorm:"column:admin_id;type:int" json:"admin_id"`                   // 管理员id
	TempMin          string `gorm:"column:temp_min;type:varchar(255)" json:"temp_min"`          // 温度最小值
	TempMax          string `gorm:"column:temp_max;type:varchar(255)" json:"temp_max"`          // 温度最大值
	PhMin            string `gorm:"column:ph_min;type:varchar(255)" json:"ph_min"`            // ph最小值
	PhMax            string `gorm:"column:ph_max;type:varchar(255)" json:"ph_max"`            // ph最大值
	DoMin            string `gorm:"column:do_min;type:varchar(255)" json:"do_min"`            // do溶解氧最小值
	DoMax            string `gorm:"column:do_max;type:varchar(255)" json:"do_max"`            // do溶解氧最大值
	CodMin           string `gorm:"column:cod_min;type:varchar(255)" json:"cod_min"`           // cod化学需氧量最小值
	CodMax           string `gorm:"column:cod_max;type:varchar(255)" json:"cod_max"`           // cod化学需氧量最大值
	Nh3NMin          string `gorm:"column:nh3_n_min;type:varchar(255)" json:"nh_3_n_min"`         // 氨氮最小值
	Nh3NMax          string `gorm:"column:nh3_n_max;type:varchar(255)" json:"nh_3_n_max"`         // 氨氮最大值
	SdMin            string `gorm:"column:sd_min;type:varchar(255)" json:"sd_min"`            // sd浊度最小值
	SdMax            string `gorm:"column:sd_max;type:varchar(255)" json:"sd_max"`            // sd浊度最大值
	ConductivityMin  string `gorm:"column:conductivity_min;type:varchar(255)" json:"conductivity_min"`  // 电导率最小值
	ConductivityMax  string `gorm:"column:conductivity_max;type:varchar(255)" json:"conductivity_max"`  // 电导率最大值
	Province         string `gorm:"column:province;type:varchar(255)" json:"province"`          // 省
	City             string `gorm:"column:city;type:varchar(255)" json:"city"`              // 市
	County           string `gorm:"column:county;type:varchar(255)" json:"county"`            // 区
	ChlorineMin      string `gorm:"column:chlorine_min;type:varchar(255)" json:"chlorine_min"`      // 余氯最小值
	ChlorineMax      string `gorm:"column:chlorine_max;type:varchar(255)" json:"chlorine_max"`      // 余氯最大值
	WaterFlowMin     string `gorm:"column:WaterFlow_min;type:varchar(255)" json:"water_flow_min"`     // 水流最小值
	WaterFlowMax     string `gorm:"column:WaterFlow_max;type:varchar(255)" json:"water_flow_max"`     // 水流最大值
	WaterPressureMin string `gorm:"column:WaterPressure_min;type:varchar(255)" json:"water_pressure_min"` // 水压最小值
	WaterPressureMax string `gorm:"column:WaterPressure_max;type:varchar(255)" json:"water_pressure_max"` // 水压最大值
}

func GetMonitorAndCount(pageNum int, pageSize int, name string, riverName string, LevelId int, uid int) (monitors []MonitoringPoint, count int){
	if LevelId == 1 {
		if riverName == "" {
			db.Table("monitoring_point").Where("name LIKE ?", "%"+name+"%").
				Offset(pageNum).Limit(pageSize).Find(&monitors)
			db.Table("monitoring_point").Where("name LIKE ?", "%"+name+"%").
				Count(&count)
		} else {
			db.Table("monitoring_point").Where("name LIKE ?", "%"+name+"%").Where("river_name = ?", riverName).
				Offset(pageNum).Limit(pageSize).Find(&monitors)
			db.Table("monitoring_point").Where("name LIKE ?", "%"+name+"%").Where("river_name = ?", riverName).
				Count(&count)
		}
	} else {
		var riverIdList []int
		db.Table("admin_river").Where("adminId = ?", uid).Model(&AdminRiver{}).Pluck("riverId", &riverIdList)
		if riverName == "" {
			db.Table("monitoring_point").Where("name LIKE ?", "%"+name+"%").Where("river_name in (?)", riverIdList).
				Offset(pageNum).Limit(pageSize).Find(&monitors)
			db.Table("monitoring_point").Where("name LIKE ?", "%"+name+"%").Where("river_name in (?)", riverIdList).
				Count(&count)
		} else {
			db.Table("monitoring_point").Where("name LIKE ?", "%"+name+"%").Where("river_name = ?", riverName).Where("river_name in (?)", riverIdList).
				Offset(pageNum).Limit(pageSize).Find(&monitors)
			db.Table("monitoring_point").Where("name LIKE ?", "%"+name+"%").Where("river_name = ?", riverName).Where("river_name in (?)", riverIdList).
				Count(&count)
		}
	}
	return
}

func AddMonitor (data map[string]interface{}) bool {
	db.Table("monitoring_point").Create(&MonitoringPoint{
		Name: data["name"].(string),
		RiverName: data["river_name"].(int),
		MNNumber: data["MN_number"].(string),
		Province: data["province"].(string),
		City: data["city"].(string),
		County: data["county"].(string),
		Address: data["address"].(string),
		Longitude: data["longitude"].(string),
		Latitude: data["latitude"].(string),
	})
	return true
}

func EditMonitor(id int, data map[string]interface{}) error {
	err := db.Table("monitoring_point").Model(&MonitoringPoint{}).Where("id = ?", id).Update(data).Error
	return err
}

func DelMonitor(id int) error {
	err := db.Table("monitoring_point").Where("id = ?", id).Delete(MonitoringPoint{}).Error
	return err
}