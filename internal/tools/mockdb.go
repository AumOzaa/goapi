package tools

import (
	"fmt"
	// "honnef.co/go/tools/printf"
	"time"
)

type mockDB struct {
}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "alex",
	},
	"jason": {
		AuthToken: "321ABC",
		Username:  "jason",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "alex",
	},
	"jason": {
		Coins:    200,
		Username: "jason",
	},
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	fmt.Printf("\nRUN-TIME : Rn the client data is %v\n", clientData)
	clientData, ok := mockLoginDetails[username]
	fmt.Printf("\nAnd now the client details is %v\n", clientData)
	if !ok {
		return nil
	}

	return &clientData
}

// func (d *mockDB) GetUserCoins(username string) *CoinDetails {
// 	time.Sleep(time.Second * 1)
//
// 	var clientData = CoinDetails{}
// 	clientData, ok := mockCoinDetails[username]
// 	if !ok {
// 		return nil
// 	}
//
// 	return &clientData
// }

func (d *mockDB) SetUpDatabase() error {
	return nil
}
