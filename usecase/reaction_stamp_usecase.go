package usecase

import (
	"nits-tips-api/model"
	"nits-tips-api/repository"
)

type IReactionStampUsecase interface {
	GetReactionStampsByArticleId(articleId string) ([]model.ReactionStampResponse, error)
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
