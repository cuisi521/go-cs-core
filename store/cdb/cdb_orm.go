// package cdb
// @Author cuisi
// @Date 2023/11/23 23:34:00
// @Desc
package cdb

import (
	"fmt"

	"github.com/xormplus/xorm"

	"github.com/cuisi521/go-cs-core/sys/clog"
	"github.com/cuisi521/go-cs-core/tool/clb/rotation"
)

const (
	DefaultName = "default"
	Master      = "Master"
	Slave       = "Slave"
)

var (
	// 数据库引擎
	dbEngine map[string]*DbEngine

	// 配置信息
	dbCnfs map[string]*DbCnfs

	// 负载均衡-轮询
	lbs map[string]*rotation.RoundRotationBalance
)

func SetCnf(cnf map[string]*DbCnfs) {
	dbCnfs = cnf
}

// DB 数据库对象
// d[0]数据库分组,d[1]指定数据库别名
// @author By Cuisi 2023/12/4 15:48:00
func DB(d ...string) (xe *xorm.Engine) {
	var _gp, _db string = DefaultName, DefaultName
	if len(d) <= 1 {
		// 判断是否负载均衡
		if g, ok := lbs[_gp]; ok {
			_db = g.Next()
		}
	} else {
		_gp = d[0]
		_db = d[1]
	}
	if v, ok := dbEngine[_db]; ok {
		fmt.Println(v.alias, v.role)
		return v.engine
	}
	return
}

func GetCnf(k string) (cnf *DbCnfs) {
	return dbCnfs[k]
}

// Install
// @author By Cuisi 2023/12/4 17:40:00
func Install() {
	lbs = make(map[string]*rotation.RoundRotationBalance)
	dbEngine = make(map[string]*DbEngine)
	for k, v := range dbCnfs {
		r := &rotation.RoundRotationBalance{}
		for _k, _v := range v.DbCnnf {
			if _v.Role == Master || _v.Role == Slave {
				r.Add(_k)
			}
			db := createDb(_v)
			de := &DbEngine{
				alias:  _k,
				engine: db,
				role:   _v.Role,
			}
			dbEngine[_k] = de
		}
		lbs[k] = r
	}
}

// createDb 创建单个数据库实例
// @author By Cuisi 2023/11/2 17:32:00
func createDb(cnf *DbCnf) *xorm.Engine {
	db, err := xorm.NewEngine(cnf.Driver, cnf.Link)
	if err != nil {
		clog.Error(err.Error())
	} else {
		db.SetConnMaxLifetime(cnf.ConnMaxLifeTime)
		db.SetMaxIdleConns(cnf.MaxIdleConn)
		db.SetMaxOpenConns(cnf.MaxOpenConn)
		db.ShowSQL(cnf.ShowSql)
	}
	// 设置缓存
	// cacher := caches.NewLRUCacher(caches.NewMemoryStore(), 1000)
	// db.SetDefaultCacher(cacher)
	return db
}

type DbEngine struct {
	alias  string
	role   string
	engine *xorm.Engine
}
