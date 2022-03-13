package models

import (
	"time"
)

type Article struct {
	ID        uint       `json:"id" xml:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt" xml:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt" xml:"updatedAt"`
	DeletedAt *time.Time `json:"-" xml:"-" sql:"index"`
	Title     string     `json:"title" xml:"title" gorm:"type:varchar(255)"`
	Body      string     `json:"body" xml:"body" gorm:"type:text"`
}
