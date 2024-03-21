// package cdb
// @Author cuisi
// @Date 2023/11/2 14:10:00
// @Desc
package test

import (
	"testing"

	"github.com/cuisi521/go-cs-core/store/cdb"
)

func install() {
	cfg := make(map[string]*cdb.DbCnfs)
	cnf := make(map[string]*cdb.DbCnf)
	df := &cdb.DbCnf{
		Alias:           "t1",
		Driver:          "postgres",
		Link:            "postgres://postgres:clm@2023@150.158.46.32:51433/dbtest?sslmode=disable",
		ShowSql:         true,
		LogLevel:        0,
		ConnMaxLifeTime: 0,
		MaxIdleConn:     0,
		MaxOpenConn:     0,
		Role:            "Slave",
	}

	cnf[df.Alias] = df
	df1 := &cdb.DbCnf{
		Alias:           "t2",
		Driver:          "postgres",
		Link:            "postgres://postgres:clm@2023@150.158.46.32:51433/dbtest1?sslmode=disable",
		ShowSql:         true,
		LogLevel:        0,
		ConnMaxLifeTime: 0,
		MaxIdleConn:     0,
		MaxOpenConn:     0,
		Role:            "Master",
	}
	cnf[df1.Alias] = df1

	dbcnfs := &cdb.DbCnfs{"default", cnf}
	cfg["default"] = dbcnfs
	cfg["v1"] = dbcnfs
	cdb.SetCnf(cfg)

}
func initDb() {
	install()
	cdb.Install()
}

func TestInstall(t *testing.T) {
	initDb()
	for i := 0; i < 10; i++ {
		cdb.DB("default")
	}

}
