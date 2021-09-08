package main

import (
	"context"
	"encoding/json"
	// "fmt"
	"net/http"

	"github.com/farbod-bigdeli/app/coin"
	"github.com/gorilla/mux"
)



type CryptoBothPrice struct {
	Usdt []coin.Coin `json:"usdt"`
	Irt	 []coin.Coin `json:"irt"`
}



var ctx = context.Background()




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
	coinCollection.UpdateCollection(coin.TopCryptos)
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

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/update/top", updateTop).Methods("POST")
	r.HandleFunc("/update/all", updateTop).Methods("POST")
	r.HandleFunc("/update/detail/{sym}", update).Methods("POST")
    r.HandleFunc("/get/top", checkToken(getTop)).Methods("GET")
	r.HandleFunc("/get/all", checkToken(getAll)).Methods("GET")
	r.HandleFunc("/get/detail/{sym}", checkToken(getDetailed)).Methods("GET")
	
	http.ListenAndServe(":8000", r)
}