package cookie

import (
	"github.com/golang-jwt/jwt/v4"

	"github.com/MrSpoony/grade-tracker/backend/logic/role"
)

var JwtKey = []byte("omcvoieansfuw0plqsntaisn2q;983l0-6lFPQ#$PL")

type Claims struct {
	Username string      `json:"username"`
	Roles    []role.Role `json:"roles"`
	jwt.StandardClaims
}
