package main

import (
	"dibantuin-be/config/database"
	"dibantuin-be/entity"
	"dibantuin-be/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitializeApp() *gin.Engine {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading ENV")
	}

	r := gin.Default()
	db := database.ConnectDatabase()

	//auto migrate
	db.AutoMigrate(&entity.Category{}, &entity.DonationProgram{}, &entity.DonationProgramRequest{},
		&entity.DonationReport{}, &entity.MoneyTransactionDonation{}, &entity.User{},
		&entity.UserLog{}, &entity.VerificationProgram{}, &entity.VerificationTransactionDonation{})

	routes.SetupRoutes(r, db)

	return r
}

func main() {
	app := InitializeApp()
	app.Run(":8080")
}
