package repository

import (
	"fmt"
	"nits-tips-api/model"

	"gorm.io/gorm"
)

type IReactionStampRepository interface {
	GetReactionStampsByArticleId(reactionStamp *[]model.ReactionStamp, articleId string) error
	CreateReactionStamp(reactionStamp *model.ReactionStamp) error
	DeleteReactionStamp(reactionStampId uint, userId string) error
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

func (rsr *reactionStampRepository) DeleteReactionStamp(reactionStampId uint, userId string) error {
	result := rsr.db.Where("stamp_id=? AND user_id=?", reactionStampId, userId).Delete(&model.ReactionStamp{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("Objectが存在しません")
	}

	return nil
}
