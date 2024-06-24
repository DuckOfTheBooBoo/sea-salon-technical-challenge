package routes

import (
	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/controllers"
	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/middlewares"
	"github.com/gin-gonic/gin"
)

func ReservationRoute(route *gin.RouterGroup) {
	reservation := route.Group("/reservations")

	reservation.POST("", middlewares.JWTMiddleware(), controllers.ReservationCreate)
}
