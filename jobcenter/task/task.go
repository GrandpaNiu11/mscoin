package task

import (
	"github.com/go-co-op/gocron"
	"jobcenter/internal/svc"
	"jobcenter/kline"
	"time"
)

type Task struct {
	s   *gocron.Scheduler
	ctx *svc.ServiceContext
}

func NewTask(ctx *svc.ServiceContext) *Task {
	return &Task{
		s:   gocron.NewScheduler(time.UTC),
		ctx: ctx,
	}
}

func (t *Task) Run() {
	t.s.Every(1).Minute().Do(func() {
		kline.NewKline(t.ctx.Config.Okx).Do("1m")
	})
}

func (t *Task) StartBlocking() {
	t.s.StartBlocking()
}
