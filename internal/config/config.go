package config

import "github.com/zeromicro/go-zero/rest"

var c *Config

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	} `json:"jwtAuth"`
	DataBase DataBase `json:"dataBase"`
}

type DataBase struct {
	Mysql Mysql `json:"mysql"`
}

type Mysql struct {
	Dsn string `json:"dsn"`
}

func init() {
	c = &Config{}
}

func GetConfig() *Config {
	return c
}
