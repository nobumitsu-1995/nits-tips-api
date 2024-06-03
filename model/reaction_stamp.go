package model

import "time"

type ReactionStamp struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserId    string    `json:"user_id" gorm:"not null"`
	StampId   int       `json:"stamp_id" gorm:"not null"`
	ArticleId string    `json:"article_id" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ReactionStampResponse struct {
	ID      uint `json:"id" gorm:"primaryKey"`
	StampId int  `json:"stamp_id" gorm:"not null"`
}
