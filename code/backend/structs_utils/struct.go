package structs_utils

import (
	"github.com/dgrijalva/jwt-go"
)

type LoginRequest struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}
type Information struct {
}

type Claims struct {
	Id   string `json:"id"`
	Plan string `json:"plan"`
	jwt.StandardClaims
}

type ErroData struct {
	Email    string
	Password string
}

type IdAttendant struct {
	ID string `json:"id_attendant" binding:"required"`
}

type IdMessage struct {
	ID string `json:"id_message" binding:"required"`
}

type RequestInfo struct {
	IP        string
	UserAgent string
	Platform  string
	Query     map[string]string
	Body      string
	Referer   string
	Cookies   map[string]string
}

type IdAttendantActive struct {
	ID string `json:"id_attendant" binding:"required"`
	Active *bool `json:"active" binding:"required"`
}