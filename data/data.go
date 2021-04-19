package data

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"sync"
)

type DbConnector struct {
	dbOnce sync.Once
	db     *gorm.DB
}

var Db *gorm.DB

func (d *DbConnector) Connect() *gorm.DB {
	d.dbOnce.Do(func() {
		dbHost := os.Getenv("DbHost")
		dbUser := os.Getenv("DbUser")
		dbPassword := os.Getenv("DbPassword")
		dbName := os.Getenv("DbName")
		connString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True", dbUser, dbPassword, dbHost, dbName)
		sqlDB, err := sql.Open("mysql", connString)
		if err != nil {
			log.Print(err)
			panic(err)
		}
		gormDB, err := gorm.Open(mysql.New(mysql.Config{
			Conn: sqlDB,
		}), &gorm.Config{})
		if err != nil {
			log.Print(err)
			panic(err)
		}
		d.db = gormDB
	})
	return d.db
}
