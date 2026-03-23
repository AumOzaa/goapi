package tools

import (
	log "github.com/sirupsen/logrus"
	// "github.com/stretchr/testify/mock"
)

type LoginDetails struct {
	AuthToken string
	Username  string
}

type CoinDetails struct {
	Coins    int64
	Username string
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetUpDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {
	log.Printf("\nRUN-TIME: The new database is initialized\n")

	var database DatabaseInterface = &mockDB{}
	log.Printf("\nThe value of the current database is : %v\n", database)

	var err error = database.SetUpDatabase()

	log.Printf("\nRUN-TIME: The value of the err right now in the database is %v\n", err)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &database, nil
}
