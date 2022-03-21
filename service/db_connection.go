package service

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Need for db driver

	"jimmy_tech_crud_gin/config"
)

// ConnectDB connects to database.
func ConnectDB() (*gorm.DB, error) {
	connectionString, driver, err := createConnectionString()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(driver, connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Creates string for SQL connection.
func createConnectionString() (string, string, error) {
	conf, err := config.GetSQLConf()
	if nil != err {
		return "", "", err
	}

	return fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
	), conf.Driver, nil
}
