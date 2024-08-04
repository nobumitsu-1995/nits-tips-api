package usecase

import (
	"context"
	"nits-tips-api/model"
	"nits-tips-api/repository"
	"nits-tips-api/validator"

	"github.com/google/uuid"
)

type ISessionDataUsecase interface {
	GetSession(ctx context.Context, sessionId string) (model.SessionData, error)
	createNewSession(ctx context.Context, sessionData *model.SessionData) error
	GenerateUUID() (string, error)
}

type sessionDataUsecase struct {
	sdr repository.ISessionDataRepository
	sv  validator.ISessionValidator
}

func NewSessionDataUsecase(sdr repository.ISessionDataRepository, sv validator.ISessionValidator) ISessionDataUsecase {
	return &sessionDataUsecase{sdr, sv}
}

func (sdu *sessionDataUsecase) GenerateUUID() (string, error) {
	uuid, err := uuid.NewV7()
	if err != nil {
		return "", nil
	}
	return uuid.String(), nil
}

func (sdu *sessionDataUsecase) createNewSession(ctx context.Context, sessionData *model.SessionData) error {
	sessionId, err := sdu.GenerateUUID()
	if err != nil {
		return err
	}
	userId, err := sdu.GenerateUUID()
	if err != nil {
		return err
	}

	sessionData.SessionId = sessionId
	sessionData.UserId = userId

	if err := sdu.sv.SessionValidator(*sessionData); err != nil {
		return err
	}

	if err = sdu.sdr.SaveSession(ctx, sessionData); err != nil {
		return err
	}
	return nil
}

func (sdu *sessionDataUsecase) GetSession(ctx context.Context, sessionId string) (model.SessionData, error) {
	sessionData := model.SessionData{
		SessionId: sessionId,
	}
	if err := sdu.sdr.GetSession(ctx, &sessionData); err != nil {
		return model.SessionData{}, err
	}

	if sessionData.UserId == "" {
		if err := sdu.createNewSession(ctx, &sessionData); err != nil {
			return model.SessionData{}, err
		}
		return sessionData, nil
	}

	return sessionData, nil
}
