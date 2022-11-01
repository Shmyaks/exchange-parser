package repository

import (
	data "github.com/Shmyaks/exchange-parser-server/app/internal/data"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/filters"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/markets"
	"github.com/Shmyaks/exchange-parser-server/app/pkg/redis"

	jsoniter "github.com/json-iterator/go"
)

type P2PRepository struct {
	Datas           []data.P2P
	redisConnection redis.Connection
}

// NewP2PRepository fabric for CurrencyRepository
func NewP2PRepository(P2PDatas []data.P2P, redisConnection redis.Connection) *P2PRepository {
	mpDatas := make([]data.P2P, len(P2PDatas))
	for _, p2pData := range P2PDatas {
		mpDatas[*p2pData.GetMarketID()-1] = p2pData
	}

	return &P2PRepository{Datas: mpDatas, redisConnection: redisConnection}
}

// GetData get P2PData of repository
func (r *P2PRepository) GetData(m markets.P2PMarket) data.P2P {
	if len(r.Datas) <= int(m)-1 {
		panic("Datas not have this market")
	}
	return r.Datas[m-1]
}

// GetAllFromData method for get P2POrders from Data
func (r *P2PRepository) GetAllFromData(market markets.P2PMarket, p2pFilter filters.P2PFilter) ([]models.P2POrder, error) {
	return r.GetData(market).GetOrdersAPI(p2pFilter)
}

// Set method: Set P2Ppair to Redis
func (r *P2PRepository) Set(market markets.P2PMarket, pairs []models.P2PPair) error {
	bs, _ := jsoniter.Marshal(&pairs)
	cmd := r.redisConnection.Pool.HSet(market.GetName(), pairs[0].GetFullName(), bs)
	return cmd.Err()
}

// GetPayMethods get PayMethods from Data
func (r *P2PRepository) GetPayMethods(market markets.P2PMarket) (map[models.Fiat][]models.PayMethod, error) {
	return r.GetData(market).GetPayMethods(), nil
}

func (r *P2PRepository) GetFromCache(market markets.P2PMarket, pairName models.CurencyPairName) []models.P2PPair {
	println(string(pairName))
	strObj, _ := r.redisConnection.Pool.HGet(market.GetName(), string(pairName)).Result()
	println(strObj)
	pairs := []models.P2PPair{}
	jsoniter.UnmarshalFromString(strObj, &pairs)
	return pairs
}

func (r *P2PRepository) GetAllFromCache(market markets.P2PMarket) []models.P2PPair {
	strObj, _ := r.redisConnection.Pool.HGetAll(market.GetName()).Result()
	pairs := []models.P2PPair{}
	for _, v := range strObj {
		var pair []models.P2PPair
		jsoniter.UnmarshalFromString(v, &pair)
		pairs = append(pairs, pair...)
	}
	return pairs
}
