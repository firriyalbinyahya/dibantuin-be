package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type JWT struct {
	AccessSecret       string
	RefreshSecret      string
	AccessExpiryInSec  int
	RefreshExpiryInSec int
}

var Config struct {
	JWT JWT
}

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, reading env variables directly")
	}

	accessExp, _ := strconv.Atoi(os.Getenv("ACCESS_EXPIRY_IN_SEC"))
	refreshExp, _ := strconv.Atoi(os.Getenv("REFRESH_EXPIRY_IN_SEC"))

	Config.JWT = JWT{
		AccessSecret:       os.Getenv("ACCESS_SECRET"),
		RefreshSecret:      os.Getenv("REFRESH_SECRET"),
		AccessExpiryInSec:  accessExp,
		RefreshExpiryInSec: refreshExp,
	}
}
