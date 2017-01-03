package api

import (
	"github.com/dgrijalva/jwt-go"
	rdbModel "foresight-app.v1/backend/apps/models/rdb"
)

type JWTClaims struct {
	jwt.StandardClaims
	UserId          int64			`json:"user_id"`
	UserName	string			`json:"user_name"`
	Email		string			`json:"email"`
	Departures	[]rdbModel.Departure	`json:"departures"`
}

func (c JWTClaims) Valid() error {
	if err := c.StandardClaims.Valid(); err != nil {
		return err
	}

	return nil
}

type TokenRequest struct {
	Email		string 	`validate:"required,email" json:"email"`
	Password	string 	`validate:"required" json:"password"`
}