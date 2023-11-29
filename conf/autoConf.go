package conf

import (
	"GoCloud/pkg/log"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
	"sync"
)

var (
	config     = viper.New()
	onceConfig sync.Once
	Change     chan struct{}
)

func Config() *viper.Viper {
	onceConfig.Do(
		func() {
			initConfig()
		})
	return config
}

// AddPath 此方法专门供测试函数使用,可以用于临时增加配置文件的路径
func AddPath(path string) {
	log.NewEntry("conf").Info("add path %s to config" + path)
	config.AddConfigPath(path)
}
func initConfig() {
	log.NewEntry("conf").Info("init config")
	config.MergeConfigMap(defaultConfig)
	configDir := "./"
	configName := "config"
	configType := "yaml"
	//设置配置文件路径

	//将默认值设置到config中
	config.AddConfigPath(configDir)
	config.SetConfigName(configName)
	config.SetConfigType(configType)

	// 配置文件出错
	if err := config.ReadInConfig(); err != nil {
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
			if err := config.WriteConfigAs(configPath); err != nil {
				panic(errors.Wrapf(err, "[error] Failed to generate default config file. %s\n", configPath))
			}
			// 再次读取配置文件
			if err := config.ReadInConfig(); err != nil {
				panic(errors.Wrapf(err, "[error] Failed to read config file. %s\n", configPath))
			}
			panic("请修改配置文件后重启程序")
		}
	}
	//再次合并，为了覆盖默认值
	config.WatchConfig()
	config.OnConfigChange(func(e fsnotify.Event) {
		ReloadConfig()
	})
}
