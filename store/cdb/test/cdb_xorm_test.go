package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/cuisi521/go-cs-core/store/cdb"
	"xorm.io/xorm"
)

func TestZcs(t *testing.T) {
	fmt.Println("测试开始。。。")
	var err error
	master, err := xorm.NewEngine("postgres", "postgres://postgres:redmoon@127.0.0.1:5432/dbtest?sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// master.ShowSQL(true)

	slave1, err := xorm.NewEngine("postgres", "postgres://postgres:redmoon@127.0.0.1:5432/dbtest1?sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// slave1.ShowSQL(true)

	slave2, err := xorm.NewEngine("postgres", "postgres://postgres:redmoon@127.0.0.1:5432/dbtest2?sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// slave2.ShowSQL(true)

	slaves := []*xorm.Engine{master, slave1, slave2}
	xg, err := xorm.NewEngineGroup(master, slaves)
	xg.SetPolicy(xorm.RoundRobinPolicy())

	// t.Log("[xg.Slave():\n", xg.Slave().DataSourceName())
	// t.Log("[xg.Slave():\n", xg.Slave().DataSourceName())
	// t.Log("[xg.Slave():\n", xg.Slave().DataSourceName())
	// t.Log("[xg.Slave():\n", xg.Slave().DataSourceName())
	// t.Log("[xg.Slave():\n", xg.Slave().DataSourceName())
	// t.Log("[xg.Slave():\n", xg.Slave().DataSourceName())
	// t.Log("[xg.Slave():\n", xg.Slave().DataSourceName())
	// t.Log("[xg.Slave():\n", xg.Slave().DataSourceName())
	// t.Log("[xg.Slave():\n", xg.Slave().DataSourceName())
	// t.Log("[xg.Slave():\n", xg.Slave().DataSourceName())
	// t.Log("[xg.Slave():\n", xg.Slave().DataSourceName())

	// sqlStr := `select * from sys_logininfor`
	for i := 0; i < 10; i++ {
		var dataList []SysLogininfor
		xg.Find(&dataList)
		fmt.Println("len:", len(dataList))
	}
	//
	// for i := 0; i < 10; i++ {
	// 	_, err := xg.Exec(`insert into sys_logininfor("msg") values(?)`, i)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	} else {
	// 		fmt.Println("success")
	// 	}
	// }
	defer xg.Close()

}

func TestFzReg(t *testing.T) {
	_install()
	cdb.Install()
	sqlStr := `select * from sys_logininfor`
	for i := 0; i < 10; i++ {
		result, _ := cdb.DB().Query(sqlStr)
		cdb.DB().Slave()
		fmt.Println("len:", len(result))
	}

	// for i := 0; i < 10; i++ {
	// 	_, err := cdb.DB().Exec(`insert into sys_logininfor("msg") values(?)`, i)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	} else {
	// 		fmt.Println("success")
	// 	}
	// }
	// t.Log("[xg.Slave():\n", cdb.DB().Slave().DataSourceName())
	// t.Log("[xg.Slave():\n", cdb.DB().Slave().DataSourceName())
	// t.Log("[xg.Slave():\n", cdb.DB().Slave().DataSourceName())
	// t.Log("[xg.Slave():\n", cdb.DB().Slave().DataSourceName())
	// t.Log("[xg.Slave():\n", cdb.DB().Slave().DataSourceName())
	// t.Log("[xg.Slave():\n", cdb.DB().Slave().DataSourceName())
}

func _install() {
	cfg := make(map[string]*cdb.DbCnfs)
	cnf := make(map[string]*cdb.DbCnf)
	df := &cdb.DbCnf{
		Alias:           "t1",
		Driver:          "postgres",
		Link:            "postgres://postgres:redmoon@127.0.0.1:5432/dbtest?sslmode=disable",
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
		Link:            "postgres://postgres:redmoon@127.0.0.1:5432/dbtest1?sslmode=disable",
		ShowSql:         true,
		LogLevel:        0,
		ConnMaxLifeTime: 0,
		MaxIdleConn:     0,
		MaxOpenConn:     0,
		Role:            "Slave",
	}
	cnf[df1.Alias] = df1

	df2 := &cdb.DbCnf{
		Alias:           "default",
		Driver:          "postgres",
		Link:            "postgres://postgres:redmoon@127.0.0.1:5432/dbtest2?sslmode=disable",
		ShowSql:         true,
		LogLevel:        0,
		ConnMaxLifeTime: 0,
		MaxIdleConn:     0,
		MaxOpenConn:     0,
		Role:            "Master",
	}
	cnf[df2.Alias] = df2

	dbcnfs := &cdb.DbCnfs{"db1", cnf}
	cfg["default"] = dbcnfs

	cdb.SetCnf(cfg)
}

type SysLogininfor struct {
	Id         int64     `xorm:"bigint pk autoincr 'id' comment('访问ID')" json:"id,string" description:"访问ID"`
	UserName   string    `xorm:"Varchar(50) 'user_name' comment('用户账号')" json:"userName" description:"用户账号"`
	Ipaddr     string    `xorm:"Varchar(128) 'ipaddr' comment('登录IP地址')" json:"ipaddr" description:"登录IP地址"`
	Status     string    `xorm:"character(2) 'status' comment('登录状态（0成功;1失败）')" json:"status" description:"登录状态（0成功;1失败）"`
	Msg        string    `xorm:"Varchar(255) 'msg' comment('提示信息')" json:"msg" description:"提示信息"`
	AccessTime time.Time `xorm:"DateTime 'access_time' comment('访问时间')" json:"accessTime" description:"访问时间"`
	Os         string    `xorm:"Varchar(65) 'os' comment('操作系统')" json:"os" description:"操作系统"`
	Module     string    `xorm:"Varchar(35) 'module' comment('登录模块')" json:"module" description:"登录模块"`
	Browser    string    `xorm:"Varchar(60) 'browser' comment('浏览器')" json:"browser" description:"浏览器"`
}

// func (*SysLogininfor) TableName() string { return "sys_logininfor" }
