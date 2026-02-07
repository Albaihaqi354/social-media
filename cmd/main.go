package main

import (
	"log"

	"github.com/Albaihaqi354/FinalPhase3.git/internal/config"
	"github.com/Albaihaqi354/FinalPhase3.git/internal/middleware"
	"github.com/Albaihaqi354/FinalPhase3.git/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"
)

// @title           Media Sodial Application
// @version         1.0
// @description     Back End using go with gin.
// @BasePath 		/

// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization
// @description                Type "Bearer" followed by a space and then your token.

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Failed to load env")
	}

	db, err := config.InitDb()
	if err != nil {
		log.Println("Failed to Connect Database")
		return
	}
	defer db.Close()

	rdb := config.InitRedis()
	defer rdb.Close()

	app := gin.Default()
	app.Use(middleware.CORSMiddleware)
	router.Init(app, db, rdb)
	app.Run(":5051")
}
