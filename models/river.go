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

func GetRiversData(id int) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	var rivers []RiverName
	var authRivers []AdminRiver
	var authData []int
	err := db.Table("river_name").Find(&rivers).Error
	if err != nil {
		return nil, err
	}
	var riversData []map[string]interface{}
	for index := range rivers {
		item := make(map[string]interface{})
		item["value"] = rivers[index].ID
		item["title"] = rivers[index].Name
		if rivers[index].Type == 0 {
			item["title"] = item["title"].(string) + "(河道)"
		}
		if rivers[index].Type == 1 {
			item["title"] = item["title"].(string) + "(水库)"
		}
		if rivers[index].Type == 2 {
			item["title"] = item["title"].(string) + "(排污口)"
		}
		if rivers[index].Type == 3 {
			item["title"] = item["title"].(string) + "(管网)"
		}
		if rivers[index].Type == 4 {
			item["title"] = item["title"].(string) + "(楼宇储水池)"
		}
		if rivers[index].Type == 5 {
			item["title"] = item["title"].(string) + "(养殖鱼塘)"
		}
		riversData = append(riversData, item)
	}

	err = db.Table("admin_river").Where("adminId = ?", id).Find(&authRivers).Error
	if err != nil {
		return nil, err
	}

	for index := range authRivers {
		authData = append(authData, authRivers[index].RiverID)
	}
	data["rivers_data"] = riversData
	data["auth_rivers"] = authData

	return data, nil
}

func GetRivers() ([]map[string]interface{}, error) {
	var data []map[string]interface{}
	var rivers []RiverName
	err := db.Table("river_name").Find(&rivers).Error
	if err == nil {
		for index := range rivers {
			item := make(map[string]interface{})
			item["key"] = rivers[index].ID
			item["value"] = rivers[index].Name
			data = append(data, item)
		}
	}
	return data, err
}