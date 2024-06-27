package routes

import (
	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/controllers"
	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/middlewares"
	"github.com/gin-gonic/gin"
)

func BranchRoute(route *gin.RouterGroup) {
	branches := route.Group("/branches")

	branches.GET("", controllers.BranchGetAll)
	branches.POST("", middlewares.JWTMiddleware(), controllers.BranchCreate)
	branches.DELETE("/:id", middlewares.JWTMiddleware(), controllers.BranchDelete)
}
