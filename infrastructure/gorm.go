package infrastructure

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialcts/mssql"
	_ "github.com/jinzhu/gorm/dialcts/postgres"
	_ "github.com/jinzhu/gorm/dialcts/sqlite"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var ormPool map[string]*DB

type DB struct {
	gorm.DB
}

func init() {
	OrmPool = make(map[string]*DB)
}

func NewOrmPool() map[string]*DB {
	return ormPool
}

func InitOrm(dbname string) {
	var orm *gorm.DB
	var err error

	// Default config
	C.Config.SetDefault(dbname, map[string]interface{}{
		"dbHost":          "127.0.0.1",
		"dbName":          "PowerCalcDev",
		"dbUser":          "root",
		"dbPassword":      "",
		"dbPort":          3306,
		"dbIdleconns_max": 0,
		"dbOpenconns_max": 20,
		"dbType":          "mysql",
	})

	dbHost := C.Config.GetString(dbname + ".dbHost")
	dbName := C.Config.GetString(dbname + ".dbName")
	dbUser := C.Config.GetString(dbname + ".dbUser")
	dbPassword := C.Config.GetString(dbname + ".dbPassword")
	dbPort := C.Config.GetString(dbname + ".dbPort")
	dbType := C.Config.GetString(dbname + ".dbType")

	connectString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName)

	// Debug mode
	// GDB.LogMode(true)

	for orm, err = gorm.Open(dbType, connectString); err != nil; {
		fmt.Println("Data connection error, retry after 5s")
		time.Sleep(5 * time.Second)
		orm, err = gorm.Open(dbType, connectString)
	}

	// size of idle connection pool
	orm.DB.SetMaxIdleConns(C.Config.GetInt(dbname + ".idleconns_max"))
	// max connection count
	orm.DB.SetMaxOpenConns(C.Config.GetInt(dbname + ".openconns_max"))
	OrmPool[dbname] = orm
}

func (this *OrmPool) GetOrm(dbname string) *DB {
	return OrmPool[dbname]
}

func (this *OrmPool) GetOrmByDefault() *DB {
	return OrmPool["dbDefault"]
}
