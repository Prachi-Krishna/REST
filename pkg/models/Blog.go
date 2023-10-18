package models

import (
	"time"
)

type Blog struct {
	Id        int `gorm:"primary_key;autoIncrement"`
	Category  string
	Title     string
	Content   string `gorm:"type:text"`
	CreatedAt time.Time
}

type Config struct {
	Service string
	Port    string
	Db      string
}
