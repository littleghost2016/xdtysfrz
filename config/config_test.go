package config

import (
	"fmt"
	"testing"
)

func TestGetConfigFromTomlFile(t *testing.T) {
	config := GetConfigFromTomlFile("./config.toml")

	// 此处如果直接使用config["login"]["username"]会显示
	// GetConfigFromTomlFile("./config.toml").login undefined (type map[string]interface {} has no field or method login)
	// 因为此处config["login"]为interface{}而不是map所以不能直接索引
	// 因此应该使用断言,将config["login"]变为map[string]interface{}后,才能继续使用键去索引
	loginConfig, ok := config["login"].(map[string]interface{})
	if !ok {
		t.Error("断言出现问题")
	}
	fmt.Println(loginConfig["username"])
}
