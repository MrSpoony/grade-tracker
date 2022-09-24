package cookie

import (
	"github.com/MrSpoony/grade-tracker/backend/logic/role"
	"github.com/MrSpoony/grade-tracker/backend/logic/user"

	"github.com/golang-jwt/jwt/v4"
)

var JwtKey = []byte("omcvoieansfuw0plqsntaisn2q;983l0-6lFPQ#$PL")

type Data struct {
	User  user.User   `json:"user"`
	Roles []role.Role `json:"roles"`
}

type Claims struct {
	Data Data
	jwt.StandardClaims
}
