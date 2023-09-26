package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
	"sync"
)

var (
	autoConfig     *viper.Viper
	onceAutoConfig sync.Once
)

func Config() *viper.Viper {
	onceAutoConfig.Do(
		func() {
			initConfig()
		})
	return autoConfig
}

// AddPath 此方法专门供测试函数使用,可以用于临时增加配置文件的路径
func AddPath(path string) {
	autoConfig.AddConfigPath(path)
}

func initConfig() {
	autoConfig.MergeConfigMap(defaultConfig)
	configDir := "./config/"
	configName := "config"
	configType := "yaml"
	//设置配置文件路径

	autoConfig.AddConfigPath(configDir)
	//将默认值设置到config中
	autoConfig.AddConfigPath(configDir)
	autoConfig.SetConfigName(configName)
	autoConfig.SetConfigType(configType)

	// 配置文件出错

	if err := autoConfig.ReadInConfig(); err != nil {
		// 如果找不到配置文件，则提醒生成配置文件并创建它
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			// 如果 config 目录不存在，则创建它
			if _, err := os.Stat(configDir); os.IsNotExist(err) {
				if err = os.MkdirAll(configDir, 0755); err != nil {
					panic(err)
				}
			}
			configPath := configDir + configName + "." + configType
			fmt.Println(errors.Wrapf(err, "[warning] Config file not found. Generating default config file at %s\n", configPath))
			if err := autoConfig.WriteConfigAs(configPath); err != nil {
				panic(errors.Wrapf(err, "[error] Failed to generate default config file. %s\n", configPath))
			}
			// 再次读取配置文件
			if err := autoConfig.ReadInConfig(); err != nil {
				panic(errors.Wrapf(err, "[error] Failed to read config file. %s\n", configPath))
			}
			panic("请修改配置文件后重启程序")
		}
	}
	autoConfig.WatchConfig()
	autoConfig.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:")
		ReloadConfig()
	})
}
func ReloadConfig() {
	// 配置文件发生变更之后会调用的回调函数

}
