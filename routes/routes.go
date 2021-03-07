package routes

import (
	"afford_meds_interview/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(o, r *gin.RouterGroup) {
	o.POST("/login", handlers.LoginHandler())
	o.POST("/register", handlers.RegisterHandler())
}
