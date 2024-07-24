package repository

import (
	"context"

	"github.com/shou1027/cookmaBackend/pkg/domain/model"
)

type Repository interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
}
