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
}



func GetDetailed(symbol string) (DetailedCoin, error){
	var detailedCoin DetailedCoin
	res, err := redisHandler.Get(symbol)
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