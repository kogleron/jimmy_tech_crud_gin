package service

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Need for db driver

	"jimmy_tech_crud_gin/config"
)

// ConnectDB connects to database.
func ConnectDB(conf config.SQLConf) (*gorm.DB, error) {
	connectionString := buildConnectionString(conf)

	db, err := gorm.Open(conf.Driver, connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Creates string for SQL connection.
func buildConnectionString(conf config.SQLConf) string {
	return fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
	)
}
