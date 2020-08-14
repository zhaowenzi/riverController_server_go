package models

type RiverName struct {
	ID      int    `gorm:"primary_key;column:id;type:int;not null" json:"id"`
	Name    string `gorm:"unique;unique;column:name;type:varchar(50);not null" json:"name"`
	Address string `gorm:"column:address;type:varchar(255)" json:"address"`
	Type    int    `gorm:"column:type;type:int;not null" json:"type"`
}

func GetRiverAndCount(pageNum int, pageSize int, name string, address string, type_ string, LevelId int, uid int) (rivers []RiverName, count int) {
	if LevelId == 1 {
		if type_ == "" {
			db.Table("river_name").Where("name LIKE ?", "%"+name+"%").Where("address LIKE ?", "%"+address+"%").
				Offset(pageNum).Limit(pageSize).Find(&rivers)
			db.Table("river_name").Where("name LIKE ?", "%"+name+"%").Where("address LIKE ?", "%"+address+"%").
				Count(&count)
		} else {
			db.Table("river_name").Where("name LIKE ?", "%"+name+"%").Where("address LIKE ?", "%"+address+"%").Where("type = ?", type_).
				Offset(pageNum).Limit(pageSize).Find(&rivers)
			db.Table("river_name").Where("name LIKE ?", "%"+name+"%").Where("address LIKE ?", "%"+address+"%").Where("type = ?", type_).
				Count(&count)
		}
	} else {
		var riverIdList []int
		db.Table("admin_river").Where("adminId = ?", uid).Model(&AdminRiver{}).Pluck("riverId", &riverIdList)
		if type_ == "" {
			db.Table("river_name").Where("name LIKE ?", "%"+name+"%").Where("address LIKE ?", "%"+address+"%").Where("id in (?)", riverIdList).
				Offset(pageNum).Limit(pageSize).Find(&rivers)
			db.Table("river_name").Where("name LIKE ?", "%"+name+"%").Where("address LIKE ?", "%"+address+"%").Where("id in (?)", riverIdList).
				Count(&count)
		} else {
			db.Table("river_name").Where("name LIKE ?", "%"+name+"%").Where("address LIKE ?", "%"+address+"%").Where("type = ?", type_).Where("id in (?)", riverIdList).
				Offset(pageNum).Limit(pageSize).Find(&rivers)
			db.Table("river_name").Where("name LIKE ?", "%"+name+"%").Where("address LIKE ?", "%"+address+"%").Where("type = ?", type_).Where("id in (?)", riverIdList).
				Count(&count)
		}
	}
	return
}

func AddRiver(data map[string]interface{}) error {

	err := db.Table("river_name").Create(&RiverName{
		Name: data["name"].(string),
		Address: data["address"].(string),
		Type: data["type"].(int),
	}).Error
	return err
}

func DelRiver(id int) error {
	err := db.Table("river_name").Where("id = ?", id).Delete(RiverName{}).Error
	return err
}

func EditRiver(id int, data map[string]interface{}) error {
	err := db.Table("river_name").Model(&RiverName{}).Where("id = ?", id).Update(data).Error
	return err
}