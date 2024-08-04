package validator

import (
	"nits-tips-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ISessionValidator interface {
	SessionValidator(sessionData model.SessionData) error
}

type sessionValidator struct{}

func NewSessionValidator() ISessionValidator {
	return &sessionValidator{}
}

func (sv *sessionValidator) SessionValidator(sessionData model.SessionData) error {
	return validation.ValidateStruct(&sessionData,
		validation.Field(
			&sessionData.SessionId,
			validation.Required.Error("SessionId is required"),
		),
		validation.Field(
			&sessionData.UserId,
			validation.Required.Error("UserId is required"),
		),
	)
}
