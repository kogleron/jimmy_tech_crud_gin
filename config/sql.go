package config

import (
	"os"
	"strconv"
)

type SQLConf struct {
	Driver   string
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

func GetDBConf() (*SQLConf, error) {
	port, err := strconv.ParseInt(os.Getenv("MYSQL_PORT"), 10, 64) //nolint:gomnd
	if nil != err {
		return nil, err
	}

	return &SQLConf{
		Driver:   "mysql",
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     int(port),
		Database: os.Getenv("MYSQL_DATABASE"),
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
	}, nil
}
