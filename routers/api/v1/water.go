package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go_server/models"
	"go_server/pkg/e"
	"go_server/pkg/logging"
	"net/http"
)

func SceneParamParser(c *gin.Context) {
	data, count := models.SceneParamParser()
	code := e.OK
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   e.GetMsg(code),
		"data":  data,
		"count": count,
	})
}

func WaterParamParser(c *gin.Context) {
	sid := c.Query("sid")
	valid := validation.Validation{}
	valid.Required(sid, "sid").Message("SID必须存在")

	code := e.PARAM_ERROR
	var data interface{}
	if !valid.HasErrors() {
		data = models.WaterParamParser(com.StrTo(sid).MustInt())
		code = e.OK
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

func EditSceneParam(c *gin.Context) {
	id := com.StrTo(c.PostForm("id")).MustInt()
	temp_min := c.PostForm("temp_min")
	temp_max := c.PostForm("temp_max")
	ph_min := c.PostForm("ph_min")
	ph_max := c.PostForm("ph_max")
	do_min := c.PostForm("do_min")
	do_max := c.PostForm("do_max")
	cod_min := c.PostForm("cod_min")
	cod_max := c.PostForm("cod_max")
	nh3_n_min := c.PostForm("nh3_n_min")
	nh3_n_max := c.PostForm("nh3_n_max")
	sd_min := c.PostForm("sd_min")
	sd_max := c.PostForm("sd_max")
	conductivity_min := c.PostForm("conductivity_min")
	conductivity_max := c.PostForm("conductivity_max")
	chlorine_min := c.PostForm("chlorine_min")
	chlorine_max := c.PostForm("chlorine_max")
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.PARAM_ERROR
	var err error
	if !valid.HasErrors() {
		data := make(map[string]interface{})
		data["temp_min"] = temp_min
		data["temp_max"] = temp_max
		data["ph_min"] = ph_min
		data["ph_max"] = ph_max
		data["do_min"] = do_min
		data["do_max"] = do_max
		data["cod_min"] = cod_min
		data["cod_max"] = cod_max
		data["nh3_n_min"] = nh3_n_min
		data["nh3_n_max"] = nh3_n_max
		data["sd_min"] = sd_min
		data["sd_max"] = sd_max
		data["conductivity_min"] = conductivity_min
		data["conductivity_max"] = conductivity_max
		data["chlorine_min"] = chlorine_min
		data["chlorine_max"] = chlorine_max
		err = models.EditSceneParam(id, data)
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

func EditScene(c *gin.Context) {
	id := com.StrTo(c.PostForm("id")).MustInt()
	temp := c.PostForm("temp")
	ph := c.PostForm("ph")
	do := c.PostForm("do")
	cod := c.PostForm("cod")
	nh3_n := c.PostForm("nh3_n")
	conductivity := c.PostForm("conductivity")
	chlorine := c.PostForm("chlorine")

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	code := e.PARAM_ERROR
	var err error
	if !valid.HasErrors() {
		data := make(map[string]interface{})
		if temp == "on" {
			data["temp"] = "1"
		} else {
			data["temp"] = "0"
		}
		if ph == "on" {
			data["ph"] = "1"
		} else {
			data["ph"] = "0"
		}
		if do == "on" {
			data["do"] = "1"
		} else {
			data["do"] = "0"
		}
		if cod == "on" {
			data["cod"] = "1"
		} else {
			data["cod"] = "0"
		}
		if nh3_n == "on" {
			data["nh3_n"] = "1"
		} else {
			data["nh3_n"] = "0"
		}
		if conductivity == "on" {
			data["conductivity"] = "1"
		} else {
			data["conductivity"] = "0"
		}
		if chlorine == "on" {
			data["chlorine"] = "1"
		} else {
			data["chlorine"] = "0"
		}
		err = models.EditScene(id, data)
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
