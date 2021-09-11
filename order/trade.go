package order

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/farbod-bigdeli/app/redisHandler"
)

type Trade struct {
	Order        
	Type string  `json:"type"`
}

type RecentTradesCollection struct {
	RecentTrades []Trade `json:"recentTrades"`
}

const (
	RecentTradesKey = "recent_trades_"
)

func GetRecentTrades(symbol string) (RecentTradesCollection, error){
	var recentTrades RecentTradesCollection
	res, err := redisHandler.Get(RecentTradesKey + symbol)
	if err == redisHandler.NotFound {
        return recentTrades, errors.New("key not found")
    } else if err != nil {
        panic("Failed")
    }
	err = json.Unmarshal([]byte(res), &recentTrades)
	if err != nil {
		return recentTrades, errors.New("unmarshal failed")
	}
	return recentTrades, nil

}

func (recentTrades RecentTradesCollection) UpdaterecentTrades(key string) error{

	data, err := json.Marshal(recentTrades)
	fmt.Println(string(data))
	if err != nil {
		return errors.New("marshal failed")
	}
	redisHandler.Store(RecentTradesKey + key, &data)

	return nil
}