package repository

import (
	"context"
	"time"

	"github.com/Albaihaqi354/FinalPhase3.git/internal/dto"
	"github.com/Albaihaqi354/FinalPhase3.git/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type AuthRepository struct {
	db    *pgxpool.Pool
	redis *redis.Client
}

func NewAuthRepository(db *pgxpool.Pool, rdb *redis.Client) *AuthRepository {
	return &AuthRepository{
		db:    db,
		redis: rdb,
	}
}

func (a AuthRepository) CreateNewUser(ctx context.Context, newUser dto.RegisterRequest, hashedPwd string) (model.Account, error) {
	tx, err := a.db.Begin(ctx)
	if err != nil {
		return model.Account{}, err
	}
	defer tx.Rollback(ctx)

	sqlAccount := "INSERT INTO accounts (email, password) VALUES ($1, $2) RETURNING id, email, password"
	var user model.Account
	err = tx.QueryRow(ctx, sqlAccount, newUser.Email, hashedPwd).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		return model.Account{}, err
	}

	sqlUser := "INSERT INTO users (account_id, name) VALUES ($1, $2) RETURNING id"
	err = tx.QueryRow(ctx, sqlUser, user.Id, user.Email).Scan(&user.Id)
	if err != nil {
		return model.Account{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		return model.Account{}, err
	}

	return user, nil
}

func (a AuthRepository) FindUserByEmail(ctx context.Context, email string) (model.Account, error) {
	sql := `SELECT u.id, a.email, a.password 
			FROM accounts a 
			JOIN users u ON a.id = u.account_id 
			WHERE a.email = $1`

	var user model.Account
	if err := a.db.QueryRow(ctx, sql, email).Scan(&user.Id, &user.Email, &user.Password); err != nil {
		return model.Account{}, err
	}
	return user, nil
}

func (a AuthRepository) SaveToken(ctx context.Context, token string, ttl time.Duration) error {
	rkey := "bian:socialMedia:whitelist:" + token
	return a.redis.Set(ctx, rkey, "active", ttl).Err()
}

func (a AuthRepository) DeleteToken(ctx context.Context, token string) error {
	rkey := "bian:socialMedia:whitelist:" + token
	return a.redis.Del(ctx, rkey).Err()
}

func (a AuthRepository) TokenWitelist(ctx context.Context, token string) (bool, error) {
	rkey := "bian:socialMedia:whitelist:" + token
	rsc, err := a.redis.Exists(ctx, rkey).Result()
	if err != nil {
		return false, err
	}
	return rsc > 0, nil
}
