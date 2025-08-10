package main

import (
	"dibantuin-be/config"
	"dibantuin-be/config/database"
	"dibantuin-be/config/redis"
	"dibantuin-be/entity"
	"dibantuin-be/routes"

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
	app.Run(":8080")
}
