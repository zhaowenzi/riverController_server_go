package e

var MsgFlags = map[int]string {
	OK: "ok",
	DB_ERROR: "数据库错误",
	PARAM_ERROR: "请求参数错误",
	AUTHORIZATION_ERROR: "认证授权错误",
	UNKNOWN_ERROR: "未知错误",
	ACC_PASS_ERROR: "账号或密码错误",
	ADMIN_DUMP: "管理员账号重复",
	KEY_DUMP: "键值重复",
	API_LOGIN_NO_PHONE: "没有登录权限，使用第三方登录",
	CODE_ERROR: "验证码错误",
	NOT_PERSON: "不是指定负责人",
	NOT_HARDWARE: "没有该设备的信息",
	NOT_ADMIN: "该用户没有权限",
	PEND_APPLY: "该用户已经申请等待通过",
	ALREADY_ALLPY: "该用户已有权限",
	NO_NAMEORPHONE: "没有填写姓名或手机号",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[UNKNOWN_ERROR]
}