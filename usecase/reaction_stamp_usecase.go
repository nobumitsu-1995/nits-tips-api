package usecase

import (
	"nits-tips-api/model"
	"nits-tips-api/repository"
)

type IReactionStampUsecase interface {
	GetReactionStampsByArticleId(articleId string) ([]model.ReactionStampResponse, error)
	CreateReactionStamp(reactionStamp model.ReactionStamp) (model.ReactionStampResponse, error)
	DeleteReactionStamp(reactionStampId uint, userId uint) error
}

type reactionStampUsecase struct {
	rsr repository.IReactionStampRepository
}

func NewReactionStampUsecase(rsr repository.IReactionStampRepository) IReactionStampUsecase {
	return &reactionStampUsecase{rsr}
}

func (rsu *reactionStampUsecase) GetReactionStampsByArticleId(articleId string) ([]model.ReactionStampResponse, error) {
	reactionStamps := []model.ReactionStamp{}

	if err := rsu.rsr.GetReactionStampsByArticleId(&reactionStamps, articleId); err != nil {
		return nil, err
	}

	reactionStampsResponse := []model.ReactionStampResponse{}
	for _, v := range reactionStamps {
		rs := model.ReactionStampResponse{
			ID:      v.ID,
			StampId: v.StampId,
		}
		reactionStampsResponse = append(reactionStampsResponse, rs)
	}

	return reactionStampsResponse, nil
}

func (rsu *reactionStampUsecase) CreateReactionStamp(reactionStamp model.ReactionStamp) (model.ReactionStampResponse, error) {
	if err := rsu.rsr.CreateReactionStamp(&reactionStamp); err != nil {
		return model.ReactionStampResponse{}, err
	}

	reactionStampResponse := model.ReactionStampResponse{
		ID:      reactionStamp.ID,
		StampId: reactionStamp.StampId,
	}
	return reactionStampResponse, nil
}

func (rsu *reactionStampUsecase) DeleteReactionStamp(reactionStampId uint, userId uint) error {
	if err := rsu.rsr.DeleteReactionStamp(reactionStampId, userId); err != nil {
		return err
	}
	return nil
}
