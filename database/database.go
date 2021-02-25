package database

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type iDatabase interface {
	Initialize()
	GetDB() *gorm.DB
}

type database struct{}

var DB iDatabase = &database{}

var (
	db *gorm.DB
)

func (database) GetDB() *gorm.DB {
	if db == nil {
		db = connect()
	}
	return db
}

func (database) Initialize() {
	if db == nil {
		db = connect()
	}
	migrate()
}

func connect() *gorm.DB {
	user := "root"
	pass := "root"
	host := "localhost"
	port := "8889"
	//name := "person-crawler"
	name := "person-crawler-2"
	dsn := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		pass,
		host,
		port,
		name,
	)
	connection, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		panic(err)
	}
	return connection
}

func migrate() {
	if err := db.AutoMigrate(&Person{}); err != nil {
		panic(err)
	}
}
