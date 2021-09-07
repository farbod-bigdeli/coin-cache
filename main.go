package main

import (
	"context"
	"encoding/json"
	// "fmt"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)


type Coin struct {
    Symbol  	string 	`json:"symbol"`
    FaName 		string 	`json:"faName"`
	Price 		float32 `json:"price"`	
	Change1D    float32	`json:"change1d"`
	Chart		string  `json:"chart"`
}

type DetailedCoin struct {
	Coin
	HighestPrice1D 	float32 `json:"highestPrice1D"`
	LowestPrice1D 	float32	`json:"lowestPrice1D"`
	Volume1D 		float32	`json:"volume1D"`
	HighestBuyPrice float32	`json:"highestBuyPrice"`
	LowestSellPrice float32	`json:"lowestSellPrice"`
}

type CryptoBothPrice struct {
	Usdt []Coin `json:"usdt"`
	Irt	 []Coin `json:"irt"`
}



var ctx = context.Background()




func getTop(w http.ResponseWriter, r *http.Request) {
	result, err := getRedis(TopCryptos)
	if err == redis.Nil {
        respondJSON(w, 500, "Key does  not exist")
		return
    } else if err != nil {
        panic(err)
	}
	var data CryptoBothPrice
	err = json.Unmarshal([]byte(result), &data)
	if err != nil {
		respondJSON(w, 500, "Failed")
		return
	}
	respondJSON(w, 201, data)

}


func getAll(w http.ResponseWriter, r *http.Request) {
	result, err := getRedis(AllCryptos)
	if err == redis.Nil {
        respondJSON(w, 500, "Key does  not exist")
		return
    } else if err != nil {
        panic(err)
	}
	var data CryptoBothPrice
	err = json.Unmarshal([]byte(result), &data)
	if err != nil {
		respondJSON(w, 500, "Failed")
		return
	}
	respondJSON(w, 201, data)

}

func get(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["sym"]
	result, err := getRedis(symbol)
	if err == redis.Nil {
        respondJSON(w, 500, "Key does  not exist")
		return
    } else if err != nil {
        panic(err)
	}
	var data DetailedCoin
	err = json.Unmarshal([]byte(result), &data)
	if err != nil {
		respondJSON(w, 500, "Failed")
		return
	}
	respondJSON(w, 201, data)

}


func updateTop (w http.ResponseWriter, r *http.Request) {
	var in CryptoBothPrice
	err := json.NewDecoder(r.Body).Decode(&in)
    
	if err != nil {
        respondJSON(w, 500, "Decode failed")
		return
    }

	result, err :=json.Marshal(in)
	if err != nil {
		respondJSON(w, 500, "Marshal failed")
		return
	}
	
	storeRedis(TopCryptos,&result)
	respondJSON(w, 200, "Succesful")

}

func update(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["sym"]
	var in DetailedCoin
	err := json.NewDecoder(r.Body).Decode(&in)
    
	if err != nil {
        respondJSON(w, 500, "Decode failed")
		return
    }

	result, err :=json.Marshal(in)
	if err != nil {
		respondJSON(w, 500, "Marshal failed")
		return
	}
	
	storeRedis(symbol, &result)
	respondJSON(w, 200, "Succesful")

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
	r.HandleFunc("/update/detail/{sym}", update).Methods("POST")
    r.HandleFunc("/get/top", checkToken(getTop)).Methods("GET")
	r.HandleFunc("/get/all", checkToken(getAll)).Methods("GET")
	r.HandleFunc("/get/detail/{sym}", checkToken(get)).Methods("GET")
	
	// r.HandleFunc("/get/all", checkToken(get)).Methods("GET")

	http.ListenAndServe(":8000", r)
}