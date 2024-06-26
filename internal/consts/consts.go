package consts

// Project config
const (
	ProjectName  = "Go开源电商实战项目"
	ProjectUsage = "Shop"
	ProjectBrief = "start http server"
)

const (
	Version            = "v0.0.1"             // 版本号
	CaptchaDefaultName = "CaptchaDefaultName" // 验证码默认存储空间名称
)

// Admin
const (
	ContextKey               = "ContextKey" // 上下文变量存储键名，前后端系统共享
	GTokenAdminPrefix        = "Admin:"     //gtoken登录 管理后台 前缀区分
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
	CtxUserName      = "CtxUserName"
	CtxUserAvatar    = "CtxUserAvatar"
	CtxUserSex       = "CtxUserSex"
	CtxUserSign      = "CtxUserSign"
	CtxUserStatus    = "CtxUserStatus"
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
	ErrNullPassword    = "账号或密码为空."
	ErrNullUser        = "没有该账号."
	ErrPassword        = "账号或密码错误."
	ErrSecretAnswerMsg = "密保问题不正确"
)

// 收藏相关
const (
	CollectionTypeGoods = iota + 1
	CollectionTypeArticle
)

// 点赞相关
const (
	PraiseTypeGoods = iota + 1
	PraiseTypeArticle
)

// 评论相关
const (
	CommentTypeGoods = iota + 1
	CommentTypeArticle
)
