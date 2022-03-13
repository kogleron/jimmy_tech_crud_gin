package config

import "os"

type TokenAuthConf struct {
	// Secret token
	Token string
	// Query param name containing secret token
	QueryParam string
}

func GetTokenAuthConf() TokenAuthConf {
	return TokenAuthConf{
		Token:      os.Getenv("AT_TOKEN"),
		QueryParam: os.Getenv("AT_QUERY_PARAM"),
	}
}
