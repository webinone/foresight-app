package rdb

import (
	"time"
)

type SimpleBoard struct {
	Id      	int64 		`gorm:"AUTO_INCREMENT;primary_key"`
	Created 	int64
	Title   	string		`validate:"required" json:"title" form:"title" query:"title"`
	Content 	string		`validate:"required" json:"content" form:"content" query:"content"`
	CreatedAt 	time.Time
}

type BoardComments struct {

}

// set User's table name to be `TB_SIMPLEBOARD`
func (SimpleBoard) TableName() string {
	return "TB_SIMPLEBOARD"
}