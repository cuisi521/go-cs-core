// Package ccfg
// @Author cuisi
// @Date 2023/10/30 18:01:00
// @Desc config
package ccfg

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	defaultDir  = "./manifest/config/setting.yaml"
	defaultName = "setting.yaml"
	defaultType = "yaml"
)

type CnfYaml struct {
	cfgName  string
	cfgType  string
	cfgPath  string
	CallBack func(fsnotify.Event)
}

var c *CnfYaml

func init() {
	c = New()

	// 设置默认路径
	c.SetCfgPath(defaultDir)

	// 默认文件名
	c.SetCfgName(defaultName)

	// 默认文件类型
	c.SetCfgType(defaultType)

}

func New() *CnfYaml {
	c := new(CnfYaml)
	return c

}

// SetCfgName sets name for the config file.
// Does not include extension.
func SetCfgName(in string) {
	c.SetCfgName(in)
}

func (c *CnfYaml) SetCfgName(in string) {
	if in != "" {
		c.cfgName = in
	}
}

// SetCfgType sets the type of the configuration returned by the
// remote source, e.g. "json".
func SetCfgType(in string) {
	c.SetCfgType(in)
}
func (c *CnfYaml) SetCfgType(in string) {
	if in != "" {
		c.cfgType = in
	}
}

// SetCfgPath adds a path for Viper to search for the config file in.
// Can be called multiple times to define multiple search paths.
func SetCfgPath(in string) {
	c.SetCfgPath(in)
}
func (c *CnfYaml) SetCfgPath(in string) {
	if in != "" {
		c.cfgPath = in
	}
}

// Get returns an interface. For a specific value use one of the Get____ methods.
func Get(key string) any {
	return viper.Get(key)
}

func GetStringMap(key string) any {
	return viper.GetStringMap(key)
}

func Unmarshal(c interface{}) (err error) {
	err = viper.Unmarshal(&c)
	return err
}

// ReadCfg will discover and load the configuration file from disk
// and key/value stores, searching in one of the defined paths.
func Install() (err error) {
	// 设置配置文件的名字
	viper.SetConfigName(c.cfgName)
	// 设置配置文件的类型
	viper.SetConfigType(c.cfgType)
	// 添加配置文件的路径，指定 config 目录下寻找
	viper.AddConfigPath(c.cfgPath)
	// 寻找配置文件并读取
	err = viper.ReadInConfig()

	if c.CallBack != nil {
		viper.WatchConfig()
		viper.OnConfigChange(c.CallBack)
	}
	err = Unmarshal(&cnf)
	return
	// fmt.Println(viper.Get("mysql"))     // map[port:3306 url:127.0.0.1]
	// fmt.Println(viper.Get("mysql.url")) // 127.0.0.1
}
