package util

import (
	"encoding/base64"
)

// 加密密码
func HashAndSalt(email, pwd []byte) string {
	//hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	//if err != nil {
	//	log.Println(err)
	//}
	//return string(hash)

	//return encode(string(email), string(pwd))

	return string(pwd)
}

//验证密码
func ComparePasswords(email []byte, plainPwd []byte, hashedPwd string) bool {

	//byteHash := []byte(hashedPwd)
	//err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	//if err != nil {
	//	log.Println(err)
	//	return false
	//}
	//return true

	//return HashAndSalt(email, plainPwd) == hashedPwd

	return string(plainPwd) == hashedPwd
}

func encode(email, password string) string {
	base := email + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(base))
}
