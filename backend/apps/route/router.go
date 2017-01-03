package route

import (
	"github.com/labstack/echo"
	echoMw "github.com/labstack/echo/middleware"
	"foresight-app.v1/backend/apps/api"
	apiModel "foresight-app.v1/backend/apps/models/api"
	"foresight-app.v1/backend/apps/handler"
	"foresight-app.v1/backend/apps/db"
	"foresight-app.v1/backend/apps/libs"
)

func Init() *echo.Echo {

	e := echo.New()

	// Set Bundle MiddleWare
	e.Use(echoMw.Logger())
	e.Use(echoMw.Recover())
	e.Use(echoMw.Gzip())
	e.Use(echoMw.CORSWithConfig(echoMw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))

	// JWT
	jwtMw := echoMw.JWTWithConfig(echoMw.JWTConfig{
		SigningKey: []byte(libs.Config.AUTH.JwtKey),
		ContextKey: "",
		Claims:&apiModel.JWTClaims{},
		//TokenLookup: "header:token" ,
	})

	// Error Handler Centralized
	e.HTTPErrorHandler = handler.JSONHTTPErrorHandler

	// DB 초기화 GORM (MySql)
	rdb := db.MySqlDBObj.Init()
	e.Use(handler.TransactionHandler(rdb))

	e.GET("/", 		api.HelloAPIObj.GetHello())
	e.POST("/login",  	api.AuthAPIObj.PostLogin())

	// Simple Board
	e.POST("/board", 		api.SimpleBoardAPIObj.CreateBoard())
	e.GET("/board",  		api.SimpleBoardAPIObj.GetBoards())
	e.GET("/board/:id", 		api.SimpleBoardAPIObj.GetBoard())
	e.DELETE("/board/:id", 		api.SimpleBoardAPIObj.DeleteBoard())
	e.PUT("/board/:id", 		api.SimpleBoardAPIObj.PutBoard())

	// Routes (JWT)
	v1 := e.Group("/api/v1", jwtMw)
	{
		// CommentBoard
		v1.POST("/comment_board", 		api.CommentBoardAPIObj.PostCommentBoard())
		v1.PUT("/comment_board", 		api.CommentBoardAPIObj.PutCommentBoard())
		v1.DELETE("/comment_board/:id", 	api.CommentBoardAPIObj.DeleteCommentBoard())

		v1.GET("/comment_board/:id", 		api.CommentBoardAPIObj.GetCommentBoard())
		v1.GET("/comment_board", 		api.CommentBoardAPIObj.GetCommentBoards())
	}

	return e
}