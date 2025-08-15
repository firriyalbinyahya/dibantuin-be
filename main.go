package main

import (
	"dibantuin-be/config"
	"dibantuin-be/config/database"
	"dibantuin-be/config/redis"
	"dibantuin-be/entity"
	"dibantuin-be/routes"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func InitializeApp() *gin.Engine {
	config.InitConfig()

	r := gin.Default()

	redis.ConnectRedis()
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

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000" // Default port
	}
	fmt.Println("Server is running on port " + port)
	app.Run(":" + port)
}
