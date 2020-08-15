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

func RiverList(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	address := c.DefaultQuery("address", "")
	type_ := c.DefaultQuery("type", "")

	accessToken := c.Query("access_token")
	claims, _ := util.ParseToken(accessToken)

	data, count := models.GetRiverAndCount(util.GetOffset(c), util.GetLimit(c), name, address, type_, claims.LevelId, claims.Uid)
	code := e.OK
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   e.GetMsg(code),
		"data":  data,
		"count": count,
	})
}

func AddRiver(c *gin.Context) {
	name := c.PostForm("name")
	address := c.PostForm("address")
	type_ := c.PostForm("type")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("name不能为空")
	valid.Required(address, "address").Message("address不能为空")
	valid.Required(type_, "type").Message("type不能为空")

	code := e.PARAM_ERROR
	var err error
	if !valid.HasErrors() {
		data := make(map[string]interface{})
		data["name"] = name
		data["address"] = address
		data["type"] = com.StrTo(type_).MustInt()
		err = models.AddRiver(data)
		if err == nil {
			code = e.OK
		} else {
			code = e.KEY_DUMP
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

func DelRiver(c *gin.Context) {
	id := c.PostForm("id")
	valid := validation.Validation{}
	valid.Required(id, "id").Message("id不能为空")

	code := e.PARAM_ERROR
	var err error
	if !valid.HasErrors() {
		err = models.DelRiver(com.StrTo(id).MustInt())
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

func EditRiver(c *gin.Context) {
	id := com.StrTo(c.PostForm("id")).MustInt()
	name := c.PostForm("name")
	address := c.PostForm("address")
	type_ := c.PostForm("type")

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.Required(name, "name").Message("name不能为空")
	valid.Required(address, "address").Message("address不能为空")
	valid.Required(type_, "type").Message("type不能为空")

	code := e.PARAM_ERROR
	var err error
	if !valid.HasErrors() {
		data := make(map[string]interface{})
		data["name"] = name
		data["address"] = address
		data["type"] = com.StrTo(type_).MustInt()
		err = models.EditRiver(id, data);
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

func AuthList(c *gin.Context) {
	id := com.StrTo(c.Query("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.PARAM_ERROR
	var data map[string]interface{}
	var err error
	if !valid.HasErrors() {
		data, err = models.GetRiversData(id)
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
		"data": data,
	})
}

func AllRiverParser(c *gin.Context) {
	code := e.PARAM_ERROR
	data, err := models.GetRivers()
	if err == nil {
		code = e.OK
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}