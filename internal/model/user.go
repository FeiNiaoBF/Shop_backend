package model

import "github.com/gogf/gf/v2/frame/g"

type RegisterInput struct {
	Name         string `json:"name" dc:"用户名"`
	Password     string `json:"password" dc:"用户密码"`
	Avatar       string `json:"avatar" dc:"头像"`
	UserSalt     string `json:"user_salt" dc:"加密盐 生成密码用"`
	Sex          int    `json:"sex" dc:"1男 2女"`
	Status       int    `json:"status" dc:"1正常 2黑名单"`
	Sign         string `json:"sign" dc:"个性签名"`
	SecretAnswer string `json:"secret_answer"  dc:"密保问题的答案"`
}

type RegisterOutput struct {
	Id uint `json:"id"`
}

type LoginInput struct {
	Name     string `json:"name"         description:"用户名" v:"required#用户名必填"`
	Password string `json:"password"     description:"密码" v:"password"`
}

type LoginOutput struct {
	Type     string `json:"type"`
	Token    string `json:"token"`
	ExpireIn int    `json:"expire_in"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Sex      uint8  `json:"sex"`
	Sign     string `json:"sign"`
	Status   uint8  `json:"status"`
}

//type InfoInput struct {
//}
//
//type InfoOutput struct {
//	Id     uint   `json:"id"`
//	Name   string `json:"name"`
//	Avatar string `json:"avatar"`
//	Sex    uint8  `json:"sex"`
//	Sign   string `json:"sign"`
//	Status uint8  `json:"status"`
//}

type UpdatePasswordInput struct {
	Password     string `json:"password"  v:"password"   description:"重置密码"`
	UserSalt     string `json:"userSalt,omitempty"     description:"加密盐 生成密码用"`
	SecretAnswer string `json:"secretAnswer" description:"密保问题的答案"`
}

type UpdatePasswordOutput struct {
	Id uint `json:"id"`
}

type UserInfoBase struct {
	g.Meta `orm:"table:user_info"`
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Sex    uint8  `json:"sex"`
	Sign   string `json:"sign"`
	Status uint8  `json:"status"`
}
