package svc

import (
	"jobcenter/database"
	config "jobcenter/internal"
)

type ServiceContext struct {
	Config      config.Config
	MongoClient *database.MongoClient
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config:      c,
		MongoClient: database.ConnectMongo(c.Mongo),
	}
}
