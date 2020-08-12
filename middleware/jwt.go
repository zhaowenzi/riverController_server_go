package jwt

import (
	"go_server/pkg/e"
	"go_server/pkg/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int

		code = e.OK
		token := c.Query("access_token")
		if token == "" {
			code = e.PARAM_ERROR
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.AUTHORIZATION_ERROR
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.AUTHORIZATION_ERROR
			}
		}

		if code != e.OK {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": "不是管理员",
			})

			c.Abort()
			return
		}

		c.Next()
	}
}

// 判断是否为超级管理员
func CheckSuperAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.Query("access_token")
		claims, _ := util.ParseToken(accessToken)
		if claims.LevelId != 1 {
			code := e.AUTHORIZATION_ERROR
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": "不是超级管理员",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
