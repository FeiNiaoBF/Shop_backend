package utility

import "github.com/gogf/gf/v2/crypto/gmd5"

// EncryptPassword 为admin密码加密
func EncryptPassword(password, salt string) string {
	return gmd5.MustEncryptString(gmd5.MustEncryptString(password) + gmd5.MustEncryptString(salt))
}
