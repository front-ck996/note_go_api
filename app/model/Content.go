package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID uint `gorm:"primarykey" json:"id"`
	//gorm.Model
	//ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Note struct {
	BaseModel
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	TopicID uint   `json:"topic_id"`
	Sort    int    `json:"sort"`
}

type TopicJson struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Parent uint   `json:"parent"`
	Sort   int    `json:"sort"`
}
type Topic struct {
	BaseModel
	TopicJson
}
