package model

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
