package consts

// Admin
const (
	Version                  = "v0.0.1"             // 版本号
	CaptchaDefaultName       = "CaptchaDefaultName" // 验证码默认存储空间名称
	ContextKey               = "ContextKey"         // 上下文变量存储键名，前后端系统共享
	GTokenAdminPrefix        = "Admin:"             //gtoken登录 管理后台 前缀区分
	CtxAdminId               = "CtxAdminId"
	CtxAdminName             = "CtxAdminName"
	CtxAdminIsAdmin          = "CtxAdminIsAdmin"
	CtxAdminRoleIds          = "CtxAdminRoleIds"
	FileMaxUploadCountMinute = 10 // 最大文件上传量
)

// User
const (
	GTokenUserPrefix = "User:" //gtoken登录 管理后台 前缀区分
	CtxUserId        = "CtxUserId"
	CtxUserName      = "CtxAUserName"
	CtxUserRoleIds   = "CtxUserRoleIds"
)

// Project config
const (
	ProjectName  = "Go开源电商实战项目"
	ProjectUsage = "Shop"
	ProjectBrief = "start http server"
)

// Gtoken config
const (
	CacheMod           = 2
	ServerName         = "shop"
	FrontendMultiLogin = false
	BackendMultiLogin  = true
	TokenType          = "Bearer"
	TokenExpireIn      = 10 * 24 * 60 * 60 //单位秒,

	// Error
	ErrNullPassword = "账号或密码为空."
	ErrNullUser     = "没有该账号."
	ErrPassword     = "账号或密码错误."
)
