package main

import (
	"context"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)


type input struct {
	Usdt []crypto `json:"usdt"`
	Irt	[]crypto	`json:"irt"`
}

type crypto struct {
    Symbol  	string 	`json:"symbol"`
    PersianName string 	`json:"persianName"`
	Price 		float32 `json:"price"`	
	Change1D    float32	`json:"change1d"`
	Chart		string  `json:"chart"`
}

var ctx = context.Background()

type redisKeys string

const (
	TopCryptos redisKeys = "top_five_cryptos"
	AllCryptos redisKeys= "all_cryptos"	
)




func update (w http.ResponseWriter, r *http.Request) {
	var in input
	err := json.NewDecoder(r.Body).Decode(&in)
    
	if err != nil {
        respondJSON(w, 500, "Decode failed")
    }

	result, err :=json.Marshal(in)
	if err != nil {
		respondJSON(w, 500, "Marshal failed")
	}
	
	storeRedis(redisKeys(TopCryptos),&result)
	respondJSON(w, 500, "Succesful")

}



func get (w http.ResponseWriter, r *http.Request) {
	result := getRedis(redisKeys(TopCryptos))
	var in input
	err := json.Unmarshal([]byte(result.Val()), &in)
	if err != nil {
		respondJSON(w, 500, "Failed")
	}
	respondJSON(w, 200, in)

}

func respondJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}
func main() {
    r := mux.NewRouter()
    r.HandleFunc("/update/top", update).Methods("POST")
    r.HandleFunc("/get/top", get).Methods("GET")
	http.ListenAndServe(":8000", r)
}