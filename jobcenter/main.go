package main

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	config "jobcenter/internal"
	"jobcenter/internal/svc"
	"jobcenter/task"
)

var configFile = flag.String("f", "etc/conf.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	t := task.NewTask(ctx)
	t.Run()
	t.StartBlocking()
}
