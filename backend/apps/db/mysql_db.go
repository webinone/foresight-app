package db

import (
	. "foresight-app.v1/backend/apps/libs"
	rdbModel "foresight-app.v1/backend/apps/models/rdb"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type MySqlDB struct {
	DB *gorm.DB
}

var MySqlDBObj = &MySqlDB{}

func (mysqlDb MySqlDB) Init() *gorm.DB {

	// Connection String
	rdb, err := gorm.Open(Config.RDB.Product, Config.RDB.ConnectString)
	if err != nil {
		panic(err)
	}

	MySqlDBObj.DB = rdb
	rdb.DB().Ping()
	rdb.DB().SetMaxIdleConns(10)
	rdb.DB().SetMaxOpenConns(100)

	rdb.LogMode(Config.RDB.Debug)

	// TODO : 생성할 테이블 목록 정의 (앞으로는 이것을 통해서 하자.)
	//------------------------------------------------------------
	//mysqlDb.DB.AutoMigrate(&rdbmodel.Board{})
	rdb.AutoMigrate(&rdbModel.SimpleBoard{})

	//if rdb.HasTable(&rdbmodel.Comment{}) {
	//	rdb.DropTable(&rdbmodel.Comment{})
	//	rdb.DropTable(&rdbmodel.User{})
	//	rdb.DropTable(&rdbmodel.Departure{})
	//	rdb.DropTable(&rdbmodel.CommentBoard{})
	//}
	//
	//// Comment Board
	//rdb.AutoMigrate(&rdbmodel.User{})
	//rdb.AutoMigrate(&rdbmodel.Comment{})
	//rdb.AutoMigrate(&rdbmodel.Departure{})
	//rdb.AutoMigrate(&rdbmodel.CommentBoard{})

	rdb.AutoMigrate(&rdbModel.Launguage{})
	rdb.AutoMigrate(&rdbModel.Movie{})
	rdb.AutoMigrate(&rdbModel.Artist{})

	return rdb
}