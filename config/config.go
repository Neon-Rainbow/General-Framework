package config

import (
	"flag"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var conf = new(Config)

type Config struct {
	Server Server `mapstructure:"server"`
	Log    Log    `mapstructure:"log"`
}

type Server struct {
	Host    string `mapstructure:"host"`
	Port    string `mapstructure:"port"`
	Mode    string `mapstructure:"mode"`
	Version string `mapstructure:"version"`
}

type Log struct {
	FilePath string `mapstructure:"file_path"`
	Level    int    `mapstructure:"level"`
}

// MustInit 初始化配置文件
// 初始化配置文件出错时，程序会panic
func MustInit() {
	configFilePath := flag.String("config", "etc/config.toml", "配置文件路径")
	flag.Parse()

	viper.SetConfigFile(*configFilePath)

	viper.WatchConfig() // 监听配置文件变化
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件被修改:", in.Name)
	})

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("读取配置文件是出错: %s \n", err))
	}

	if err := viper.Unmarshal(&conf); err != nil {
		panic(fmt.Errorf("配置文件解码错误: %s \n", err))
	}
}

// Get 获取配置信息
func Get() *Config {
	if conf == nil {
		panic("配置文件未初始化")
	}
	return conf
}
