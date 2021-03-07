package handlers

import (
	"afford_meds_interview/models"
	"afford_meds_interview/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if bindErr := c.Bind(&user); bindErr != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"result": "Data not in Json Format", "err": bindErr.Error()})
			return
		}
		token, loginErr := services.UserLogin(user)
		if loginErr != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"result": "", "err": loginErr.Error()})
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"result": "User Login Sucessfully, Please use the Token:" + token, "err": nil})
	}
}

func RegisterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if bindErr := c.Bind(&user); bindErr != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"result": "Data not in Json Format", "err": bindErr.Error()})
			return
		}
		err := services.UserRegistration(user)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"result": nil, "err": err.Error()})
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"result": "User Registered SuccesFully!!"})
	}
}
