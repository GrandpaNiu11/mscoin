package config

import (
	"jobcenter/database"
	"jobcenter/kline"
)

type Config struct {
	Okx   kline.OkxConfig
	Mongo database.MongoConfig
}
