package tokenhelper

import (
	"afford_meds_interview/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type jwtClaims struct {
	Username string `json:"userName"`
	jwt.StandardClaims
}

// GenerateToken - to generate jwt token for authenctication of restricted routes
func GenerateToken(username string) (string, error) {
	expirationtime := time.Hour * 24
	claims := jwtClaims{
		Username: username,
	}
	claims.ExpiresAt = int64(expirationtime)
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString([]byte(""))
	if err != nil {
		return "", err
	}
	return tokenString, err
}

// GetTokenForLogin - to validate
func GetTokenForLogin(c *gin.Context, username string) {
	token := c.GetHeader("Authorization")
	DecodeToken(token, "", models.AppConfig.JwtKey)
}

func DecodeToken(token, username, jwtKey string) error {
	var claims jwtClaims
	claims.Username = username
	valToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {

		return []byte(jwtKey), nil
	})
	if err != nil {
		return err
	}
	valErr := valToken.Claims.Valid()
	if valErr != nil {
		return valErr
	}
	return nil
}
