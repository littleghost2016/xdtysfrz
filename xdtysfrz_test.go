package xdtysfrz

import (
	"bytes"
	"encoding/base64"
	"testing"
)

func TestEncryptAndDecrypt(t *testing.T) {
	password := "123456"
	key := "dddddddddddddddd"
	iv := "1111111111111111"

	encryptedPassword := GetEncryptedPassword(password, key, iv)
	encryptedPasswordAfterBase64 := base64.StdEncoding.EncodeToString(encryptedPassword)
	t.Log("encryptedPasswordAfterBase64: ", encryptedPasswordAfterBase64)

	decryptedPassword := DecryptPassword(encryptedPassword, key, iv)
	if !bytes.Equal(decryptedPassword[64:], []byte(password)) {
		t.Log(decryptedPassword[64:])
		t.Error("EncryptAndDecrypt过程出错")
	}
}
