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

func AnnouncementList(c *gin.Context) {
	title := c.DefaultQuery("title", "");
	data, count := models.GetAnnouncementAndCount(util.GetOffset(c), util.GetLimit(c), title)
	code := e.OK
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   e.GetMsg(code),
		"data":  data,
		"count": count,
	})
}

func EditAnnouncementStatus(c *gin.Context) {
	id := c.PostForm("id")
	status := c.PostForm("status")

	valid := validation.Validation{}
	valid.Required(id, "id").Message("id不能为空")
	valid.Required(status, "staus").Message("status不能为空")

	code := e.PARAM_ERROR
	var err error
	if !valid.HasErrors() {
		err = models.EditAnnouncementStatus(com.StrTo(id).MustInt(), com.StrTo(status).MustInt())
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

func DelAnnouncement(c *gin.Context) {
	id := c.PostForm("id")
	valid := validation.Validation{}
	valid.Required(id, "id").Message("id不能为空")

	code := e.PARAM_ERROR
	var err error
	if !valid.HasErrors() {
		err = models.DelAnnouncement(com.StrTo(id).MustInt())
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

func ShowAnnouncementContent(c *gin.Context) {
	id := c.Query("id")
	code := e.OK
	data := models.GetAnnouncement(com.StrTo(id).MustInt())
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func EditAnnouncement(c *gin.Context) {
	id := com.StrTo(c.PostForm("id")).MustInt()
	title := c.PostForm("title")
	source := c.PostForm("source")
	content := c.PostForm("content")

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.Required(title, "title").Message("title不能为空")
	valid.Required(source, "source").Message("source不能为空")
	valid.Required(content, "content").Message("content不能为空")

	code := e.PARAM_ERROR
	var err error
	if !valid.HasErrors() {
		data := make(map[string]interface{})
		data["title"] = title
		data["source"] = source
		data["content"] = content
		data["update_time"] = util.GenerateCurrentTimeStamp()
		err = models.EditAnnouncement(id, data)
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

func AddAnnouncement(c *gin.Context) {
	title := c.PostForm("title")
	source := c.PostForm("source")
	content := c.PostForm("content")

	valid := validation.Validation{}
	valid.Required(title, "title").Message("title不能为空")
	valid.Required(source, "source").Message("source不能为空")
	valid.Required(content, "content").Message("content不能为空")

	accessToken := c.Query("access_token")
	claims, _ := util.ParseToken(accessToken)

	code := e.PARAM_ERROR
	var err error
	if !valid.HasErrors() {
		data := make(map[string]interface{})
		data["uid"] = claims.Uid
		data["title"] = title
		data["source"] = source
		data["content"] = content
		data["create_time"] = util.GenerateCurrentTimeStamp()
		data["update_time"] = util.GenerateCurrentTimeStamp()
		data["release_time"] = int64(0)
		data["status"] = 0
		models.AddAnnouncement(data)
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