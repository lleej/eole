package config

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

func init() {

}

//Setup 加载配置项
// 1. 命令行
// 2. 环境变量
// 3. 配置文件 settings.yaml
// 4. 默认值
func Setup(path string) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("没有找到配置文件: ", path, err.Error)
	}
}
