package api

import (
	"github.com/labstack/echo"
	"net/http"
	"foresight-app.v1/backend/apps/libs"
	"github.com/Sirupsen/logrus"
)

type HelloAPI struct {
	Logger *logrus.Logger
}

var HelloAPIObj = &HelloAPI{
	Logger : libs.GetLogger("api-hello"),
}

// INDEX PAGE
func (api HelloAPI) GetHello() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		api.Logger.Info("Hello , API !!!")

		// ERROR 발생
		//return echo.NewHTTPError(http.StatusBadRequest, "Must provide both email and password")
		return c.String(http.StatusOK, "Hello , API !!!")
	}
}