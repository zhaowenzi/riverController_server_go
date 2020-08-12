package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go_server/pkg/setting"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}

	return result
}

func GetOffset(c *gin.Context) int {
	page := 1
	limit := 10

	result := 0

	page, _ = com.StrTo(c.Query("page")).Int()
	limit, _ = com.StrTo(c.Query("limit")).Int()

	if page > 0 {
		result = (page - 1) * limit
	}

	return result
}

func GetLimit(c *gin.Context) int {
	limit := 10
	limit, _ = com.StrTo(c.Query("limit")).Int()
	return limit
}
