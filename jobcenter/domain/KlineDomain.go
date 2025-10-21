package domain

import (
	"context"
	"jobcenter/dao"
	"jobcenter/database"
	"jobcenter/model"
	"jobcenter/repo"
	"log"
)

type KlineDomain struct {
	KlineRepo repo.KlineRepo
}

func (d *KlineDomain) SaveBeath(data [][]string, symbol string, period string) {
	klines := make([]*model.Kline, len(data))
	for i, v := range data {
		klines[i] = model.NewKline(v, period)
	}
	err := d.KlineRepo.SaveBatch(context.Background(), klines, symbol, period)
	if err != nil {
		log.Println(err)
	}
}

func NewKlineDomain(cli *database.MongoClient) *KlineDomain {
	return &KlineDomain{
		KlineRepo: dao.NewKlineDao(cli.Db),
	}
}
