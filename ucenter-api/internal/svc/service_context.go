package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"grpc-common/uclient"
	"ucenter-api/internal/config"
)

type ServiceContext struct {
	Config        config.Config
	UCRegisterRpc uclient.Register
	UCLoginRpc    uclient.Login
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UCRegisterRpc: uclient.NewRegister(zrpc.MustNewClient(c.UCenterRpc)),
		UCLoginRpc:    uclient.NewLogin(zrpc.MustNewClient(c.UCenterRpc)),
	}
}
