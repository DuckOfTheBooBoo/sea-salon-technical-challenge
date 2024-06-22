package routes

import (
	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(route *gin.RouterGroup) {
	users := route.Group("/users")

	users.POST("/", controllers.UserCreate)
}
