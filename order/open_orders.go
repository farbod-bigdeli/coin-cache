package order

import (
	"encoding/json"
	"errors"

	"github.com/farbod-bigdeli/app/redisHandler"
)

type OpenOrders struct {
	OpenSellOrders []Order `json:"openSellOrders"`
	OpenBuyOrders  []Order `json:"openBuyOrders"` 
}


const (
	OpenOrdersKey = "open_order_"
)

func GetOpenOrders(symbol string) (OpenOrders, error){
	var openOrders OpenOrders
	res, err := redisHandler.Get(OpenOrdersKey + symbol)
	if err == redisHandler.NotFound {
        return openOrders, errors.New("key not found")
    } else if err != nil {
        panic("Failed")
    }
	err = json.Unmarshal([]byte(res), &openOrders)
	if err != nil {
		return openOrders, errors.New("unmarshal failed")
	}
	return openOrders, nil

}

func (openOrders OpenOrders) UpdateOpenOrders(key string) error{

	data, err := json.Marshal(openOrders)
	if err != nil {
		return errors.New("marshal failed")
	}
	redisHandler.Store(OpenOrdersKey + key, &data)

	return nil
}
