package models

type SceneParam struct {
	ID            int    `gorm:"primary_key;column:id;type:int;not null" json:"id"`
	Sid           int16  `gorm:"column:sid;type:smallint;not null" json:"sid"`               // 场景id
	Temp          string `gorm:"column:temp;type:enum('0','1');not null" json:"temp"`         // 温度
	Ph            string `gorm:"column:ph;type:enum('0','1');not null" json:"ph"`           // pH值
	Do            string `gorm:"column:do;type:enum('0','1');not null" json:"do"`           // do溶解氧
	Cod           string `gorm:"column:cod;type:enum('0','1');not null" json:"cod"`          // cod化学需氧量
	Nh3N          string `gorm:"column:nh3_n;type:enum('0','1');not null" json:"nh3_n"`        // 氨氮
	Sd            string `gorm:"column:sd;type:enum('0','1');not null" json:"sd"`           // sd浊度
	Conductivity  string `gorm:"column:conductivity;type:enum('0','1');not null" json:"conductivity"` // 电导率
	Chlorine      string `gorm:"column:chlorine;type:enum('0','1');not null" json:"chlorine"`     // 余氯
	WaterPressure string `gorm:"column:WaterPressure;type:enum('0','1')"`         // 水压
	WaterFlow     string `gorm:"column:WaterFlow;type:enum('0','1');not null"`    // 水流
	CreateTime    int    `gorm:"column:create_time;type:int;not null"`
	UpdateTime    int    `gorm:"column:update_time;type:int;not null"`
	MN            string `gorm:"column:MN;type:varchar(255)"`
}

type WaterParam struct {
	ID               int    `gorm:"primary_key;column:id;type:int;not null" json:"id"`
	Wid              int    `gorm:"column:wid;type:int;not null" json:"wid"`               // 水质类别id
	Sid              int    `gorm:"column:sid;type:int;not null" json:"sid"`               // 场景id
	TempMin          string `gorm:"column:temp_min;type:varchar(10)" json:"temp_min"`           // 温度范围MIN值
	TempMax          string `gorm:"column:temp_max;type:varchar(10)" json:"temp_max"`           // 温度范围MAX值
	PhMin            string `gorm:"column:ph_min;type:varchar(10)" json:"ph_min"`             // pH范围MIN值
	PhMax            string `gorm:"column:ph_max;type:varchar(10)" json:"ph_max"`             // pH范围MAX值
	DoMin            string `gorm:"column:do_min;type:varchar(15)" json:"do_min"`             // DO溶解氧MIN值
	DoMax            string `gorm:"column:do_max;type:varchar(15)" json:"do_max"`             // DO溶解氧MAX值
	CodMin           string `gorm:"column:cod_min;type:varchar(15)" json:"cod_min"`            // COD化学需氧量MIN值
	CodMax           string `gorm:"column:cod_max;type:varchar(15)" json:"cod_max"`            // COD化学需氧量MAX值
	Nh3NMin          string `gorm:"column:nh3_n_min;type:varchar(15)" json:"nh3_n_min"`          // NH3-N氨氮MIN值
	Nh3NMax          string `gorm:"column:nh3_n_max;type:varchar(15)" json:"nh3_n_max"`          // NH3-N氨氮MAX值
	SdMin            string `gorm:"column:sd_min;type:varchar(15)" json:"sd_min"`             // SD浊度MIN值
	SdMax            string `gorm:"column:sd_max;type:varchar(15)" json:"sd_max"`             // SD浊度MAX值
	ConductivityMin  string `gorm:"column:conductivity_min;type:varchar(15)" json:"conductivity_min"`   // 导电率范围min值
	ConductivityMax  string `gorm:"column:conductivity_max;type:varchar(15)" json:"conductivity_max"`   // 导电率范围max值
	WaterPressureMin string `gorm:"column:WaterPressure_min;type:varchar(255)"` // 水压MIN值
	WaterPressureMax string `gorm:"column:WaterPressure_max;type:varchar(255)"` // 水压MAX值
	WaterFlowMin     string `gorm:"column:WaterFlow_min;type:varchar(255)"`     // 水流MIN值
	WaterFlowMax     string `gorm:"column:WaterFlow_max;type:varchar(255)"`     // 水流MAX值
	ChlorineMin      string `gorm:"column:chlorine_min;type:varchar(255)" json:"chlorine_min"`      // 余氯MIN值
	ChlorineMax      string `gorm:"column:chlorine_max;type:varchar(255)" json:"chlorine_max"`      // 余氯MAX值
	CreateTime       int    `gorm:"column:create_time;type:int;not null"`       // 创建时间
	UpdateTime       int    `gorm:"column:update_time;type:int;not null"`       // 更新时间
}

func SceneParamParser() (sceneParams []SceneParam, count int) {
	db.Table("scene_param").Find(&sceneParams)
	db.Table("scene_param").Count(&count)
	return
}

func WaterParamParser(sid int) (waterParam []WaterParam) {
	db.Table("water_param").Where("sid = ?", sid).Find(&waterParam)
	return
}

func EditSceneParam(id int, data map[string]interface{}) error {
	err := db.Table("water_param").Model(&WaterParam{}).Where("id = ?", id).Update(data).Error
	return err
}

func EditScene(id int, data map[string]interface{}) error {
	err := db.Table("scene_param").Model(&SceneParam{}).Where("id = ?", id).Update(data).Error
	return err
}