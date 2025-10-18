package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UCenterRpc zrpc.RpcClientConf
	Mysql      MysqlConfig
	CacheRedis cache.CacheConf
}

type MysqlConfig struct {
	DataSource string
}
