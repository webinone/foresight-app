package rdb

type Launguage struct {
	ID uint `gorm:"primary_key"`
	Name string
}

func (Launguage) TableName() string {
	return "TB_LANGUAGE"
}

type Movie struct {
	ID uint `gorm:"primary_key"`
	Title string
	LaunguageID uint
	Launguage Launguage
}

func (Movie) TableName() string {
	return "TB_MOVIE"
}

type Artist struct {
	ID uint `gorm:"primary_key"`
	Name string
	Movies []Movie `gorm:"many2many:TB_ARTIST_MOVIES"`
}

func (Artist) TableName() string {
	return "TB_ARTIST"
}