package frontend

import "github.com/gogf/gf/v2/frame/g"

type RegisterReq struct {
	g.Meta       `path:"/register" method:"post" tags:"用户" sm:"用户注册接口"`
	Name         string `json:"name" description:"用户名" v:"required#用户名必填"`
	Password     string `json:"password" description:"用户密码" v:"password"`
	Avatar       string `json:"avatar" description:"头像"`
	UserSalt     string `json:"user_salt" description:"加密盐 生成密码用"`
	Sex          int    `json:"sex" description:"1男 2女"`
	Status       int    `json:"status" description:"1正常 2黑名单"`
	Sign         string `json:"sign" description:"个性签名" v:"max-length:30#个性签名最大30个字"`
	SecretAnswer string `json:"secret_answer"  description:"密保问题的答案"`
}

type RegisterRes struct {
	Id uint `json:"id"`
}
