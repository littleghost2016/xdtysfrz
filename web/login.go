package web

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	myCrypto "xdtysfrz/crypto"

	"github.com/antchfx/htmlquery"
)

// LoginAndGetCookie 登录并获取cookie
func LoginAndGetCookie(myMethod, myURL, myUsername, myPassword string) (myCookie string) {

	// 创建复用的客户端和cookiejar
	// 初始化cookie池
	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: cookieJar,
	}

	// 获取所需参数
	requiredParametersForLogin := getRequiredParametersForLogin(client, "http://ids.xidian.edu.cn/authserver/login")

	// TODO: 检查是否需要填写验证码
	// 使用密码登录错误三次后，将需要输入验证码，对应的js代码为，其中ts与时间有关，且猜测验证码将根据用户名和时间生成
	// if (dllt != '' && dllt == 'dynamicLogin') {
	// 	}else{
	// 		$("#captchaImg").attr("src","captcha.html?ts=" + new Date().getMilliseconds());
	// 	}
	// captchaCheckFlag := captchaCheck(client, myUsername)

	// 准备表单内容
	formData := make(url.Values)
	keysThatWorked := []string{"lt", "dllt", "execution", "_eventId", "rmShown"}
	for _, eachKey := range keysThatWorked {
		// 设置每一项
		formData.Set(eachKey, requiredParametersForLogin[eachKey])
	}
	// 以下一行仅供测试pwdDefaultEncryptSalt
	// fmt.Println(requiredParametersForLogin["pwdDefaultEncryptSalt"])
	formData.Set("username", myUsername)
	encryptedPassword := myCrypto.EncryptPassword(myPassword, requiredParametersForLogin["pwdDefaultEncryptSalt"], "1111111111111111")
	// 测出的iv本应该如下一行所示使用随机字符串，但对于统一身份认证系统来说，iv值是多少不重要
	// encryptedPassword := myCrypto.EncryptPassword(myPassword, requiredParametersForLogin["pwdDefaultEncryptSalt"], string(myCrypto.GetARandomByteSliceOfTheSpecifiedLength(16)))
	encryptedPasswordAfterBase64 := base64.StdEncoding.EncodeToString(encryptedPassword)
	formData.Set("password", encryptedPasswordAfterBase64)
	formDataEncoded := formData.Encode()
	// 因为EncryptPassword返回的内容里面没有对+、=等符号进行url编码，在上一行的Encode()之后被正常urlencode
	// 注意：当=等符号被转换成%3D等符号以后，再进行Encode()则会对%再次进行编码。
	// 因此若不想让+被转换成%3D等情况，可如下一行所示使用url.QueryUnescape()逆行编码
	// formDataUnescaped, _ := url.QueryUnescape(formData.Encode())
	// 以下一行仅供测试formDataUnescaped
	// fmt.Printf("formData.Encode() %s\n", formDataUnescaped)

	// 创建新请求
	postRequest, _ := http.NewRequest("POST", myURL, strings.NewReader(formDataEncoded))
	setLoginHeader(postRequest)
	// fmt.Println("request.Body", request.Body)

	response, err := client.Do(postRequest)
	if err != nil {
		panic("POST登录发送请求出错")
	}

	defer response.Body.Close()

	// 以下一行仅供测试response.StatusCode
	fmt.Println("response.StatusCode", response.StatusCode)

	// 以下一行仅供测试cookie
	// fmt.Println(response.Cookies())

	// 当使用如下两行时，程序会卡在网络通信的过程中停止，在wireshark中表现为客户端强制发送RST标志以终止连接
	// 猜测为string函数遇到了无法转换的字符导致崩溃，进而客户端终止了链接
	// b, _ := ioutil.ReadAll(response.Body)
	// fmt.Println(string(b))

	// TODO: 继续完成其他功能
	return ""
}

func getRequiredParametersForLogin(client *http.Client, loginURL string) (parameters map[string]string) {

	// 初始化请求
	request, _ := http.NewRequest("GET", loginURL, nil)

	response, err := client.Do(request)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	// 以下一行供测试请求返回状态码，200为正常
	// fmt.Printf("httpResp.Status=%s", resp.Status)

	// b, _ := ioutil.ReadAll(response.Body)
	// return string(b)

	parameters = make(map[string]string, 6)

	doc, _ := htmlquery.Parse(response.Body)
	formInputSlice := htmlquery.Find(doc, "//input[@type='hidden']")

	// 以下一行仅供测试formInputSlice的长度
	// fmt.Println("len(formInputSlice)", len(formInputSlice))

	for _, eachInput := range formInputSlice {
		var tempKey string
		tempValue := htmlquery.SelectAttr(eachInput, "value")

		if htmlquery.SelectAttr(eachInput, "name") != "" {
			tempKey = htmlquery.SelectAttr(eachInput, "name")
		} else if htmlquery.SelectAttr(eachInput, "id") != "" {
			tempKey = htmlquery.SelectAttr(eachInput, "id")
		}

		parameters[tempKey] = tempValue
	}

	// 以下一行仅供测试formInputMap的内容
	// fmt.Println(formInputMap)

	return
}
