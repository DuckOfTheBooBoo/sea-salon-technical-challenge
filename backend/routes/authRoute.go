package routes

import (
	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoute(route *gin.RouterGroup) {
	auth := route.Group("/auth")

	auth.POST("/login", controllers.AuthLogIn)
}
