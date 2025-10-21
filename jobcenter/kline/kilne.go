package kline

import (
	"encoding/json"
	"jobcenter/database"
	"jobcenter/domain"
	"log"
	"mscoin-common/tools"
	"sync"
	"time"
)

type OkxConfig struct {
	ApiLey    string
	SecretKey string
	Pass      string
	Host      string
}

type kline struct {
	wg          sync.WaitGroup
	c           OkxConfig
	KlineDomain *domain.KlineDomain
}
type OkxKlineRes struct {
	Code string     `json:"code"`
	Msg  string     `json:"msg"`
	Data [][]string `json:"data"`
}

func (k *kline) Do(period string) {
	log.Println("============启动k线数据拉取==============")
	k.wg.Add(2)
	go k.syncToMongo("BTC-USDT", "BTC/USDT", period)
	go k.syncToMongo("ETH-USDT", "ETH/USDT", period)
	k.wg.Wait()
	log.Println("===============k线数据拉取结束===============")
}

func (k *kline) syncToMongo(instId string, symbol string, period string) {
	// ✅ 正确：添加完整的 URL
	api := k.c.Host + "/api/v5/market/candles?instId=" + instId + "&bar=" + period

	timestamp := tools.ISO(time.Now())
	// ✅ 注意：签名中的路径要去掉域名部分
	sign := tools.ComputeHmacSha256(timestamp+"GET"+"/api/v5/market/candles?instId="+instId+"&bar="+period, k.c.SecretKey)

	header := make(map[string]string)
	header["OK-ACCESS-KEY"] = k.c.ApiLey
	header["OK-ACCESS-SIGN"] = sign
	header["OK-ACCESS-TIMESTAMP"] = timestamp
	header["OK-ACCESS-PASSPHRASE"] = k.c.Pass

	resp, err := tools.GetWithHeader(api, header, "")
	if err != nil {
		log.Println("HTTP请求失败:", err)
		k.wg.Done()
		return
	}

	var result = &OkxKlineRes{}
	errc := json.Unmarshal(resp, &result)
	if errc != nil {
		log.Println("JSON解析失败:", errc)
		k.wg.Done()
		return
	}

	log.Println("==============获取到的k线数据============")
	log.Println("instId:", instId, "period", period)
	log.Println("result:", string(resp))
	log.Println("==========执行存储===================")
	if result.Code == "0" {
		//代表成功
		k.KlineDomain.SaveBeath(result.Data, symbol, period)
	}
	k.wg.Done()
}

func NewKline(c OkxConfig, client *database.MongoClient) *kline {
	return &kline{
		c:           c,
		KlineDomain: domain.NewKlineDomain(client),
	}
}
