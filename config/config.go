package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

// Mysql 配置
type Mysql struct {
	Host     string
	User     string
	Password string
	DB       string
	Port     string
}


// Config 对应config.yml文件的位置
type Config struct {
	//Debug      bool `toml:"debug"`
	//Port       string
	//Secret     string
	//JobWorkers int    `toml:"job_workers"`
	//JwtSecret  string `toml:"jwt_secret"`
	Mysql      `toml:"mysql"`
}

// config
var config Config

// 配置文件路径
var configFile = ""

// Get 获取config
func Get() Config {
	if config.Host == "" {
		// 默认配置文件在同级目录
		filepath := getPath()

		// 解析配置文件
		if _, err := toml.DecodeFile(filepath, &config); err != nil {
			log.Fatal("配置文件读取失败！", err)
		}
	}

	return config
}

// SetPath 设置Config文件的路径
func SetPath(path string) {
	configFile = path
}

// 获取文件路径
func getPath() string {

	// 默认配置文件在同级目录

	path, _ := os.Getwd()
	filepath := path + "/config/config.toml"

	return filepath
}
