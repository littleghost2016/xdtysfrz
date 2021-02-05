package web

import (
	"fmt"
	"testing"

	myConfig "xdtysfrz/config"
)

func TestLoginAndGetCookie(t *testing.T) {
	url := "http://ids.xidian.edu.cn/authserver/login"

	config := myConfig.GetConfigFromTomlFile("../config/config.toml")
	loginConfig, ok := config["login"].(map[string]interface{})
	if !ok {
		t.Error("断言出现问题")
	}

	result := LoginAndGetCookie("POST", url, loginConfig["username"].(string), loginConfig["password"].(string))
	fmt.Println(result)
}

// func TestGetRequiredParametersForLogin(t *testing.T) {
// 	client := &http.Client{}
// 	url := "http://ids.xidian.edu.cn/authserver/login"

// 	// 获取登陆所需参数
// 	parameters := getRequiredParametersForLogin(client, url)

// 	// 检验参数数量是否正确
// 	if len(parameters) != 7 {
// 		t.Error("获取到的参数数量不是7")
// 	}

// 	// 检验参数名称是否符合
// 	// 以下为所需要包含的参数名称
// 	keysThatShouldBeIncluded := []string{
// 		"_eventId",
// 		"dllt",
// 		"execution",
// 		"isSliderCaptcha",
// 		"lt",
// 		"pwdDefaultEncryptSalt",
// 		"rmShown",
// 	}
// 	// 循环验证
// 	for _, eachKey := range keysThatShouldBeIncluded {
// 		if _, ok := parameters[eachKey]; !ok {
// 			t.Error("登陆所需参数中不包含：", eachKey)
// 		}
// 	}
// }

// func TestLoginAndGetCookie(t *testing.T) {
// 	LoginAndGetCookie()
// }
