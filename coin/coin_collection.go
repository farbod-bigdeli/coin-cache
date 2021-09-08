package coin

import(
	"github.com/farbod-bigdeli/app/redisHandler"
	"errors"
	"encoding/json"
)

type CoinCollection struct {
	Usdt []Coin `json:"usdt"`
	Irt	 []Coin `json:"irt"`
}

const (
	TopCryptos string = "top_five_cryptos"
	AllCryptos string= "all_cryptos"	
)

func GetCollection(symbol string) (CoinCollection, error){
	var coinCollection CoinCollection
	res, err := redisHandler.Get(symbol)
	if err == redisHandler.NotFound {
        return coinCollection, errors.New("key not found")
    } else if err != nil {
        panic("Failed")
    }
	err = json.Unmarshal([]byte(res), &coinCollection)
	if err != nil {
		return coinCollection, errors.New("unmarshal failed")
	}
	return coinCollection, nil

}


func (coinCollection CoinCollection) UpdateCollection(key string) error{

	data, err := json.Marshal(coinCollection)
	if err != nil {
		return errors.New("marshal failed")
	}
	redisHandler.Store(key, &data)

	return nil
}