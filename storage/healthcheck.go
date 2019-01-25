package storage

import (
	"context"
	"users.git/models"
)

func (s Storage) Ping(ctx context.Context) (*models.Ping, error) {
	// TODO: Conduct database transaction
	successfulCheck := &models.Ping{Pong : "pong"}
	return successfulCheck, nil
}