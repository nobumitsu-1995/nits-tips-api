package repository

import (
	"context"
	"nits-tips-api/model"

	"github.com/redis/go-redis/v9"
)

type ISessionDataRepository interface {
	SaveSession(ctx context.Context, sessionData *model.SessionData) error
	GetSession(ctx context.Context, sessionData *model.SessionData) error
}

type sessionDataRepository struct {
	db *redis.Client
}

func NewSessionDataRepository(db *redis.Client) ISessionDataRepository {
	return &sessionDataRepository{db}
}

func (sdr *sessionDataRepository) SaveSession(ctx context.Context, sessionData *model.SessionData) error {
	return sdr.db.Set(ctx, sessionData.SessionId, sessionData.UserId, 0).Err()
}

func (sdr *sessionDataRepository) GetSession(ctx context.Context, sessionData *model.SessionData) error {
	value, err := sdr.db.Get(ctx, sessionData.SessionId).Result()
	if err == redis.Nil {
		sessionData.UserId = ""
		return nil
	} else if err != nil {
		return err
	}
	sessionData.UserId = value
	return nil
}
