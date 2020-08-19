package routers

import (
	"github.com/gin-gonic/gin"
	jwt "go_server/middleware"
	"go_server/pkg/setting"
	"go_server/routers/api"
	v1 "go_server/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.POST("/login", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)

		// 管理员列表
		apiv1.GET("/adminUsers", jwt.CheckSuperAdmin(), v1.AdminUsers)
		// 新增管理员
		apiv1.POST("/addAdmin", jwt.CheckSuperAdmin(), v1.AddAdmin)
		// 删除管理员
		apiv1.POST("/delAdmin", jwt.CheckSuperAdmin(), v1.DelAdmin)
		// 编辑管理员
		apiv1.POST("/editAdmin", jwt.CheckSuperAdmin(), v1.EditAdmin)
		// 管理员菜单
		apiv1.GET("/adminMenu", v1.AdminMenu)

		// 公告列表
		apiv1.GET("/announcementList", v1.AnnouncementList)
		// 更改公告状态
		apiv1.POST("/editAnnouncementStatus", v1.EditAnnouncementStatus)
		// 删除公告
		apiv1.POST("/delAnnouncement", v1.DelAnnouncement)
		// 单个公告
		apiv1.GET("/showAnnouncementContent", v1.ShowAnnouncementContent)
		// 编辑公告
		apiv1.POST("/editAnnouncement", v1.EditAnnouncement)
		// 新增公告
		apiv1.POST("/addAnnouncement", v1.AddAnnouncement)

		// 新闻列表
		apiv1.GET("/newsList", v1.NewList)
		// 更改新闻状态
		apiv1.POST("/editNewStatus", v1.EditNewStatus)
		// 删除新闻
		apiv1.POST("/delNew", v1.DelNew)
		// 单个新闻
		apiv1.GET("/showNewContent", v1.ShowNewContent)
		// 编辑新闻
		apiv1.POST("/editNew", v1.EditNew)
		// 新增新闻
		apiv1.POST("/newNew", v1.AddNew)
		// TODO 上传图片接口
		// TODO 导入新闻列表
		// TODO 导入新闻具体信息

		// 监测点列表
		apiv1.GET("/monitorsList", v1.MonitorList)
		// 新增监测点
		apiv1.POST("/addMonitor", v1.AddMonitor)
		// 编辑监测点
		apiv1.POST("/editMonitor", v1.EditMonitor)
		// 删除监测点
		apiv1.POST("/delMonitor", v1.DelMonitor)
		// TODO 获取地图上的监测点

		// 河流列表
		apiv1.GET("/riverList", v1.RiverList)
		// 新增河流
		apiv1.POST("/addRiver", v1.AddRiver)
		// 删除河流
		apiv1.POST("/delRiver", v1.DelRiver)
		// 编辑河流
		apiv1.POST("/editRiver", v1.EditRiver)
		// 获取河流授权列表
		apiv1.GET("/authList", v1.AuthList)
		// TODO authRiver
		apiv1.GET("/allRiverParser", v1.AllRiverParser)

		// 用户列表
		apiv1.GET("/userList", v1.UserList)
		// 删除用户
		apiv1.POST("/delUser", v1.DelUser)
		// 巡河记录
		apiv1.GET("/patrolList", v1.PatrolList)
		// 获取申请授权
		apiv1.GET("/authUser", v1.AuthUser)
		// 对用户授权
		apiv1.POST("/authUserRiver",v1.AuthUserRiver)

		// 举报列表
		apiv1.GET("/reportList", v1.ReportList)
		// 获取举报内容
		apiv1.GET("/getReportContent", v1.GetReportContent)
		// 获取所有可以分配的用户
		apiv1.GET("/getDelegate", v1.GetDelegate)
		// 分配负责人
		apiv1.POST("/delegateUser", v1.DelegateUser)
		// 审核举报
		apiv1.POST("/audit", v1.Audit)

		// 获取场景参数
		apiv1.GET("/sceneParamParser", v1.SceneParamParser)
		// 获取水质评价标准
		apiv1.GET("/waterParam", v1.WaterParamParser)
		// 修改水质评价标准
		apiv1.POST("/editSceneParam", v1.EditSceneParam)
		// 修改水质场景
		apiv1.POST("/editScene", v1.EditScene)

	}

	return r
}