package repositoryimpl

import (
	"context"
	"database/sql"
	"log"

	"github.com/shou1027/cookmaBackend/pkg/domain/model"
	"github.com/shou1027/cookmaBackend/pkg/domain/repository"
	"github.com/shou1027/cookmaBackend/pkg/infrastructure/postgresql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repositoryImpl struct {
	db DBTX
}

func NewRepositoryImpl(db DBTX) repository.Repository {
	return &repositoryImpl{db: db}
}

func (ri *repositoryImpl) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	var lastInsertId int
	query := "INSERT INTO users(name, email, password, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW()) returning id"
	err := ri.db.QueryRowContext(ctx, query, user.GetName(), user.GetEmail(), user.GetPassword()).Scan(&lastInsertId)
	if err != nil {
		return nil, err
	}
	log.Println(user)

	user, err = model.Reconstruct(
		int64(lastInsertId),
		user.GetName(),
		user.GetEmail(),
		user.GetPassword(),
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ri *repositoryImpl) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	u := postgresql.User{}
	query := "SELECT id, name, email, password FROM users WHERE email = $1"
	err := ri.db.QueryRowContext(ctx, query, email).Scan(&u.Id, &u.Name, &u.Email, &u.Password)
	if err != nil {
		return nil, nil
	}

	user, err := model.Reconstruct(
		u.Id,
		u.Name,
		u.Email,
		u.Password,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
