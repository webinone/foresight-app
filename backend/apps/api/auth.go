package api

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/Sirupsen/logrus"
	"foresight-app.v1/backend/apps/libs"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/dgrijalva/jwt-go"
	apiModel "foresight-app.v1/backend/apps/models/api"
	rdbModel "foresight-app.v1/backend/apps/models/rdb"
	"foresight-app.v1/backend/apps/handler"
	"gopkg.in/go-playground/validator.v9"
)

type AuthAPI struct {
	Logger 		*logrus.Logger
	Validator 	*validator.Validate
}

var AuthAPIObj = &AuthAPI{
	Logger : libs.GetLogger(),
	Validator : validator.New(),
}

func (api AuthAPI) PostLogin() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		//logger.Info("Hello , API !!!")

		// Login 정보를 받고
		payload := &apiModel.TokenRequest{}
		c.Bind(payload)

		if err = api.Validator.Struct(payload); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.(validator.ValidationErrors))
		} else {
			api.Logger.Debug(" success validate !!!!")
		}

		// 암호전 패스워드
		fmt.Println(">>>> general pass : " , payload.Password)

		enc_pass := libs.Sha256Encoding(payload.Password)
		payload.Password = enc_pass

		fmt.Println(">>>> enc_pass : " , enc_pass)

		user := &rdbModel.User{
			Email:payload.Email,
			Password:enc_pass,
		}

		// 사용자가 있는지 확인한다.
		var count = 0
		tx := c.Get("Tx").(*gorm.DB)
		tx.Preload("Departures").Where("email = ? and password = ? ", user.Email, user.Password).Find(user).Count(&count)

		if count == 0 {
			// 존재하지 않는다면...
			return echo.NewHTTPError(http.StatusUnauthorized, "email or password not matched")
		}

		// JWT KEY 값을 전송한다.
		fmt.Println(">>>> user_id : " , user.Id)
		fmt.Println(">>>> user_password : " , user.Password)
		fmt.Println(">>>> user_email : " , user.Email)

		// JWT Key Get
		jwtKey := []byte(libs.Config.AUTH.JwtKey)
		token := jwt.New(jwt.SigningMethodHS256)

		claims := &apiModel.JWTClaims{
			UserId:user.Id,
			UserName:user.Name,
			Email:user.Email,
			Departures:user.Departures,
		}

		token.Claims = claims
		tokenString, _ := token.SignedString(jwtKey)

		fmt.Println("tokenString : ", tokenString)

		return handler.APIResultHandler(c, true, http.StatusOK, map[string]interface{}{"token_key": tokenString})
	}
}
