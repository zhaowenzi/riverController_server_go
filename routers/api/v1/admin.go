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

// 返回管理员列表
func AdminUsers(c *gin.Context) {

	name := c.DefaultQuery("name", "")
	nickname := c.DefaultQuery("nickname", "")
	data, count := models.GetAdminListsAndCount(util.GetOffset(c), util.GetLimit(c), name, nickname)
	code := e.OK
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   e.GetMsg(code),
		"data":  data,
		"count": count,
	})
}

// 新增管理员
func AddAdmin(c *gin.Context) {

	name := c.PostForm("name")
	nickname := c.PostForm("nickname")
	level_ID := c.PostForm("level_ID")
	phone := c.PostForm("phone")
	password := c.PostForm("password")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("账号不能为空")
	valid.Required(nickname, "nickname").Message("昵称不能为空")
	valid.Required(level_ID, "level_ID").Message("权限不能为空")
	valid.Required(phone, "phone").Message("手机号不能为空")
	valid.Required(password, "password").Message("密码不能为空")

	code := e.PARAM_ERROR
	var err error
	if !valid.HasErrors() {
		err = models.AddAdmin(name, nickname, level_ID, phone, util.PasswordMd5(password))
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

// 删除管理员
func DelAdmin(c *gin.Context) {
	id := c.PostForm("id")
	valid := validation.Validation{}
	valid.Required(id, "id").Message("id不能为空")

	code := e.PARAM_ERROR
	var err error
	if !valid.HasErrors() {
		err = models.DelAdmin(com.StrTo(id).MustInt())
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

// 编辑管理员
func EditAdmin(c *gin.Context) {
	valid := validation.Validation{}

	id := com.StrTo(c.PostForm("id")).MustInt()
	name := c.PostForm("name")
	nickname := c.PostForm("nickname")
	level_ID := c.PostForm("level_ID")
	phone := c.PostForm("phone")

	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.Required(name, "name").Message("账号不能为空")
	valid.Required(nickname, "nickname").Message("昵称不能为空")
	valid.Required(level_ID, "level_ID").Message("权限不能为空")
	valid.Required(phone, "phone").Message("手机号不能为空")

	code := e.PARAM_ERROR
	var err error
	if !valid.HasErrors() {
		data := make(map[string]interface{})
		data["name"] = name
		data["nickname"] = nickname
		data["level_ID"] = level_ID
		data["phone"] = phone
		err = models.EditAdmin(id, data)
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

// 管理员菜单
func AdminMenu(c *gin.Context) {
	accessToken := c.Query("access_token")
	claims, _ := util.ParseToken(accessToken)
	if claims.LevelId == 1 {
		c.JSON(http.StatusOK, gin.H{
			"code": e.OK,
			"msg": e.GetMsg(e.OK),
			"data": util.SuperAdminMenu(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": e.OK,
			"msg": e.GetMsg(e.OK),
			"data": util.NormAdminMenu(),
		})
	}
}
