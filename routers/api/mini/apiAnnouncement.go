package mini

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go_server/models"
	"go_server/pkg/e"
	"go_server/pkg/logging"
	"net/http"
)

func ApiAnnouncementList(c *gin.Context) {
	data := models.ApiAnnouncementList()
	code := e.OK
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func ApiAnnouncementDetail(c *gin.Context) {
	id := c.PostForm("id")
	valid := validation.Validation{}
	valid.Required(id, "id").Message("id不能为空")

	code := e.PARAM_ERROR
	var data interface{}
	if !valid.HasErrors() {
		data = models.ApiAnnouncementDetail(com.StrTo(id).MustInt())
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