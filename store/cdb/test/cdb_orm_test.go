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
	"github.com/xormplus/xorm"
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

func TestDb(t *testing.T) {
	var err error
	// 1.创建db引擎
	db, err = xorm.NewPostgreSQL("postgres://postgres:clm@2023@150.158.46.32:5433/dbtest?sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}

	// 2.显示sql语句
	db.ShowSQL(true)

	// insertObj()
	// insertExec()
	isnertExecute()
	// 4.2.3使用sql配置文件管理语句,两种载入配置的方式LoadSqlMap()和RegisterSqlMap(),以及SqlMapClient()替代SQL()
	if false {
		err = db.LoadSqlMap("./sql.xml")
		// err = db.RegisterSqlMap(xorm.Xml("./","sql.xml"))
		if err != nil {
			fmt.Println(err)
		}
		db.SqlMapClient("insert_1", "ft7").Execute()
	}

}
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

func isnertExecute() {
	// 4.2.2 db.SQL().Execute():单事务准备了Statement处理sql语句
	sql := "insert into public.user(name) values(?)"
	db.SQL(sql, "ft5").Execute()

}

func joinTest() {

}

// 读写分离
func fz() {
	var err error
	master, err := xorm.NewEngine("postgres", "postgres://postgres:root@localhost:5432/test?sslmode=disable")
	if err != nil {
		return
	}

	slave1, err := xorm.NewEngine("postgres", "postgres://postgres:root@localhost:5432/test1?sslmode=disable")
	if err != nil {
		return
	}

	slave2, err := xorm.NewEngine("postgres", "postgres://postgres:root@localhost:5432/test2?sslmode=disable")
	if err != nil {
		return
	}

	slaves := []*xorm.Engine{slave1, slave2}
	masters := []*xorm.Engine{master, slave2}
	eg, err := xorm.NewEngineGroup(masters, slaves)
	eg.DB()
}

func TestZc(t *testing.T) {
	var err error
	master, err := xorm.NewEngine("postgres", "postgres://postgres:clm@2023@150.158.46.32:5433/dbtest?sslmode=disable")
	if err != nil {
		return
	}

	slave1, err := xorm.NewEngine("postgres", "postgres://postgres:clm@2023@150.158.46.32:5433/dbtest?sslmode=disable")
	if err != nil {
		return
	}

	slave2, err := xorm.NewEngine("postgres", "postgres://postgres:clm@2023@150.158.46.32:5433/dbtest?sslmode=disable")
	if err != nil {
		return
	}

	slaves := []*xorm.Engine{slave1, slave2}
	eg, err := xorm.NewEngineGroup(master, slaves)
	fmt.Println(eg)
}
