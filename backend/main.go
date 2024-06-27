package main

import (
	"log"

	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/configs"
	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/middlewares"
	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/routes"
	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	utils.LoadEnv()
}

func main() {
	r := gin.Default()

	// Set up database connection
	db, err := configs.ConnectToDB()

	if err != nil {
		log.Fatal(err)
		return
	}

	// Allow cors
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowAllOrigins = true
	corsMiddleware := cors.New(corsConfig)
	r.Use(corsMiddleware)


	// Define root endpoint
	api := r.Group("/api")
	
	api.Use(middlewares.DBMiddleware(db.GetDB()))

	//Register routes
	routes.ReservationRoute(api)
	routes.ReviewRoute(api)
	routes.UserRoute(api)
	routes.AuthRoute(api)
	routes.BranchRoute(api)
	routes.ServiceRoute(api)

	r.Run(":3000")
}
