package main

import (
	"encoding/json"
	// "fmt"
	"net/http"

	"github.com/farbod-bigdeli/app/coin"
	"github.com/farbod-bigdeli/app/order"
	"github.com/gorilla/mux"
)



type CryptoBothPrice struct {
	Usdt []coin.Coin `json:"usdt"`
	Irt	 []coin.Coin `json:"irt"`
}



func getTop(w http.ResponseWriter, r *http.Request) {
	detailedCoin, err := coin.GetCollection(coin.TopCryptos)
	if err != nil {
		respondJSON(w, 200, err.Error())
		return
	}
	respondJSON(w, 200, detailedCoin)

}


func getAll(w http.ResponseWriter, r *http.Request) {
	detailedCoin, err := coin.GetCollection(coin.AllCryptos)
	if err != nil {
		respondJSON(w, 200, err.Error())
		return
	}
	respondJSON(w, 200, detailedCoin)

}

func getDetailed(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["sym"]
	detailedCoin, err := coin.GetDetailed(symbol)
	if err != nil {
		respondJSON(w, 200, err.Error())
		return
	}
	respondJSON(w, 200, detailedCoin)
}

func getOpenOrders(w http.ResponseWriter, r *http.Request) {
	sym := mux.Vars(r)["sym"]
	openOrders, err := order.GetOpenOrders(sym)
	if err != nil {
		respondJSON(w, 200, err.Error())
		return
	}
	respondJSON(w, 200, openOrders)

}
func getRecentTrades(w http.ResponseWriter, r *http.Request) {
	sym := mux.Vars(r)["sym"]
	recentTrades, err := order.GetRecentTrades(sym)
	if err != nil {
		respondJSON(w, 200, err.Error())
		return
	}
	respondJSON(w, 200, recentTrades)

}

func updateTop (w http.ResponseWriter, r *http.Request) {
	var coinCollection coin.CoinCollection
	err := json.NewDecoder(r.Body).Decode(&coinCollection)
	if err != nil {
        respondJSON(w, 500, "Decode failed")
		return
    }
	coinCollection.UpdateCollection(coin.TopCryptos)
	respondJSON(w, 200, "Successful")

}

func updateAll (w http.ResponseWriter, r *http.Request) {
	var coinCollection coin.CoinCollection
	err := json.NewDecoder(r.Body).Decode(&coinCollection)
	if err != nil {
        respondJSON(w, 500, "Decode failed")
		return
    }
	coinCollection.UpdateCollection(coin.AllCryptos)
	respondJSON(w, 200, "Successful")

}


func update(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["sym"]
	var detailedCoin coin.DetailedCoin
	err := json.NewDecoder(r.Body).Decode(&detailedCoin)
	if err != nil {
        respondJSON(w, 500, "Decode failed")
		return
    }
	err = detailedCoin.UpdateDetailed(symbol)
	if err != nil {
		respondJSON(w, 500, "Failed")
		return
	}
	respondJSON(w, 200, "Successful")
}



func updateOpenOrders(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["sym"]
	var openOrders order.OpenOrders
	err := json.NewDecoder(r.Body).Decode(&openOrders)
	if err != nil {
        respondJSON(w, 500, "Decode failed")
		return
    }
	err = openOrders.UpdateOpenOrders(symbol)
	if err != nil {
		respondJSON(w, 500, "Failed")
		return
	}
	respondJSON(w, 200, "Successful")
}

func updateRecentTrades(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["sym"]
	var recentTrades order.RecentTradesCollection
	err := json.NewDecoder(r.Body).Decode(&recentTrades)
	if err != nil {
        respondJSON(w, 500, "Decode failed")
		return
    }
	err = recentTrades.UpdaterecentTrades(symbol)
	if err != nil {
		respondJSON(w, 500, "Failed")
		return
	}
	respondJSON(w, 200, "Successful")
}

func respondJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}


func checkToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	  header := r.Header
	  
	  if _, ok := header["Token"]; ok{
		next(w, r)
		return
	  }
	  respondJSON(w, 401, "Unauthorized.")
	}
}

// func test (w http.ResponseWriter, r *http.Request) {
// 	a := order.RecentTradesCollection{}
// 	b := order.Trade {}
// 	b.Amount = 23
// 	a.RecentTrades = append(a.RecentTrades, b)
// 	respondJSON(w, 200, a)
// }

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/update/top", updateTop).Methods("POST")
	r.HandleFunc("/update/all", updateAll).Methods("POST")
	r.HandleFunc("/update/detail/{sym}", update).Methods("POST")
    r.HandleFunc("/get/top", checkToken(getTop)).Methods("GET")
	r.HandleFunc("/get/all", checkToken(getAll)).Methods("GET")
	r.HandleFunc("/get/detail/{sym}", checkToken(getDetailed)).Methods("GET")
	r.HandleFunc("/order/get/open/{sym}", checkToken(getOpenOrders)).Methods("GET")
	r.HandleFunc("/order/update/open/{sym}", checkToken(updateOpenOrders)).Methods("POST")
	r.HandleFunc("/trade/get/recent/{sym}", checkToken(getRecentTrades)).Methods("GET")
	r.HandleFunc("/trade/update/recent/{sym}", checkToken(updateRecentTrades)).Methods("POST")
	r.HandleFunc("/trade/update/recent/{sym}", checkToken(updateRecentTrades)).Methods("POST")
	// r.HandleFunc("/test", checkToken(test)).Methods("GET")

	
	http.ListenAndServe(":8000", r)
}