package repository

import (
	"context"

	"github.com/Albaihaqi354/FinalPhase3.git/internal/dto"
	"github.com/jackc/pgx/v5/pgxpool"
)

type InteractionRepository struct {
	db *pgxpool.Pool
}

func NewInteractionRepository(db *pgxpool.Pool) *InteractionRepository {
	return &InteractionRepository{
		db: db,
	}
}

func (r *InteractionRepository) LikePost(ctx context.Context, postId, userId int) error {
	sql := `INSERT INTO likes (post_id, user_id, created_at) 
			VALUES ($1, $2, CURRENT_TIMESTAMP)
			ON CONFLICT (post_id, user_id) DO NOTHING`
	_, err := r.db.Exec(ctx, sql, postId, userId)
	return err
}

func (r *InteractionRepository) UnlikePost(ctx context.Context, postId, userId int) error {
	sql := `DELETE FROM likes WHERE post_id = $1 AND user_id = $2`
	_, err := r.db.Exec(ctx, sql, postId, userId)
	return err
}

func (r *InteractionRepository) CreateComment(ctx context.Context, postId, userId int, content string) (int, error) {
	sql := `INSERT INTO comments (post_id, user_id, content, created_at, updated_at) 
			VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, sql, postId, userId, content).Scan(&id)
	return id, err
}

func (r *InteractionRepository) GetCommentsByPostId(ctx context.Context, postId int) ([]dto.CommentResponse, error) {
	sql := `SELECT c.id, c.post_id, c.user_id, u.name, COALESCE(u.avatar_url, ''), c.content, c.created_at, c.updated_at
			FROM comments c
			JOIN users u ON c.user_id = u.id
			WHERE c.post_id = $1
			ORDER BY c.created_at ASC`

	rows, err := r.db.Query(ctx, sql, postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []dto.CommentResponse
	for rows.Next() {
		var comment dto.CommentResponse
		err := rows.Scan(
			&comment.Id, &comment.PostId, &comment.UserId, &comment.UserName, &comment.UserAvatarUrl,
			&comment.Content, &comment.CreatedAt, &comment.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
