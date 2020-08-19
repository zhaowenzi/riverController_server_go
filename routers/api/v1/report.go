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

func ReportList(c *gin.Context) {
	scene := c.DefaultQuery("scene", "")
	status := c.DefaultQuery("status", "")

	data, count := models.GetReportAndCount(util.GetOffset(c), util.GetLimit(c), scene, status)
	code := e.OK
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   e.GetMsg(code),
		"data":  data,
		"count": count,
	})
}

func GetReportContent(c *gin.Context)  {
	id := com.StrTo(c.Query("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.PARAM_ERROR
	var data interface{}
	if !valid.HasErrors() {
		data = models.GetReport(id)
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

func GetDelegate(c *gin.Context) {
	data := models.GetDelegate()
	code := e.OK
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func DelegateUser(c *gin.Context) {
	id := com.StrTo(c.PostForm("id")).MustInt()
	delegate := com.StrTo(c.PostForm("delegate")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.Min(delegate, 1, "delegate").Message("delegate必须大于0")

	code := e.PARAM_ERROR
	var err error
	if !valid.HasErrors() {
		err = models.DelegateUser(id, delegate)
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

func Audit(c *gin.Context) {
	id := com.StrTo(c.PostForm("id")).MustInt()
	type_ := com.StrTo(c.PostForm("type")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.Min(type_, 1, "type").Message("type必须大于0")

	code := e.PARAM_ERROR
	var err error
	if !valid.HasErrors() {
		err = models.Audit(id, type_)
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
