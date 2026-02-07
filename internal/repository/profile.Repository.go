package repository

import (
	"context"

	"github.com/Albaihaqi354/FinalPhase3.git/internal/dto"
	"github.com/Albaihaqi354/FinalPhase3.git/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProfileRepository struct {
	db *pgxpool.Pool
}

func NewProfileRepository(db *pgxpool.Pool) *ProfileRepository {
	return &ProfileRepository{
		db: db,
	}
}

func (r *ProfileRepository) GetProfileByAccountId(ctx context.Context, accountId int) (dto.ProfileResponse, error) {
	sql := `SELECT COALESCE(u.id, 0), a.email, COALESCE(u.name, ''), COALESCE(u.avatar_url, ''), COALESCE(u.bio, '') 
			FROM accounts a 
			LEFT JOIN users u ON a.id = u.account_id 
			WHERE a.id = $1`

	var profile dto.ProfileResponse
	err := r.db.QueryRow(ctx, sql, accountId).Scan(
		&profile.Id,
		&profile.Email,
		&profile.Name,
		&profile.AvatarUrl,
		&profile.Bio,
	)

	if err != nil {
		return dto.ProfileResponse{}, err
	}
	return profile, nil
}

func (r *ProfileRepository) UpdateProfile(ctx context.Context, accountId int, profile model.Users) error {
	sql := `UPDATE users SET name = $1, avatar_url = $2, bio = $3, updated_at = CURRENT_TIMESTAMP 
			WHERE account_id = $4`

	_, err := r.db.Exec(ctx, sql, profile.Name, profile.AvatarUrl, profile.Bio, accountId)
	return err
}
