package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"math/rand"
	"time"
)

// PKCS7Padding 使用PKCS7进行填充，填充过程是在加密之前
// PKCS7的填充规则为，还差几个字节能达到blockSize的大小，则填充的每个字节为这个数字
// 例：[97, 98, 99]需要填充到16字节还差13个字节，则填充后的内容为[97 98 99 13 13 13 13 13 13 13 13 13 13 13 13 13]
// 输入：未填充未加密的原始数据，数据块大小
// 输出：填充过的未加密数据
func PKCS7Padding(uncryptedData []byte, blockSize int) (paddedUncryptedData []byte) {
	paddingLength := blockSize - (len(uncryptedData) % blockSize)

	paddingText := bytes.Repeat([]byte{byte(paddingLength)}, paddingLength)
	paddedUncryptedData = append(uncryptedData, paddingText...)

	return
}

// PKCS7UnPadding 去除PKCS7填充
// 读取最后一个字节，去除掉对应长度的填充
// 输入：填充过的未加密数据
// 输出：未填充未加密的原始数据
func PKCS7UnPadding(paddedUncryptedData []byte) (unpaddedUncryptedData []byte) {
	paddedUncryptedDataLength := len(paddedUncryptedData)
	paddingLength := int(paddedUncryptedData[paddedUncryptedDataLength-1])

	unpaddedUncryptedData = paddedUncryptedData[:(paddedUncryptedDataLength - paddingLength)]

	return
}

// AESEncryptUsingCBCMode AES加密
// 输入：原始数据（明文），密钥
// 输出：加密后的数据（密文），错误
func AESEncryptUsingCBCMode(rawData, key, iv []byte) (cryptedData []byte, functionError error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	//填充原始数据
	blockSize := block.BlockSize()
	paddedUncryptedData := PKCS7Padding(rawData, blockSize)

	cryptedDataLength := len(paddedUncryptedData)
	cryptedData = make([]byte, cryptedDataLength)

	// if _, err := io.ReadFull(rand.Reader, iv); err != nil {
	// 	panic(err)
	// }

	// 初始向量iv在加密的时候很重要，但因为逆向分析以后发现在统一身份认证系统中iv的值并不影响结果，所以iv可以随便取值
	// 关于初始向量的分析可参考readme.md，其中IV的影响参考自博客 https://www.jianshu.com/p/45848dd484a9
	// 最开始用的是下面这一行设置iv，其值为[1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1]
	// 即int类型的1直接转换成byte类型的1，然后重复16遍，得到的加密后的内容与在线AES https://oktools.net/aes 算出来的有所不同
	// iv = bytes.Repeat([]byte{byte(1)}, blockSize)
	// 然后使用了下面这一行，直接将string转成[]byte，得到的结果与在线加密一致...
	// iv = []byte("1111111111111111")
	// 下面一行的输出仅作iv值的输出测试，与函数功能无关
	// fmt.Println("iv", iv)
	// 但是！实际使用的iv是传进来的

	//block大小和初始向量大小一定要一致
	mode := cipher.NewCBCEncrypter(block, iv)

	// 下面三行得输出仅作cryptedData的输出测试，与函数功能无关
	// fmt.Println("paddedUncryptedData", paddedUncryptedData, len(paddedUncryptedData))
	// fmt.Println("cryptedData", cryptedData, len(cryptedData))
	// fmt.Println("cryptedData[blockSize:]", cryptedData[blockSize:], len(cryptedData[blockSize:]))

	mode.CryptBlocks(cryptedData, paddedUncryptedData)

	functionError = nil

	return
}

// AESDecryptUsingCBCMode AES解密
func AESDecryptUsingCBCMode(encryptedData, key, iv []byte) (unpaddedUncryptedData []byte, functionError error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	blockSize := block.BlockSize()
	if len(encryptedData) < blockSize {
		panic("encryptedData too short")
	}

	// CBC mode always works in whole blocks.
	if len(encryptedData)%blockSize != 0 {
		panic("encryptedData is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	paddedUncryptedData := make([]byte, len(encryptedData))
	mode.CryptBlocks(paddedUncryptedData, encryptedData)
	//解填充
	unpaddedUncryptedData = PKCS7UnPadding(paddedUncryptedData)

	functionError = nil

	return
}

// GetARandomByteSliceOfTheSpecifiedLength 获取随机字节切片
func GetARandomByteSliceOfTheSpecifiedLength(randomByteSliceLength int) (randomByteSliceOfTheSpecifiedLength []byte) {
	// 方法一：此处随机字符串的获取方式参考自 https://zhuanlan.zhihu.com/p/90830253
	src := rand.NewSource(time.Now().UnixNano())
	const initialString = "ABCDEFGHJKMNPQRSTWXYZabcdefhijkmnprstwxyz2345678"
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)

	randomByteSliceOfTheSpecifiedLength = make([]byte, randomByteSliceLength)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := randomByteSliceLength-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(initialString) {
			randomByteSliceOfTheSpecifiedLength[i] = initialString[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	// // 方法二：
	// randomByteSliceOfTheSpecifiedLength = make([]byte, randomByteSliceLength)
	// if _, err := io.ReadFull(crypto_rand.Reader, randomByteSliceOfTheSpecifiedLength); err != nil {
	// 	panic(err)
	// }

	return
}
