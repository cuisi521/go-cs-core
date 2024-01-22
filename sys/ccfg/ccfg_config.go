// package ccfg
// @Author cuisi
// @Date 2023/12/11 10:22:00
// @Desc
package ccfg

import (
	"time"
)

var (
	cnf *Config
)

// Config 配置
// @author By Cuisi 2023/12/11 14:50:00
type Config struct {
	Server   ServerCnf                   `mapstructure:"server"`
	Logger   LoggerCnf                   `mapstructure:"logger"`
	Ssl      SslCnf                      `mapstructure:"ssl"`
	Database map[string]map[string]DbCnf `mapstructure:"database"`
	Cache    CacheCnf                    `mapstructure:"cache"`
	Token    TokenCnf                    `mapstructure:"token"`
}

// ServerCnf 服务配置
// @author By Cuisi 2023/12/11 14:49:00
type ServerCnf struct {
	// 主机地址
	Host string `mapstructure:"host"`

	// 是否调试模式
	Debug bool `mapstructure:"debug"`

	// https 端口
	HttpsPort int `mapstructure:"httpsPort"`

	// http 端口
	HttpPort int `mapstructure:"httpPort"`

	// 请求头大小限制 默认：20KB
	MaxHeaderBytes string `mapstructure:"maxHeaderBytes"`

	// #客户端上传文件大小限制 默认：200MB
	ClientMaxBodySize string `mapstructure:"clientMaxBodySize"`

	// 分布式节点(0-1024)
	Node int64 `mapstructure:"node"`
}

// LoggerCnf 日志配置
// @author By Cuisi 2023/12/11 14:49:00
type LoggerCnf struct {
	// 级别
	Level string `mapstructure:"level"`

	// 是否输出
	StdOut bool `mapstructure:"stdout"`

	// 保存路径
	Path string `mapstructure:"path"`

	// 文件名称
	file string `mapstructure:"file"`

	// 输出颜色
	OutColor bool `mapstructure:"outColor"`
}

// SslCnf
// @author By Cuisi 2023/12/21 09:20:00
type SslCnf struct {
	Enable bool   `mapstructure:"enable"`
	Pem    string `mapstructure:"pem"`
	Key    string `mapstructure:"key"`
}

// DbCnf 单节点数据库配置
// @author By Cuisi 2023/12/11 15:27:00
type DbCnf struct {
	// 数据库驱动
	Driver string `mapstructure:"driver"`

	// 数据库链接
	Link string `mapstructure:"link"`

	// 编码
	Charset string `mapstructure:"charset"`

	// 是否调试模式
	Debug bool `mapstructure:"debug"`

	// 连接池最大闲置的连接数
	MaxIdle int `mapstructure:"maxIdle"`

	// 连接池最大打开的连接数
	MaxOpen int `mapstructure:"maxOpen"`

	// (单位秒)连接对象可重复使用的时间长度
	MaxLifetime time.Duration `mapstructure:"maxLifetime"`
}

// CacheCnf 缓存配置
// @author By Cuisi 2023/12/11 15:28:00
type CacheCnf struct {
	Mod   int                 `mapstructure:"mod"`
	Redis map[string]RedisCnf `mapstructure:"redis"`
}

// RedisCnf redis 配置
// @author By Cuisi 2023/12/11 15:28:00
type RedisCnf struct {
	Host        string `mapstructure:"host"`
	Db          int    `mapstructure:"db"`
	idleTimeout int    `mapstructure:"idleTimeout"`
	maxActive   int    `mapstructure:"maxActive"`
	Password    string `mapstructure:"pwd"`
	PoolSize    int    `mapstructure:"poolSize"`
}

// TokenCnf
// @author By Cuisi 2023/12/25 14:23:00
type TokenCnf struct {
	Key string `mapstructure:"key"`
	// 0 默认 jwt，1 内存 2 redis
	Mod int `mapstructure:"mod"`
	// redis 别名
	Redis        string `mapstructure:"redis"`
	CacheKey     string `mapstructure:"cacheKey"`
	ExcludePaths string `mapstructure:"excludePaths"`
	LoginPath    string `mapstructure:"loginPath"`
	LoginOutPath string `mapstructure:"loginOutPath"`
	TimeOut      int64  `mapstructure:"timeOut"`
	MaxRefresh   int64  `mapstructure:"maxRefresh"`
}

func SysCnf() *Config {
	return cnf
}
