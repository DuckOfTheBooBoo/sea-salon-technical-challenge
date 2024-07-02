package routes

import (
	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/controllers"
	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/middlewares"
	"github.com/gin-gonic/gin"
)

func BranchRoute(route *gin.RouterGroup) {
	branches := route.Group("/branches")

	branches.GET("", middlewares.JWTMiddleware(), controllers.BranchGetAll)
	branches.POST("", middlewares.JWTMiddleware(), controllers.BranchCreate)
	branches.PUT("/:id", middlewares.JWTMiddleware(), controllers.BranchUpdate)
	branches.DELETE("/:id", middlewares.JWTMiddleware(), controllers.BranchDelete)
}
