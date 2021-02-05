package xdtysfrz

import (
	myCrypto "xdtysfrz/crypto"
)

// GetEncryptedPassword 获取加密密码口令
func GetEncryptedPassword(password, key, iv string) (encryptedPassword []byte) {
	encryptedPassword = myCrypto.EncryptPassword(password, key, iv)

	return
}

// GetDecryptedPassword 获取解密后的密码口令
func GetDecryptedPassword(cryptedPassword []byte, key, iv string) (password []byte) {
	password = myCrypto.DecryptPassword(cryptedPassword, key, iv)

	return
}

// GetCookie 获取模拟登录后的cookie
// func GetCookie(username, password string) {

// }
