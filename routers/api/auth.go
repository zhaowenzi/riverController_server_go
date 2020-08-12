package api

import (
	"go_server/models"
	"go_server/pkg/e"
	"go_server/pkg/logging"
	"go_server/pkg/util"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.PostForm("name")
	password := c.PostForm("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.PARAM_ERROR
	if ok {
		passwordMd5 := util.PasswordMd5(password)
		isExist, uId, levelID := models.CheckAuth(username, passwordMd5)
		if isExist {
			token, err := util.GenerateToken(uId, username, levelID)
			if err != nil {
				code = e.AUTHORIZATION_ERROR
			} else {
				data["access_token"] = token
				data["level_ID"] = levelID

				code = e.OK
			}

		} else {
			code = e.AUTHORIZATION_ERROR
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
