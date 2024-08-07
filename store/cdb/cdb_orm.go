// package cdb
// @Author cuisi
// @Date 2023/11/23 23:34:00
// @Desc
package cdb

import (
	"fmt"

	"xorm.io/xorm"

	"github.com/cuisi521/go-cs-core/sys/clog"
)

const (
	DefaultName = "default"
	Master      = "Master"
	Slave       = "Slave"
)

var (
	// 数据库引擎
	dbEngine      map[string]*DbEngine
	dbEngineGroup map[string]*xorm.EngineGroup

	// 配置信息
	dbCnfs map[string]*DbCnfs

	// 负载均衡-轮询
	// lbs map[string]*rotation.RoundRotationBalance
)

func SetCnf(cnf map[string]*DbCnfs) {
	dbCnfs = cnf
}

// DB 数据库对象
// d[0]数据库分组,d[1]指定数据库别名
// @author By Cuisi 2023/12/4 15:48:00
// func DB(d ...string) (xe *xorm.Engine) {
// 	var _gp, _db string = DefaultName, DefaultName
// 	if len(d) <= 1 {
// 		// 判断是否负载均衡
// 		if g, ok := lbs[_gp]; ok {
// 			if g.Next() != "" {
// 				_db = g.Next()
// 			}
// 		}
// 	} else {
// 		_gp = d[0]
// 		_db = d[1]
// 	}
// 	if v, ok := dbEngine[_db]; ok {
// 		fmt.Println(v.alias, v.role)
// 		return v.engine
// 	}
// 	return
// }

func DB(d ...string) (xe *xorm.EngineGroup) {
	var (
		_gp string = DefaultName
	)
	if len(d) > 0 {
		_gp = d[0]
	}

	if g, ok := dbEngineGroup[_gp]; ok {
		xe = g
	} else {
		for k, v := range dbCnfs {
			for _, _v := range v.DbCnnf {
				if _v.Role == Master {
					_gp = k
				}
			}
		}
	}
	return
}

func GetCnf(k string) (cnf *DbCnfs) {
	return dbCnfs[k]
}

// Install
// @author By Cuisi 2023/12/4 17:40:00
// func Install() {
// 	lbs = make(map[string]*rotation.RoundRotationBalance)
// 	dbEngine = make(map[string]*DbEngine)
// 	for k, v := range dbCnfs {
// 		r := &rotation.RoundRotationBalance{}
// 		dbStatus := false
// 		for _k, _v := range v.DbCnnf {
// 			db := createDb(_v)
// 			if db != nil {
// 				dbStatus = true
// 				if _v.Role == Master || _v.Role == Slave {
// 					r.Add(_k)
// 				}
// 				de := &DbEngine{
// 					alias:  _k,
// 					engine: db,
// 					role:   _v.Role,
// 				}
// 				dbEngine[_k] = de
// 			}
// 		}
// 		if dbStatus {
// 			lbs[k] = r
// 		}
// 	}
// }

func Install() {
	dbEngineGroup = make(map[string]*xorm.EngineGroup)
	for k, v := range dbCnfs {
		var dbMaster *xorm.Engine
		var dbSlaves = make([]*xorm.Engine, 0)
		for _k, _v := range v.DbCnnf {
			db := createDb(_v)
			if db != nil {
				dbAlias := k + "$" + _k
				db.Alias(dbAlias)
				switch _v.Role {
				case Master:
					dbMaster = db
					// dbSlaves = append(dbSlaves, db)
				case Slave:
					dbSlaves = append(dbSlaves, db)
				default:
					dbMaster = db
				}
			}
		}
		if dbMaster == nil {
			continue
		}
		// 策略组,轮训
		xg, err := xorm.NewEngineGroup(dbMaster, dbSlaves, xorm.RoundRobinPolicy())
		if err != nil {
			fmt.Println(err.Error())
		} else {
			dbEngineGroup[k] = xg
		}
	}
}

func CloseDb() {
	for _, v := range dbEngineGroup {
		if v != nil {
			v.Close()
		}
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
	// 连接测试
	if db != nil {
		err = db.Ping()
		if err != nil {
			db = nil
		}
	}
	return db
}

type DbEngine struct {
	alias  string
	role   string
	engine *xorm.Engine
}
