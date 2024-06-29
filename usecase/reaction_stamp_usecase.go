package usecase

import (
	"nits-tips-api/model"
	"nits-tips-api/repository"
)

type IReactionStampUsecase interface {
	GetReactionStampsByArticleId(articleId string, userId string) (model.ReactionStampSummaryResponse, error)
	CreateReactionStamp(reactionStamp model.ReactionStamp) (model.ReactionStampResponse, error)
	DeleteReactionStamp(reactionStampId uint, userId string) error
	calcStampSummary(reactionStamps []model.ReactionStamp) []model.ReactionStampSummary
	getStampByUserId(reactionStamps []model.ReactionStamp, userId string) []model.ReactedStamp
}

type reactionStampUsecase struct {
	rsr repository.IReactionStampRepository
}

func NewReactionStampUsecase(rsr repository.IReactionStampRepository) IReactionStampUsecase {
	return &reactionStampUsecase{rsr}
}

func (rsu *reactionStampUsecase) GetReactionStampsByArticleId(articleId string, userId string) (model.ReactionStampSummaryResponse, error) {
	reactionStamps := []model.ReactionStamp{}

	if err := rsu.rsr.GetReactionStampsByArticleId(&reactionStamps, articleId); err != nil {
		return model.ReactionStampSummaryResponse{}, err
	}

	reactionStampsResponse := model.ReactionStampSummaryResponse{
		ReactionStampSummary: rsu.calcStampSummary(reactionStamps),
		ReactedStamp:         rsu.getStampByUserId(reactionStamps, userId),
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

func (rsu *reactionStampUsecase) DeleteReactionStamp(reactionStampId uint, userId string) error {
	if err := rsu.rsr.DeleteReactionStamp(reactionStampId, userId); err != nil {
		return err
	}
	return nil
}

func (rsu *reactionStampUsecase) calcStampSummary(reactionStamps []model.ReactionStamp) []model.ReactionStampSummary {
	counts := make(map[int]int)

	for _, stamp := range reactionStamps {
		if stamp.StampId >= 1 && stamp.StampId <= 6 {
			counts[stamp.StampId]++
		}
	}

	var results []model.ReactionStampSummary

	for i := 1; i <= 6; i++ {
		results = append(results, model.ReactionStampSummary{
			StampId: i,
			Total:   counts[i],
		})
	}

	return results
}

func (rsu *reactionStampUsecase) getStampByUserId(reactionStamps []model.ReactionStamp, userId string) []model.ReactedStamp {
	var reactedStamps []model.ReactedStamp

	for _, stamp := range reactionStamps {
		if stamp.UserId == userId {
			reactedStamps = append(reactedStamps, model.ReactedStamp{
				ID:      stamp.ID,
				StampId: stamp.StampId,
			})
		}
	}

	return reactedStamps
}
