package db

import (
	"fmt"
	"github.com/ctbsea/Go-Message/config"
	"github.com/ctbsea/Go-Message/util/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"log"
)

type Db struct {
	Mysql *gorm.DB
}

func InitDb(zapLogger *zap.Logger ,config config.Config) *Db {
	var db Db
	db.Mysql = inItMySqlConn(zapLogger ,config)
	return &db
}

func inItMySqlConn(zapLogger *zap.Logger ,config config.Config) *gorm.DB {
	//"user:password@tcp(ip:port)/dbname?charset=utf8&parseTime=True&loc=Local"
	user := config.MySQL.User
	password := config.MySQL.Password
	dbname := config.MySQL.Database
	ip := config.MySQL.IP
	port := config.MySQL.Port
	mysqlStr := user + ":" + password + "@tcp(" + ip + ":" + port + ")/" + dbname +
		 "?parseTime=true&charset=utf8&loc=Local"
	fmt.Println(mysqlStr)
	db, err := gorm.Open("mysql", mysqlStr)
	if err != nil {
		log.Fatal(err.Error())
	}
	db.LogMode(config.MySQL.Debug)
	db.DB().SetMaxIdleConns(config.MySQL.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.MySQL.MaxOpenConns)
	db.SetLogger(&logger.SqlLoggerMiddleware{Zap: zapLogger})
	return db
}
