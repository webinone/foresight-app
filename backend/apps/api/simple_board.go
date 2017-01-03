package api

import (
	"github.com/labstack/echo"
	"net/http"
	"foresight-app.v1/backend/apps/models/rdb"
	"foresight-app.v1/backend/apps/handler"
	"github.com/jinzhu/gorm"
	"github.com/Sirupsen/logrus"
	"foresight-app.v1/backend/apps/libs"
	"strconv"
	"gopkg.in/go-playground/validator.v9"
	"fmt"
)

type SimpleBoardAPI struct {
	Logger 		*logrus.Logger
	Validator 	*validator.Validate
}

var SimpleBoardAPIObj = &SimpleBoardAPI{
	Logger 		: libs.GetLogger("api-simpleboard"),
	Validator	: validator.New(),
}

// Board Insert
func (api SimpleBoardAPI) CreateBoard() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		board := &rdb.SimpleBoard{}

		c.Bind(board)

		fmt.Scanln(board)

		fmt.Println("board : ", board)

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
		//return c.JSON(http.StatusOK, "Hello , API !!!")
		return handler.APIResultHandler(c, true, http.StatusOK, "Created Board")
	}
}

// Board PUT (Update)
func (api SimpleBoardAPI) PutBoard() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		board := &rdb.SimpleBoard{}
		c.Bind(board)

		id, _ := strconv.ParseInt(c.Param("id"), 0, 64)

		tx := c.Get("Tx").(*gorm.DB)
		tx.Model(board).Where("id = ?", id).Updates(
			map[string]interface{}{
				"title": board.Title,
				"content": board.Content,
			})

		return handler.APIResultHandler(c, true, http.StatusOK, "Update Board")
	}
}


// Board DELETE (Delete)
func (api SimpleBoardAPI) DeleteBoard() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		id, _ := strconv.ParseInt(c.Param("id"), 0, 64)
		board := &rdb.SimpleBoard{}

		tx := c.Get("Tx").(*gorm.DB)
		tx.Delete(board, "id = ? ", id)

		// ERROR 발생
		//return echo.NewHTTPError(http.StatusBadRequest, "Must provide both email and password")
		//return c.JSON(http.StatusOK, "Hello , API !!!")
		return handler.APIResultHandler(c, true, http.StatusOK, "Delete Board")
	}
}

// Board GET
func (api SimpleBoardAPI) GetBoard() echo.HandlerFunc {

	return func(c echo.Context) (err error) {
		id, _ := strconv.ParseInt(c.Param("id"), 0, 64)

		api.Logger.Debug("id : ", id)

		board := rdb.SimpleBoard{}

		tx := c.Get("Tx").(*gorm.DB)
		tx.Where("id = ? ", id).Find(&board)

		return handler.APIResultHandler(c, true, http.StatusOK, board)
	}
}

// Boards GET
func (api SimpleBoardAPI) GetBoards() echo.HandlerFunc {

	return func(c echo.Context) (err error) {

		// Paging 정보
		offset, _ := strconv.Atoi(c.QueryParam("offset"))
		limit, _  := strconv.Atoi(c.QueryParam("limit"))

		api.Logger.Debug("offset : ", offset)
		api.Logger.Debug("limit ", limit)

		boards := []rdb.SimpleBoard{}
		tx := c.Get("Tx").(*gorm.DB)

		// TOTAL COUNT 구하기
		var count = 0
		tx.Find(&boards).Count(&count)

		tx.Order("id desc").Offset(offset).Limit(limit).Find(&boards)

		return handler.APIResultHandler(c, true, http.StatusOK,
			map[string]interface{}{"totalcount": count, "rows": boards})
	}
}
