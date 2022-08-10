package infra

import (
	"log"
	"time"

	// _ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlConfiguration struct {
	addr            string
	maxOpenConns    int
	maxIdleConns    int
	connMaxLifetime time.Duration
}
type MysqlConnPool struct {
	Conn   *gorm.DB
	config MysqlConfiguration
}

func NewMysqlConnPool(config MysqlConfiguration) (*MysqlConnPool, error) {
	db, err := gorm.Open(mysql.Open(config.addr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	pool, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	pool.SetMaxIdleConns(config.maxIdleConns)
	pool.SetMaxOpenConns(config.maxOpenConns)
	pool.SetConnMaxLifetime(config.connMaxLifetime)

	return &MysqlConnPool{Conn: db, config: config}, err
}

func NewMysqlConfiguration() MysqlConfiguration {
	const poolsize = 5
	const lifetimeMin = 60
	return MysqlConfiguration{
		addr:            "vuongnguyen3:vuong@tcp(127.0.0.1:3306)/1stpj?charset=utf8mb4&parseTime=True&loc=Local",
		maxOpenConns:    poolsize,
		maxIdleConns:    poolsize,
		connMaxLifetime: time.Duration(lifetimeMin) * time.Minute,
	}
}
