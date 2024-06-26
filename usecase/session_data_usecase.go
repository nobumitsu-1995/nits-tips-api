package usecase

import (
	"context"
	"nits-tips-api/model"
	"nits-tips-api/repository"

	"github.com/google/uuid"
)

type ISessionDataUsecase interface {
	GetSession(ctx context.Context, sessionId string) (model.SessionData, error)
	createNewSession(ctx context.Context, sessionData *model.SessionData) error
	generateUUID() (string, error)
}

type sessionDataUsecase struct {
	sdr repository.ISessionDataRepository
}

func NewSessionDataUsecase(sdr repository.ISessionDataRepository) ISessionDataUsecase {
	return &sessionDataUsecase{sdr}
}

func (sdu *sessionDataUsecase) generateUUID() (string, error) {
	uuid, err := uuid.NewV7()
	if err != nil {
		return "", nil
	}
	return uuid.String(), nil
}

func (sdu *sessionDataUsecase) createNewSession(ctx context.Context, sessionData *model.SessionData) error {
	sessionId, err := sdu.generateUUID()
	if err != nil {
		return err
	}
	userId, err := sdu.generateUUID()
	if err != nil {
		return err
	}

	sessionData.SessionId = sessionId
	sessionData.UserId = userId

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
