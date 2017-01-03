package rdb

import (
	"time"
)

type CommentBoard struct {
	Id      	int64 		`gorm:"AUTO_INCREMENT;primary_key" json:"id"`
	Title   	string		`validate:"required" json:"title" form:"title" query:"title"`
	Content 	string		`validate:"required" json:"content" form:"content" query:"content"`
	User		User		`gorm:"ForeignKey:UserId" json:"user"`
	UserId          int64		`json:"user_id" form:"user_id" query:"user_id"`
	Comments	[]Comment	`gorm:"AssociationForeignKey:CommentBoardId" json:"comments"`
	CreatedAt 	time.Time	`json:"created_at"`
}

func (CommentBoard) TableName() string {
	return "TB_COMMENTBOARD"
}

type Comment struct {
	Id      	int64 		`gorm:"AUTO_INCREMENT;primary_key"`
	Content 	string
	CommentBoardId	int64
	User		User		`gorm:"ForeignKey:UserId"`
	UserId          int64
	CreatedAt 	time.Time
}
func (Comment) TableName() string {
	return "TB_COMMENT"
}

type User struct {
	Id		int64		`gorm:"AUTO_INCREMENT;primary_key" json:"id"`
	Name		string		`json:"name"`
	Email		string		`json:"email"`
	Password	string		`json:"password"`
	Departures	[]Departure	`gorm:"many2many:TB_USER_DEPARTURES;" json:"departures"`
}

func (User) TableName() string {
	return "TB_USER"
}

type UserRole struct {
	Id              int64		`gorm:"AUTO_INCREMENT;primary_key" json:"id"`
	Name		string
}

func (UserRole) TableName() string {
	return "TB_ROLE"
}

type Departure struct {
	Id		int64		`gorm:"AUTO_INCREMENT;primary_key" json:"id"`
	Name            string		`json:"name"`
}

func (Departure) TableName() string {
	return "TB_DEPARTURE"
}