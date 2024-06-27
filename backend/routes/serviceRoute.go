package routes

import (
	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/controllers"
	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/middlewares"
	"github.com/gin-gonic/gin"
)

func ServiceRoute(route *gin.RouterGroup) {
	services := route.Group("/services")

	services.GET("", middlewares.JWTMiddleware(), controllers.ServiceGetAll)
	services.POST("", middlewares.JWTMiddleware(), controllers.ServiceCreate)
}
