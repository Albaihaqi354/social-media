package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type FollowRepository struct {
	db *pgxpool.Pool
}

func NewFollowRepository(db *pgxpool.Pool) *FollowRepository {
	return &FollowRepository{
		db: db,
	}
}

func (r *FollowRepository) FollowUser(ctx context.Context, followerId, followingId int) error {
	sql := `INSERT INTO follows (follower_id, following_id, created_at) 
			VALUES ($1, $2, CURRENT_TIMESTAMP)
			ON CONFLICT (follower_id, following_id) DO NOTHING`
	_, err := r.db.Exec(ctx, sql, followerId, followingId)
	return err
}

func (r *FollowRepository) UnfollowUser(ctx context.Context, followerId, followingId int) error {
	sql := `DELETE FROM follows WHERE follower_id = $1 AND following_id = $2`
	_, err := r.db.Exec(ctx, sql, followerId, followingId)
	return err
}

func (r *FollowRepository) IsFollowing(ctx context.Context, followerId, followingId int) (bool, error) {
	sql := `SELECT EXISTS(SELECT 1 FROM follows WHERE follower_id = $1 AND following_id = $2)`
	var exists bool
	err := r.db.QueryRow(ctx, sql, followerId, followingId).Scan(&exists)
	return exists, err
}
