package config

import (
	"os"
	"strconv"
)

type ServerConf struct {
	Port int64
}

func GetServerConf() (*ServerConf, error) {
	port, err := strconv.ParseInt(os.Getenv("SERVER_PORT"), 10, 32) //nolint:gomnd
	if err != nil {
		return nil, err
	}

	return &ServerConf{
		Port: port,
	}, nil
}
