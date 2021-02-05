package crypto

// DecryptPassword 解密乱码口令
func DecryptPassword(cryptedPassword []byte, key, iv string) (password []byte) {
	passwordWithRandomByteSlice, _ := AESDecryptUsingCBCMode(cryptedPassword, []byte(key), []byte(iv))

	// 去掉前64位随机字符串
	password = passwordWithRandomByteSlice[64:]

	return
}
