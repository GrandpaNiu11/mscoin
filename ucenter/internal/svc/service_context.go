package svc

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"mscoin-common/msdb"
	"ucenter/internal/config"
	"ucenter/internal/database"
)

type ServiceContext struct {
	Config config.Config
	Cache  cache.Cache
	DB     *msdb.MsDB
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisCache := cache.New(c.CacheRedis, nil, cache.NewStat("mscoin"), nil, func(o *cache.Options) {})

	return &ServiceContext{
		Config: c,
		Cache:  redisCache,
		DB:     database.ConnMysql(c.Mysql.DataSource),
	}
}
