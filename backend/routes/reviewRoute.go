package routes

import (
	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/controllers"
	"github.com/gin-gonic/gin"
)

func ReviewRoute(route *gin.RouterGroup) {
	review := route.Group("/review")

	review.GET("/", controllers.ReviewGet)
	review.POST("/", controllers.ReviewCreate)
}
