package repository

import (
	"nits-tips-api/model"

	"gorm.io/gorm"
)

type IReactionStampRepository interface {
	GetReactionStampsByArticleId(reactionStamp *[]model.ReactionStamp, articleId string) error
	CreateReactionStamp(reactionStamp *model.ReactionStamp) error
}

type reactionStampRepository struct {
	db *gorm.DB
}

func NewReactionStampRepository(db *gorm.DB) IReactionStampRepository {
	return &reactionStampRepository{db}
}

func (rsr *reactionStampRepository) GetReactionStampsByArticleId(reactionStamps *[]model.ReactionStamp, articleId string) error {
	if err := rsr.db.Where("article_id=?", articleId).Find(reactionStamps).Error; err != nil {
		return err
	}
	return nil
}

func (rsr *reactionStampRepository) CreateReactionStamp(reactionStamp *model.ReactionStamp) error {
	if err := rsr.db.Create(reactionStamp).Error; err != nil {
		return err
	}
	return nil
}
