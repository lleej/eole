package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func init() {

}

//Setup 加载配置项
// 1. 命令行
// 2. 环境变量
// 3. 配置文件 settings.yaml
// 4. 默认值
func Setup(path string) {
	viper.SetConfigFile()
}