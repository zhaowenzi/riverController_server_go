package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go_server/models"
	"go_server/pkg/e"
	"go_server/pkg/logging"
	"go_server/pkg/util"
	"net/http"
)

func MonitorList(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	riverName := c.DefaultQuery("RiverName", "")

	accessToken := c.Query("access_token")
	claims, _ := util.ParseToken(accessToken)

	data, count := models.GetMonitorAndCount(util.GetOffset(c), util.GetLimit(c), name, riverName, claims.LevelId, claims.Uid)
	code := e.OK
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   e.GetMsg(code),
		"data":  data,
		"count": count,
	})

}

func AddMonitor(c *gin.Context) {
	name := c.PostForm("name")
	RiverName := c.PostForm("RiverName")
	MN_number := c.PostForm("MN_number")
	province := c.PostForm("province")
	city := c.PostForm("city")
	county := c.PostForm("county")
	address := c.PostForm("address")
	longitude := c.PostForm("longitude")
	latitude := c.PostForm("latitude")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("name不能为空")
	valid.Required(RiverName, "RiverName").Message("RiverName不能为空")
	valid.Required(MN_number, "MN_number").Message("MN_number不能为空")
	valid.Required(province, "province").Message("province不能为空")
	valid.Required(city, "city").Message("city不能为空")
	valid.Required(county, "county").Message("county不能为空")
	valid.Required(address, "address").Message("address不能为空")
	valid.Required(longitude, "longitude").Message("longitude不能为空")
	valid.Required(latitude, "latitude").Message("latitude不能为空")

	code := e.PARAM_ERROR
	var err error
	if !valid.HasErrors() {
		data := make(map[string]interface{})
		data["name"] = name
		data["river_name"] = com.StrTo(RiverName).MustInt()
		data["MN_number"] = MN_number
		data["province"] = province
		data["city"] = city
		data["county"] = county
		data["address"] = address
		data["longitude"] = longitude
		data["latitude"] = latitude
		models.AddMonitor(data)
		code = e.OK
	} else {
		for _, validErr := range valid.Errors {
			logging.Info("err.key: %s, err.message: %s", validErr.Key, validErr.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": err,
	})
}

func EditMonitor(c *gin.Context) {
	id := com.StrTo(c.PostForm("id")).MustInt()
	name := c.PostForm("name")
	RiverName := c.PostForm("RiverName")
	MN_number := c.PostForm("MN_number")
	province := c.PostForm("province")
	city := c.PostForm("city")
	county := c.PostForm("county")
	address := c.PostForm("address")
	longitude := c.PostForm("longitude")
	latitude := c.PostForm("latitude")

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.Required(name, "name").Message("name不能为空")
	valid.Required(RiverName, "RiverName").Message("RiverName不能为空")
	valid.Required(MN_number, "MN_number").Message("MN_number不能为空")
	valid.Required(province, "province").Message("province不能为空")
	valid.Required(city, "city").Message("city不能为空")
	valid.Required(county, "county").Message("county不能为空")
	valid.Required(address, "address").Message("address不能为空")
	valid.Required(longitude, "longitude").Message("longitude不能为空")
	valid.Required(latitude, "latitude").Message("latitude不能为空")

	code := e.PARAM_ERROR
	var err error
	if !valid.HasErrors() {
		data := make(map[string]interface{})
		data["name"] = name
		data["river_name"] = com.StrTo(RiverName).MustInt()
		data["MN_number"] = MN_number
		data["province"] = province
		data["city"] = city
		data["county"] = county
		data["address"] = address
		data["longitude"] = longitude
		data["latitude"] = latitude
		models.EditMonitor(id, data)
		if err == nil {
			code = e.OK
		}
	} else {
		for _, validErr := range valid.Errors {
			logging.Info("err.key: %s, err.message: %s", validErr.Key, validErr.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": err,
	})
}

func DelMonitor(c *gin.Context) {
	id := c.PostForm("id")
	valid := validation.Validation{}
	valid.Required(id, "id").Message("id不能为空")

	code := e.PARAM_ERROR
	var err error
	if !valid.HasErrors() {
		err = models.DelMonitor(com.StrTo(id).MustInt())
		if err == nil {
			code = e.OK
		}
	} else {
		for _, validErr := range valid.Errors {
			logging.Info("err.key: %s, err.message: %s", validErr.Key, validErr.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": err,
	})
}