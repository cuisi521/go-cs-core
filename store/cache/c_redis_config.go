// package credis
// @Author cuisi
// @Date 2023/12/5 09:33:00
// @Desc
package cache

import (
	"crypto/tls"
	"time"
)

type Config struct {
	Address         string        `json:"address"`
	Db              int           `json:"db"`
	User            string        `json:"user"`
	Pass            string        `json:"pass"`
	MinIdle         int           `json:"minIdle"`
	MaxIdle         int           `json:"maxIdle"`
	MaxActive       int           `json:"maxActive"`
	MaxConnLifetime time.Duration `json:"maxConnLifetime"`
	IdleTimeout     time.Duration `json:"idleTimeout"`
	WaitTimeout     time.Duration `json:"waitTimeout"`
	DialTimeout     time.Duration `json:"dialTimeout"`
	ReadTimeout     time.Duration `json:"readTimeout"`
	WriteTimeout    time.Duration `json:"writeTimeout"`
	MasterName      string        `json:"masterName"`
	TLS             bool          `json:"tls"`
	TLSSkipVerify   bool          `json:"tlsSkipVerify"`
	TLSConfig       *tls.Config   `json:"-"`
	SlaveOnly       bool          `json:"slaveOnly"`
	PoolSize        int           `json:"poolSize"`
}
