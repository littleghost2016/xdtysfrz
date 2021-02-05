package config

import (
	"github.com/BurntSushi/toml"
)

// GetConfigFromTomlFile 从指定位置的toml文件获取配置信息
func GetConfigFromTomlFile(filePath string) (config map[string]interface{}) {
	// BurntSushi/toml可以将toml文件转为struct或者map,以前使用的是struct,每次更改结构体比较麻烦
	if _, err := toml.DecodeFile(filePath, &config); err != nil {
		panic(err)
	}

	return
}
