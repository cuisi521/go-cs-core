// package cdb
// @Author cuisi
// @Date 2023/11/1 17:03:00
// @Desc
package test

import (
	"fmt"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

// 用户表结构
type User struct {
	Id      int       `xorm:"not null pk autoincr INTEGER"`
	Name    string    `xorm:"VARCHAR(20)"`
	Created time.Time `xorm:"default 'now()' DATETIME"`
	ClassId int       `xorm:"default 1 INTEGER"`
}

// Class表结构
type Class struct {
	Id   int    `xorm:"not null pk autoincr INTEGER"`
	Name string `xorm:"VARCHAR(20)"`
}

// 临时表结构
type UserClass struct {
	User `xorm:"extends"`
	Name string
}

// 此方法仅用于orm查询时，查询表认定
func (UserClass) TableName() string {
	return "public.user"
}

var db *xorm.Engine

func insertObj() {
	// 4.执行插入语句的几种方式
	// 4.1 orm插入方式:不好控制，如果仅仅插入的对象的属性是name='ftq',那么其他的零值会一同insert，orm方式对零值的处理有点不太好
	user := new(User)
	user.Name = "ftq"
	_, err := db.Insert(user)

	if err != nil {
		fmt.Println(err)
	}
}

func insertExec() {
	// 4.2 命令插入方式
	// 4.2.1 db.Exec():单事务单次提交
	sql := "insert into public.user(name) values(?)"
	db.Exec(sql, "ft4")
}

func TestZc(t *testing.T) {
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

	t.Log("[xg.Slave():\n", xg.Slave().DataSourceName())
	t.Log("[xg.Slave():\n", xg.Slave().DataSourceName())
	t.Log("[xg.Slave():\n", xg.Slave().DataSourceName())
	t.Log("[xg.Slave():\n", xg.Slave().DataSourceName())
	t.Log("[xg.Slave():\n", xg.Slave().DataSourceName())
	t.Log("[xg.Slave():\n", xg.Slave().DataSourceName())
	t.Log("[xg.Slave():\n", xg.Slave().DataSourceName())
	t.Log("[xg.Slave():\n", xg.Slave().DataSourceName())
	t.Log("[xg.Slave():\n", xg.Slave().DataSourceName())
	t.Log("[xg.Slave():\n", xg.Slave().DataSourceName())
	t.Log("[xg.Slave():\n", xg.Slave().DataSourceName())

	// sqlStr := `select * from sys_logininfor`
	// for i := 0; i < 10; i++ {
	// 	result := xg.QueryRe(sqlStr)
	// 	fmt.Println("len:", len(result.Result))
	// }

	defer xg.Close()

}
