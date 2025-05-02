package token_app

import (
	"code/hash_app"
	"code/structs_utils"
	"code/utils"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	JWT_KEY []byte
)

func ValidateJWT(tokenString string) (*structs_utils.Claims, error) {
	claims := &structs_utils.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return JWT_KEY, nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("token inválido")
	}
	return claims, nil
}


func GenerateJWT(data_of_user bson.M, c *gin.Context) error {
	expiration_time := time.Now().Add(12 * time.Hour)
	claims := &structs_utils.Claims{
		Id: data_of_user["id"].(string),
		Plan: hash_app.HashString(data_of_user["plan"].(string)),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration_time.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWT_KEY)
	if err != nil {
		return err
	}

	expiration_time_in_second := strconv.Itoa(int(time.Until(expiration_time).Seconds()))

	c.Header("Set-Cookie", fmt.Sprintf("jwt_token=%s; Path=/; Domain=%s; Max-Age=%s; Secure; HttpOnly; SameSite=Strict; Priority=High", tokenString, utils.GetDomain(c), expiration_time_in_second))

	return nil
}

func TokenAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString, _ := c.Cookie("jwt_token")
        if tokenString == "" {
            c.Redirect(http.StatusFound, "/entrar")
            c.Abort()
            return
        }

        claims, err := ValidateJWT(tokenString)

        if err != nil {
			c.SetCookie("jwt_token", "", -1, "/", utils.GetDomain(c), false, true)
            c.Redirect(http.StatusFound, "/entrar")
            c.Abort()
            return
        }

        
		

        c.Set("id", claims.Id)
        c.Set("plan", claims.Plan)
        c.Next()
    }
}


func GenerateUUIDFromEmail(email string) uuid.UUID {
	
	return uuid.NewSHA1(uuid.NameSpaceDNS, []byte(email))
}