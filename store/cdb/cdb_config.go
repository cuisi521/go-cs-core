// package cdb
// @Author cuisi
// @Date 2023/11/1 16:51:00
// @Desc
package cdb

import (
	"time"
)

// DbCfg 数据库配置
type DbCnf struct {
	// 配置名称
	Alias string `json:"alias"`

	// 驱动名称
	Driver string `json:"driver"`

	// 连接串
	Link string `json:"link"`

	// 打印Sql
	ShowSql bool `json:"showSql"`

	// 日志级别
	LogLevel int `json:"logLevel"`

	// 连接可能被重用的最大时间
	ConnMaxLifeTime time.Duration `json:"connMaxLifeTime"`

	// 最大空闲连接数
	MaxIdleConn int `json:"maxIdleConn"`

	// 打开连接数
	MaxOpenConn int `json:"maxOpenConn"`

	// 角色，一般用于读写分离
	Role string `json:"role"`
}

type DbCnfs struct {
	// 配置别名
	Alias string

	// 配置信息
	DbCnnf map[string]*DbCnf
}
