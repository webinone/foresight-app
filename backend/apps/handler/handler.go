package handler

import (
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"foresight-app.v1/backend/apps/models/api"
	"net/http"
)

func JSONHTTPErrorHandler(err error, c echo.Context) {
	code := fasthttp.StatusInternalServerError
	msg := "Internal Server Error"
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
	}

	APIResultHandler(c, false, code, msg)
}

// transaction middleware
func TransactionHandler(db *gorm.DB) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {

			tx := db.Begin()

			c.Set("Tx", tx)

			if err := next(c); err != nil {
				tx.Rollback()
				logrus.Debug("Transction Rollback: ", err)
				return err
			}
			logrus.Debug("Transaction Commit")
			tx.Commit()

			return nil
		})
	}
}

func APIResultHandler(c echo.Context, httpSuccess bool,  httpStatus int, data interface{}) error {

	apiResult := api.APIResult{
		Success : httpSuccess,
		ResultCode : httpStatus,
		ResultData: data,
	}

	//fmt.Println(apiResult)
	//
	//// CORS 설정
	//ctx.Response.Header.Add("Access-Control-Allow-Origin", "*")
	//ctx.Response.Header.Add("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS, DELETE")
	//ctx.Response.Header.Add("Access-Control-Max-Age", "3600")
	//ctx.Response.Header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With, Origin")
	//
	//return ctx.JSON(iris.StatusOK, &apiResult)
	return c.JSON(http.StatusOK, &apiResult)
}


//func setUpRequest ( c echo.HandlerFunc ) echo.HandlerFunc {
//	return func(ctx echo.Context) error {
//		req := ctx.Request()
//
//		Logger := logger.NewLogger()
//		// add some default fields to the logger ~ on all messages
//		logger := api.log.WithFields(logrus.Fields{
//			"method":     req.Method(),
//			"path":       req.URL().Path(),
//			"request_id": uuid.NewRandom().String(),
//		})
//		ctx.Set(loggerKey, logger)
//		startTime := time.Now()
//
//		defer func() {
//			rsp := ctx.Response()
//			// at the end we will want to log a few more interesting fields
//			logger.WithFields(logrus.Fields{
//				"status_code":  rsp.Status(),
//				"runtime_nano": time.Since(startTime).Nanoseconds(),
//			}).Info("Finished request")
//		}()
//
//		// now we will log out that we have actually started the request
//		logger.WithFields(logrus.Fields{
//			"user_agent":     req.UserAgent(),
//			"content_length": req.ContentLength(),
//		}).Info("Starting request")
//
//		err := f(ctx)
//		if err != nil {
//			ctx.Error(err)
//		}
//
//		return err
//	}
//}