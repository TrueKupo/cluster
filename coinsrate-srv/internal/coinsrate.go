package coinsrate

import (
	"context"
	"sync"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/gocraft/dbr/v2"
	"github.com/sirupsen/logrus"

	"github.com/truekupo/cluster/coinsrate-srv/lib/config"
	"github.com/truekupo/cluster/coinsrate-srv/lib/db/middleware"
	"github.com/truekupo/cluster/coinsrate-srv/lib/db/models"
	"github.com/truekupo/cluster/coinsrate-srv/lib/db/repo"

	"github.com/truekupo/cluster/lib/logger"
)

type CoinsRateHandler struct {
	sync.Mutex

	conf *config.Config
	dbs  *dbr.Session
}

var (
	log *logrus.Entry = nil
)

func NewCoinsRateHandler(conf *config.Config) (*CoinsRateHandler, error) {
	log = logger.LogModule("coinsrate-handler")

	h := &CoinsRateHandler{
		conf: conf,
	}

	// connect to DB
	dsd, err := middleware.NewDSD(conf)
	if err != nil {
		log.Error("coinsrate:handler:new ", err)
		return nil, err
	}

	h.dbs = dsd.Conn.NewSession(nil)

	// start watch handler
	go h.coinsrate_process()

	return h, nil
}

func (h *CoinsRateHandler) coinsrate_process() {
	for {
		pairs, err := repo.Pair(h.dbs).FindActive()
		if err == nil {
			for _, p := range pairs {
				klines, err := h.GetKlines(p.CoinName + p.CurrencySymbol)
				if err != nil {
					log.WithField("error", err).Error("GetKlines")
					continue
				}

				log.WithField("pair_id", p.Id).WithField("symbol", p.CoinName+p.CurrencySymbol).WithField("kline_len", len(klines)).Debug("GetKlines")

				for i := len(klines) - 1; i >= 0; i-- {
					k := klines[i]

					_, err = repo.Kline(h.dbs).Insert(&models.Kline{
						PairId:    p.Id,
						OpenTime:  k.OpenTime,
						CloseTime: k.CloseTime,
						Open:      k.Open,
						Close:     k.Close,
						High:      k.High,
						Low:       k.Low,
					})
					if err != nil {
						break
					}
				}
			}
		}

		time.Sleep(30 * time.Second)
	}
}

func (h *CoinsRateHandler) GetKlines(Symbol string) ([]*binance.Kline, error) {
	client := binance.NewClient("", "")

	klines, err := client.NewKlinesService().Symbol(Symbol).Interval("1m").Do(context.Background())
	if err != nil {
		return nil, err
	}

	return klines, err
}

func (h *CoinsRateHandler) LastRate(Coin string, Currency string) (int64, string, error) {
	p, err := repo.Pair(h.dbs).FindByNames(Coin, Currency)
	if err != nil {
		return 0, "", err
	}

	k, err := repo.Kline(h.dbs).FindLast(p.Id)
	if err != nil {
		return 0, "", err
	}

	return k.CloseTime, k.Close, nil
}
