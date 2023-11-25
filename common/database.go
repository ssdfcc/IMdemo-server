package common

import (
	"ariga.io/entcache"
	"database/sql"
	entsql "entgo.io/ent/dialect/sql"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/logx"
	redis2 "github.com/zeromicro/go-zero/core/stores/redis"
	"os"
	"time"
)

// DatabaseConf stores database configurations.
type DatabaseConf struct {
	Host         string `json:",env=DATABASE_HOST"`
	Port         int    `json:",env=DATABASE_PORT"`
	Username     string `json:",default=root,env=DATABASE_USERNAME"`
	Password     string `json:",optional,env=DATABASE_PASSWORD"`
	DBName       string `json:",default=simple_admin,env=DATABASE_DBNAME"`
	SSLMode      string `json:",optional,env=DATABASE_SSL_MODE"`
	Type         string `json:",default=mysql,options=[mysql,postgres,sqlite3],env=DATABASE_TYPE"`
	MaxOpenConn  int    `json:",optional,default=100,env=DATABASE_MAX_OPEN_CONN"`
	CacheTime    int    `json:",optional,default=10,env=DATABASE_CACHE_TIME"`
	DBPath       string `json:",optional,env=DATABASE_DBPATH"`
	MysqlConfig  string `json:",optional,env=DATABASE_MYSQL_CONFIG"`
	PGConfig     string `json:",optional,env=DATABASE_PG_CONFIG"`
	SqliteConfig string `json:",optional,env=DATABASE_SQLITE_CONFIG"`
}

// NewCacheDriver returns an Ent driver with cache.
func (c DatabaseConf) NewCacheDriver(redisConf redis2.RedisConf) *entcache.Driver {
	db, err := sql.Open(c.Type, c.GetDSN())
	logx.Must(err)

	db.SetMaxOpenConns(c.MaxOpenConn)
	driver := entsql.OpenDB(c.Type, db)

	rdb := redis.NewClient(&redis.Options{Addr: redisConf.Host})

	cacheDrv := entcache.NewDriver(
		driver,
		entcache.TTL(time.Duration(c.CacheTime)*time.Second),
		entcache.Levels(
			entcache.NewLRU(256),
			entcache.NewRedis(rdb),
		),
	)

	return cacheDrv
}

// NewNoCacheDriver returns an Ent driver without cache.
func (c DatabaseConf) NewNoCacheDriver() *entsql.Driver {
	db, err := sql.Open(c.Type, c.GetDSN())
	logx.Must(err)

	db.SetMaxOpenConns(c.MaxOpenConn)
	driver := entsql.OpenDB(c.Type, db)

	return driver
}

// MysqlDSN returns mysql DSN.
func (c DatabaseConf) MysqlDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True%s", c.Username, c.Password, c.Host, c.Port, c.DBName, c.MysqlConfig)
}

// PostgresDSN returns Postgres DSN.
func (c DatabaseConf) PostgresDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s %s", c.Host, c.Username, c.Password, c.DBName, c.Port,
		c.SSLMode, c.PGConfig)
}

// SqliteDSN returns Sqlite DSN.
func (c DatabaseConf) SqliteDSN() string {
	if c.DBPath == "" {
		logx.Must(errors.New("the database file path cannot be empty"))
	}

	if _, err := os.Stat(c.DBPath); os.IsNotExist(err) {
		f, err := os.OpenFile(c.DBPath, os.O_CREATE|os.O_RDWR, 0600)
		if err != nil {
			logx.Must(fmt.Errorf("failed to create SQLite database file %q", c.DBPath))
		}
		if err := f.Close(); err != nil {
			logx.Must(fmt.Errorf("failed to create SQLite database file %q", c.DBPath))
		}
	} else {
		if err := os.Chmod(c.DBPath, 0660); err != nil {
			logx.Must(fmt.Errorf("unable to set permission code on %s: %v", c.DBPath, err))
		}
	}

	return fmt.Sprintf("file:%s?_busy_timeout=100000&_fk=1%s", c.DBPath, c.SqliteConfig)
}

// GetDSN returns DSN according to the database type.
func (c DatabaseConf) GetDSN() string {
	switch c.Type {
	case "mysql":
		return c.MysqlDSN()
	case "postgres":
		return c.PostgresDSN()
	case "sqlite3":
		return c.SqliteDSN()
	default:
		return "mysql"
	}
}
