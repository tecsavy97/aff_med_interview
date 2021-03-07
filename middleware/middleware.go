package middleware

import (
	"afford_meds_interview/routes"
	"log"

	"github.com/gin-gonic/gin"
)

// InitMiddleware - to initates jwt middleware
func InitMiddleware(g *gin.Engine) {
	o := &gin.RouterGroup{}
	o.Group("/o", OpenMiddleWare())
	r := &gin.RouterGroup{}
	r.Group("/r", RestrictedMiddleWare())
	routes.InitRoutes(o, r)
}

// OpenMiddleWare - For open routes with context "/o"
func OpenMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Open Middleware Used!!")
	}
}

// RestrictedMiddleWare - For restricted routes with context "/r"
func RestrictedMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// token := c.GetHeader("Authorization")
	}
}
