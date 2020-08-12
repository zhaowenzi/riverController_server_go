package util

type MenuList struct {
	Name  string `json:"name"`
	Title string `json:"title"`
	Jump  string `json:"jump"`
}

type AdminMenu struct {
	Name  string `json:"name"`
	Title string `json:"title"`
	Icon  string `json:"icon"`
	List  []MenuList `json:"list"`
}

func SuperAdminMenu() []AdminMenu {
	var AdminMenus []AdminMenu

	menu := AdminMenu{Name: "user", Title: "管理员列表", Icon: "layui-icon-user", List: []MenuList{{Name: "administrators-list", Title: "管理员列表", Jump: "user/administrators/list"}}}
	AdminMenus = append(AdminMenus, menu)

	menu = AdminMenu{Name: "news", Title: "新闻管理", Icon: "layui-icon-list", List: []MenuList{{Name: "news-list", Title: "新闻列表", Jump: "news/list"}}}
	AdminMenus = append(AdminMenus, menu)

	menu = AdminMenu{Name: "announcement", Title: "公告管理", Icon: "layui-icon-tabs", List: []MenuList{{Name: "announcement-list", Title: "公告列表", Jump: "announcement/list"}}}
	AdminMenus = append(AdminMenus, menu)

	menu = AdminMenu{Name: "report", Title: "举报管理", Icon: "layui-icon-app", List: []MenuList{{Name: "report-list", Title: "举报列表", Jump: "report/list"}}}
	AdminMenus = append(AdminMenus, menu)

	menu = AdminMenu{Name: "river", Title: "河流管理", Icon: "layui-icon-survey", List: []MenuList{{Name: "river-list", Title: "河流列表", Jump: "river/list"}}}
	AdminMenus = append(AdminMenus, menu)

	menu = AdminMenu{Name: "users", Title: "用户管理", Icon: "layui-icon-login-wechat", List: []MenuList{{Name: "users-list", Title: "用户列表", Jump: "users/list"}, {Name: "users-auth", Title: "授权列表", Jump: "users/auth"}}}
	AdminMenus = append(AdminMenus, menu)

	menu = AdminMenu{Name: "monitors", Title: "数据管理", Icon: "layui-icon-chart", List: []MenuList{{Name: "monitors-map", Title: "一张图", Jump: "monitors/map"}, {Name: "monitors-list", Title: "监测点", Jump: "monitors/list"}}}
	AdminMenus = append(AdminMenus, menu)

	menu = AdminMenu{Name: "water", Title: "水质评价管理", Icon: "layui-icon-chart-screen", List: []MenuList{{Name: "scene-param", Title: "水质评价", Jump: "water/sceneParam"}, {Name: "scene", Title: "水质场景", Jump: "water/scene"}, {Name: "water-report", Title: "水质报告", Jump: "water/report"}}}
	AdminMenus = append(AdminMenus, menu)
	return AdminMenus
}

func NormAdminMenu() []AdminMenu {
	var AdminMenus []AdminMenu

	menu := AdminMenu{Name: "news", Title: "新闻管理", Icon: "layui-icon-list", List: []MenuList{{Name: "news-list", Title: "新闻列表", Jump: "news/listNorm"}}}
	AdminMenus = append(AdminMenus, menu)

	menu = AdminMenu{Name: "announcement", Title: "公告管理", Icon: "layui-icon-tabs", List: []MenuList{{Name: "announcement-list", Title: "公告列表", Jump: "announcement/listNorm"}}}
	AdminMenus = append(AdminMenus, menu)

	menu = AdminMenu{Name: "report", Title: "举报管理", Icon: "layui-icon-app", List: []MenuList{{Name: "report-list", Title: "举报列表", Jump: "report/listNorm"}}}
	AdminMenus = append(AdminMenus, menu)

	menu = AdminMenu{Name: "river", Title: "河流管理", Icon: "layui-icon-survey", List: []MenuList{{Name: "river-list", Title: "河流列表", Jump: "river/list"}}}
	AdminMenus = append(AdminMenus, menu)

	menu = AdminMenu{Name: "users", Title: "用户管理", Icon: "layui-icon-login-wechat", List: []MenuList{{Name: "users-list", Title: "用户列表", Jump: "users/listNorm"}, {Name: "users-auth", Title: "授权列表", Jump: "users/auth"}}}
	AdminMenus = append(AdminMenus, menu)

	menu = AdminMenu{Name: "monitors", Title: "数据管理", Icon: "layui-icon-chart", List: []MenuList{{Name: "monitors-map", Title: "一张图", Jump: "monitors/map"}, {Name: "monitors-list", Title: "监测点", Jump: "monitors/list"}}}
	AdminMenus = append(AdminMenus, menu)
	return AdminMenus
}
