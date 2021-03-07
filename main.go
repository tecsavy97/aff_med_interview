package main

import (
	"afford_meds_interview/helpers/confighelper"
	"afford_meds_interview/middleware"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	if initeErr := confighelper.InitConfig(); initeErr != nil {
		log.Println("Failed to init Config")
	}
	if startErr := startServer(); startErr != nil {
		log.Println(startErr)
	}
}

func startServer() error {
	router := gin.Default()
	s := &http.Server{
		Handler: router,
		Addr:    ":7000",
	}
	router.GET("/", checkStatus())
	middleware.InitMiddleware(router)
	if serveErr := s.ListenAndServe(); serveErr != nil {
		return serveErr
	}
	return nil
}

func checkStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Server is Running")
		c.AbortWithStatus(http.StatusOK)
	}
}
