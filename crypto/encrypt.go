package crypto

// EncryptPassword 加密密码口令
func EncryptPassword(password, key, iv string) (encryptedPassword []byte) {

	// 生成所需要的前64位随机字符串
	randomByteSlice1 := GetARandomByteSliceOfTheSpecifiedLength(64)
	constructedInput := append(randomByteSlice1, []byte(password)...)

	encryptedPassword, _ = AESEncryptUsingCBCMode(constructedInput, []byte(key), []byte(iv))
	return
}
