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

func UserList(c *gin.Context) {
	realName := c.DefaultQuery("RealName", "")
	riverName := c.DefaultQuery("RiverName", "")

	accessToken := c.Query("access_token")
	claims, _ := util.ParseToken(accessToken)

	data, count := models.GetUserAndCount(util.GetOffset(c), util.GetLimit(c), realName, riverName, claims.LevelId, claims.Uid)
	code := e.OK
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   e.GetMsg(code),
		"data":  data,
		"count": count,
	})
}

func DelUser(c *gin.Context) {
	id := c.PostForm("id")
	valid := validation.Validation{}
	valid.Required(id, "id").Message("id不能为空")

	code := e.PARAM_ERROR
	var err error
	if !valid.HasErrors() {
		err = models.DelUser(com.StrTo(id).MustInt())
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

func PatrolList(c *gin.Context) {
	id := c.Query("id")
	valid := validation.Validation{}
	valid.Required(id, "id").Message("id不能为空")

	code := e.PARAM_ERROR
	if !valid.HasErrors() {
		data, err := models.PatrolList(com.StrTo(id).MustInt())
		if err == nil {
			code = e.OK
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			return
		}
	} else {
		for _, validErr := range valid.Errors {
			logging.Info("err.key: %s, err.message: %s", validErr.Key, validErr.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": "",
	})
}

func AuthUser(c *gin.Context) {
	status := c.DefaultQuery("status", "")
	accessToken := c.Query("access_token")
	claims, _ := util.ParseToken(accessToken)

	data, count := models.GetAuthUser(util.GetOffset(c), util.GetLimit(c), status, claims.LevelId, claims.Uid)
	code := e.OK
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   e.GetMsg(code),
		"data":  data,
		"count": count,
	})
}

func AuthUserRiver(c *gin.Context) {
	id := com.StrTo(c.PostForm("id")).MustInt()
	status := com.StrTo(c.PostForm("status")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.Min(status, 1, "status").Message("status必须大于0")

	code := e.PARAM_ERROR
	var err error
	if !valid.HasErrors() {
		err = models.AuthUserRiver(id, status);
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