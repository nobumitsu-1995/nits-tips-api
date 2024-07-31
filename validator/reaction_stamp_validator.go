package validator

import (
	"nits-tips-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IReactionStampValidator interface {
	ReactionStampValidator(reactionStamp model.ReactionStamp) error
}

type reactionStampValidator struct{}

func NewReactionStampValidator() IReactionStampValidator {
	return &reactionStampValidator{}
}

func (rsv *reactionStampValidator) ReactionStampValidator(reactionStamp model.ReactionStamp) error {
	return validation.ValidateStruct(&reactionStamp,
		validation.Field(
			&reactionStamp.UserId,
			validation.Required.Error("UserId is required"),
		),
		validation.Field(
			&reactionStamp.StampId,
			validation.Required.Error("StampId is required"),
			validation.Min(1).Error("StampId must be at least 1"),
			validation.Max(6).Error("StampId must be at most 6"),
		),
		validation.Field(
			&reactionStamp.ArticleId,
			validation.Required.Error("ArticleId is required"),
		),
	)
}
