package usecase

import (
	"nits-tips-api/model"
	"nits-tips-api/repository"
	"nits-tips-api/validator"
	"sort"
)

type IReactionStampUsecase interface {
	GetReactionStampsByArticleId(articleId string, userId string) (model.ReactionStampSummaryResponse, error)
	CreateReactionStamp(reactionStamp model.ReactionStamp) (model.ReactionStampResponse, error)
	DeleteReactionStamp(reactionStampId uint, userId string) error
	calcStampSummary(reactionStamps []model.ReactionStamp) []model.ReactionStampSummary
	getStampByUserId(reactionStamps []model.ReactionStamp, userId string) []int
}

type reactionStampUsecase struct {
	rsr repository.IReactionStampRepository
	rsv validator.IReactionStampValidator
}

func NewReactionStampUsecase(rsr repository.IReactionStampRepository, rsv validator.IReactionStampValidator) IReactionStampUsecase {
	return &reactionStampUsecase{rsr, rsv}
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
	if err := rsu.rsv.ReactionStampValidator(reactionStamp); err != nil {
		return model.ReactionStampResponse{}, err
	}

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

func (rsu *reactionStampUsecase) getStampByUserId(reactionStamps []model.ReactionStamp, userId string) []int {
	var reactedStamps []int

	if userId == "" {
		return []int{}
	}

	for _, stamp := range reactionStamps {
		if stamp.UserId == userId {
			reactedStamps = append(reactedStamps, stamp.StampId)
		}
	}

	if len(reactedStamps) <= 0 {
		return []int{}
	}

	sort.Ints(reactedStamps)
	return reactedStamps
}
