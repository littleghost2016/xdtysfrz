package xdtysfrz

func GetEncryptedPassword(password, key, iv string) (encryptedPassword []byte) {
	randomByteSlice1 := getARandomByteSliceOfTheSpecifiedLength(64)
	constructedInput := append(randomByteSlice1, []byte(password)...)

	encryptedPassword, _ = AESEncryptUsingCBCMode(constructedInput, []byte(key), []byte(iv))
	return
}

func DecryptPassword(cryptedPassword []byte, key, iv string) (password []byte) {
	password, _ = AESDecryptUsingCBCMode(cryptedPassword, []byte(key), []byte(iv))

	return
}
