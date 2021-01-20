package xdtysfrz

import (
	"bytes"
	"encoding/base64"
	"os"
	"testing"
)

func TestPKCS7Padding(t *testing.T) {
	result := PKCS7Padding([]byte("abc"), 16)

	if !bytes.Equal(result, []byte{97, 98, 99, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13}) {
		t.Error("TestPKCS7Padding")
	}
}

func TestPKCS7UnPadding(t *testing.T) {
	result := PKCS7UnPadding([]byte{97, 98, 99, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13})

	if !bytes.Equal(result, []byte{97, 98, 99}) {
		t.Error("PKCS7UnPadding解填充错误")
	}
}

func TestAESEncryptUsingCBCMode(t *testing.T) {
	result, _ := AESEncryptUsingCBCMode([]byte("abc"), []byte("dddddddddddddddd"), []byte("1111111111111111"))

	resultAfterBase64Encode := base64.StdEncoding.EncodeToString(result)
	theCorrectResult := "vSCLMI8i2I38I/Z0Bm6Vmw=="
	if resultAfterBase64Encode != theCorrectResult {
		t.Error("TestAESEncryptUsingCBCMode加密过程错误")
	}
}

func TestAESDecryptUsingCBCMode(t *testing.T) {
	encryptedData, _ := base64.StdEncoding.DecodeString("vSCLMI8i2I38I/Z0Bm6Vmw==")
	result, _ := AESDecryptUsingCBCMode([]byte(encryptedData), []byte("dddddddddddddddd"), []byte("1111111111111111"))

	theCorrectResult := []byte("abc")
	if !bytes.Equal(result, theCorrectResult) {
		t.Log("AESDecryptUsingCBCMode解密错误")
	}
}

func TestGetARandomByteSliceOfTheSpecifiedLength(t *testing.T) {
	stringLength := 16

	result := getARandomByteSliceOfTheSpecifiedLength(stringLength)
	if len(result) != stringLength {
		t.Error("TestGetARandomStringOfTheSpecifiedLength获取随机字符串函数错误")
		// } else {
		// 	// 方法一没问题：
		// 	t.Log(*(*string)(unsafe.Pointer(&result)))

		// 	// 方法二在测试
		// 	// t.Log(result)
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
