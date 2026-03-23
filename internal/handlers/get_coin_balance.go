package handlers

import (
	"encoding/json"
	// "go/token"
	"net/http"

	"github.com/AumOzaa/goapi/api"
	"github.com/AumOzaa/goapi/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"

	"fmt"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\nIn teh get coin balance")
	var params = api.CoinBalanceParams{} // getting the values from
	fmt.Printf("The value of params is %v", params)

	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query()) // Grab the params from the url, and set the to the values in the struct

	fmt.Printf("\nRuntime : The value of decoder is : %v", err)

	if err != nil { // woild be use when there is an error in decoding the params
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.CoinBalanceResponse{
		Balance: (*&tokenDetails).Coins,
		Code:    http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
