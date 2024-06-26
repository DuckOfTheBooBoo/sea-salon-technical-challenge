package routes

import (
	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/controllers"
	"github.com/DuckOfTheBooBoo/sea-salon-technical-challenge/backend/middlewares"
	"github.com/gin-gonic/gin"
)

func BranchRoute(route *gin.RouterGroup) {
	branches := route.Group("/branches")

	branches.POST("", middlewares.JWTMiddleware(), controllers.BranchCreate)
}
