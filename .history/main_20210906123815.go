package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/go-redis/redis/v8"

)


type mainData struct {
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


func update (w http.ResponseWriter, r *http.Request) {
	var in mainData
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
	
	storeRedis(redisKeys(TopCryptos),&result)
	respondJSON(w, 200, "Succesful")

}



func get (w http.ResponseWriter, r *http.Request) {
	result, err := getRedis(redisKeys(TopCryptos))
	if err == redis.Nil {
        respondJSON(w, 500, "Key does  not exist")
    } else if err != nil {
        panic(err)
	}
	var data mainData
	err = json.Unmarshal([]byte(result), &data)
	if err != nil {
		respondJSON(w, 500, "Failed")
		return
	}
	respondJSON(w, 200, data)

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