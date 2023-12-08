// Package ccfg
// @Author cuisi
// @Date 2023/10/30 18:01:00
// @Desc config
package ccfg

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	cfgName string
	cfgType string
	cfgPath string
}

var c *Config

func init() {
	c = New()
}

func New() *Config {
	c := new(Config)
	return c

}

// SetCfgName sets name for the config file.
// Does not include extension.
func SetCfgName(in string) {
	c.SetCfgName(in)
}

func (c *Config) SetCfgName(in string) {
	if in != "" {
		c.cfgName = in
	}
}

// SetCfgType sets the type of the configuration returned by the
// remote source, e.g. "json".
func SetCfgType(in string) {
	c.SetCfgType(in)
}
func (c *Config) SetCfgType(in string) {
	if in != "" {
		c.cfgType = in
	}
}

// SetCfgPath adds a path for Viper to search for the config file in.
// Can be called multiple times to define multiple search paths.
func SetCfgPath(in string) {
	c.SetCfgPath(in)
}
func (c *Config) SetCfgPath(in string) {
	if in != "" {
		c.cfgPath = in
	}
}

// Get returns an interface. For a specific value use one of the Get____ methods.
func Get(key string) any {
	return viper.Get(key)
}

// ReadCfg will discover and load the configuration file from disk
// and key/value stores, searching in one of the defined paths.
func Install() {
	// 设置配置文件的名字
	viper.SetConfigName(c.cfgName)
	// 设置配置文件的类型
	viper.SetConfigType(c.cfgType)
	// 添加配置文件的路径，指定 config 目录下寻找
	viper.AddConfigPath(c.cfgPath)
	// 寻找配置文件并读取
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// fmt.Println(viper.Get("mysql"))     // map[port:3306 url:127.0.0.1]
	// fmt.Println(viper.Get("mysql.url")) // 127.0.0.1
}
