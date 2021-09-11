package coin 

import (
	"github.com/farbod-bigdeli/app/redisHandler"
	"encoding/json"
	"errors"
)

type DetailedCoin struct {
	Coin
	HighestPrice1D 	float32 `json:"highestPrice1D"`
	LowestPrice1D 	float32	`json:"lowestPrice1D"`
	Volume1D 		float32	`json:"volume1D"`
	HighestBuyPrice float32	`json:"highestBuyPrice"`
	LowestSellPrice float32	`json:"lowestSellPrice"`
	DecimalLimits 	`json:"decimalLimits"`
}

type DecimalLimits struct{
	Buy 	DecimalLimit	`json:"buy"` 
	Sell 	DecimalLimit	`json:"sell"`			
}

type DecimalLimit struct {
	Coin 	int	`json:"coin"`
	Irt 	int	`json:"irt"`
	Usdt 	int	`json:"usdt"`
}
const (
	DetailedCoinKey = "detailed_coin_"
)



func GetDetailed(symbol string) (DetailedCoin, error){
	var detailedCoin DetailedCoin
	res, err := redisHandler.Get(DetailedCoinKey + symbol)
	if err == redisHandler.NotFound {
        return detailedCoin, errors.New("key not found")
    } else if err != nil {
        panic("Failed")
    }
	err = json.Unmarshal([]byte(res), &detailedCoin)
	if err != nil {
		return detailedCoin, errors.New("unmarshal failed")
	}
	return detailedCoin, nil

}

func (detaledCoin DetailedCoin) UpdateDetailed (key string) error {
	data, err := json.Marshal(detaledCoin)
	if err != nil {
		return errors.New("marshal failed")
	}
	redisHandler.Store(DetailedCoinKey + key, &data)

	return nil
}
