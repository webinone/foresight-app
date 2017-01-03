package api

import (
	"github.com/Sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	"foresight-app.v1/backend/apps/libs"
	"github.com/labstack/echo"
	rdbModel "foresight-app.v1/backend/apps/models/rdb"
	"net/http"
	"github.com/jinzhu/gorm"
	"foresight-app.v1/backend/apps/handler"
	"fmt"
	"strconv"
)

type CommentBoardAPI struct {
	Logger 		*logrus.Logger
	Validator 	*validator.Validate
}

var CommentBoardAPIObj = &CommentBoardAPI{
	Logger 		: libs.GetLogger("api-commentboard"),
	Validator	: validator.New(),
}

// CommentBoard Insert
func (api CommentBoardAPI) PostCommentBoard() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		board := &rdbModel.CommentBoard{}

		c.Bind(board)

		claims 		:= libs.GetJWTClaims(c)
		user_id 	:= claims.UserId
		user_name 	:= claims.UserName
		email 		:= claims.Email
		departures 	:= claims.Departures

		api.Logger.Debug(" user_id : ", user_id)
		api.Logger.Debug(" user_name : ", user_name)
		api.Logger.Debug(" email : ", email)
		api.Logger.Debug(" departures : ", departures)

		board.UserId = user_id

		// validate 체크
		if err = api.Validator.Struct(board); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.(validator.ValidationErrors))
		} else {
			api.Logger.Debug(" success validate !!!!")
		}

		tx := c.Get("Tx").(*gorm.DB)
		tx.Create(board)

		// ERROR 발생
		//return echo.NewHTTPError(http.StatusBadRequest, "Must provide both email and password")
		return handler.APIResultHandler(c, true, http.StatusCreated, "Created Board")
	}
}

// CommentBoard Update
func (api CommentBoardAPI) PutCommentBoard() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		board := &rdbModel.CommentBoard{}

		c.Bind(board)

		claims 		:= libs.GetJWTClaims(c)
		user_id 	:= claims.UserId

		board.UserId = user_id

		tx := c.Get("Tx").(*gorm.DB)
		//tx.Model(board).Where("id = ?", board.Id).Updates(
		count := tx.Model(board).Updates(
			map[string]interface{}{
				"title": board.Title,
				"content": board.Content,
				"user_id" : board.UserId,
			}).RowsAffected

		fmt.Println(">>> count : ", count)

		if count == 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "NOT UPDATE")
		}

		return handler.APIResultHandler(c, true, http.StatusOK, "Update Board")
	}
}

// CommentBoard DELETE (Delete)
func (api CommentBoardAPI) DeleteCommentBoard() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		id, _ := strconv.ParseInt(c.Param("id"), 0, 64)
		board := &rdbModel.CommentBoard{}

		tx := c.Get("Tx").(*gorm.DB)
		count := tx.Delete(board, "id = ? ", id).RowsAffected

		if count == 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "NOT DELETE")
		}

		// ERROR 발생
		//return echo.NewHTTPError(http.StatusBadRequest, "Must provide both email and password")
		//return c.JSON(http.StatusOK, "Hello , API !!!")
		return handler.APIResultHandler(c, true, http.StatusOK, "Delete Board")
	}
}

// CommentBoard GET
func (api CommentBoardAPI) GetCommentBoard() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		id, _ := strconv.ParseInt(c.Param("id"), 0, 64)
		board := &rdbModel.CommentBoard{}

		api.Logger.Debug("id : ", id)

		tx := c.Get("Tx").(*gorm.DB)

		tx.Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name")
		}).Where("id = ? ", id).Preload("User.Departures").Find(&board)


		return handler.APIResultHandler(c, true, http.StatusOK, board)
	}
}

// CommentBoard GET ALL
func (api CommentBoardAPI) GetCommentBoards() echo.HandlerFunc {

	return func(c echo.Context) (err error) {

		// 검색 기능을 넣자.
		// Paging 정보
		offset, _ := strconv.Atoi(c.QueryParam("offset"))
		limit, _  := strconv.Atoi(c.QueryParam("limit"))

		title 	   := c.QueryParam("title")
		content    := c.QueryParam("content")

		api.Logger.Debug("offset : ", offset)
		api.Logger.Debug("limit ", limit)

		api.Logger.Debug("title ", title)
		api.Logger.Debug("content ", content)

		boards := []rdbModel.CommentBoard{}
		tx := c.Get("Tx").(*gorm.DB)

		// Where Condition
		if title != "" {
			tx = tx.Where("title LIKE  ?", "%" + title + "%")
		}

		if content != "" {
			tx = tx.Where("content LIKE  ?", "%" + content + "%")
		}

		// TOTAL COUNT 구하기
		var count = 0
		tx.Find(&boards).Count(&count)

		tx.Preload("User").Preload("User.Departures").Order("id desc").Offset(offset).Limit(limit).Find(&boards)

		return handler.APIResultHandler(c, true, http.StatusOK,
			map[string]interface{}{"totalcount": count, "rows": boards})
	}
}