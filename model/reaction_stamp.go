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

type ReactionStampSummary struct {
	StampId int
	Total   int
}

type ReactionStampSummaryResponse struct {
	ReactionStampSummary []ReactionStampSummary
	ReactedStamp         []int
}

type ReactionStampResponse struct {
	ID      uint
	StampId int
}
