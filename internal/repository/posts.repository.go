package repository

import (
	"context"

	"github.com/Albaihaqi354/FinalPhase3.git/internal/dto"
	"github.com/Albaihaqi354/FinalPhase3.git/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostRepository struct {
	db *pgxpool.Pool
}

func NewPostRepository(db *pgxpool.Pool) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (r *PostRepository) CreatePost(ctx context.Context, post model.Post) (int, error) {
	sql := `INSERT INTO posts (user_id, content, image_url, created_at, updated_at) 
			VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, sql, post.UserId, post.Content, post.ImageUrl).Scan(&id)
	return id, err
}

func (r *PostRepository) GetPostById(ctx context.Context, id int) (dto.PostResponse, error) {
	sql := `SELECT p.id, p.content, COALESCE(p.image_url, ''), p.user_id, u.name, COALESCE(u.avatar_url, ''), 
			(SELECT COUNT(*) FROM likes WHERE post_id = p.id) as like_count,
			(SELECT COUNT(*) FROM comments WHERE post_id = p.id) as comment_count,
			p.created_at, p.updated_at
			FROM posts p
			JOIN users u ON p.user_id = u.id
			WHERE p.id = $1`

	var post dto.PostResponse
	err := r.db.QueryRow(ctx, sql, id).Scan(
		&post.Id, &post.Content, &post.ImageUrl, &post.UserId, &post.UserName, &post.UserAvatarUrl,
		&post.LikeCount, &post.CommentCount, &post.CreatedAt, &post.UpdatedAt,
	)
	return post, err
}

func (r *PostRepository) GetFeed(ctx context.Context, userId int) ([]dto.PostResponse, error) {
	sql := `SELECT p.id, p.content, COALESCE(p.image_url, ''), p.user_id, u.name, COALESCE(u.avatar_url, ''), 
			(SELECT COUNT(*) FROM likes WHERE post_id = p.id) as like_count,
			(SELECT COUNT(*) FROM comments WHERE post_id = p.id) as comment_count,
			p.created_at, p.updated_at
			FROM posts p
			JOIN users u ON p.user_id = u.id
			WHERE p.user_id IN (SELECT following_id FROM follows WHERE follower_id = $1)
			OR p.user_id = $1
			ORDER BY p.created_at DESC`

	rows, err := r.db.Query(ctx, sql, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []dto.PostResponse
	for rows.Next() {
		var post dto.PostResponse
		err := rows.Scan(
			&post.Id, &post.Content, &post.ImageUrl, &post.UserId, &post.UserName, &post.UserAvatarUrl,
			&post.LikeCount, &post.CommentCount, &post.CreatedAt, &post.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
