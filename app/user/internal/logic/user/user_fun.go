package user

import "github.com/gogf/gf/v2/crypto/gmd5"

func EncryptPwd(password string) (string, error) {
	return gmd5.EncryptString(password)
}
