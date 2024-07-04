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
	// for i := 0; i < 10; i++ {
	// 	result, _ := xg.Query(sqlStr)
	// 	fmt.Println("len:", len(result))
	// }
	for i := 0; i < 10; i++ {
		_, err := xg.Exec(`insert into sys_logininfor("msg") values(?)`, i)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("success")
		}
	}
	defer xg.Close()

}

func TestFzReg(t *testing.T) {
	_install()
	cdb.Install()
	sqlStr := fmt.Sprintf(`SELECT DISTINCT ON (a.attname) a.attname AS column_name,
										pg_catalog.format_type(a.atttypid, a.atttypmod) AS data_type,
										pg_catalog.col_description(a.attrelid, a.attnum) AS description,
										CASE
											WHEN i.indisprimary THEN 'PRIMARY KEY'
											WHEN pg_indexes.indexdef LIKE '%s' || a.attname || '%s' THEN 'HAS INDEX'
											ELSE ''
										END AS key_or_index
									FROM
										pg_catalog.pg_attribute a
										JOIN pg_catalog.pg_class c ON a.attrelid = c.oid
										LEFT JOIN pg_catalog.pg_namespace n ON n.oid = c.relnamespace
										LEFT JOIN pg_catalog.pg_index i ON i.indrelid = a.attrelid AND a.attnum = ANY(i.indkey)
										LEFT JOIN pg_indexes ON pg_indexes.tablename = c.relname AND pg_indexes.schemaname = n.nspname
									WHERE
										c.relname = '%s' -- 表名
										AND a.attnum > 0
										AND NOT a.attisdropped
										AND n.nspname = 'public' -- 如果表在 public schema
									ORDER BY
										a.attname,
										CASE
											WHEN i.indisprimary THEN 0
											WHEN pg_indexes.indexdef LIKE '%s' || a.attname || '%s' THEN 1
											ELSE 2
										END;`, "%", "%", "sys_logininfor", "%", "%")
	resultMap, err := cdb.DB().Query(sqlStr)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, v := range resultMap {
		// fmt.Println("======", k, v)
		if value, ok := v["column_name"]; ok {
			ColumnName := string(value)
			fmt.Println("ColumnName:", ColumnName)
		}
		if value, ok := v["description"]; ok {
			ColumnComment := string(value)
			fmt.Println("ColumnComment:", ColumnComment)
		}
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
