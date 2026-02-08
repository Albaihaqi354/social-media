package service

import (
	"context"

	"github.com/Albaihaqi354/FinalPhase3.git/internal/repository"
)

type FollowService struct {
	followRepo *repository.FollowRepository
}

func NewFollowService(followRepo *repository.FollowRepository) *FollowService {
	return &FollowService{
		followRepo: followRepo,
	}
}

func (s *FollowService) FollowUser(ctx context.Context, followerId, followingId int) error {
	return s.followRepo.FollowUser(ctx, followerId, followingId)
}

func (s *FollowService) UnfollowUser(ctx context.Context, followerId, followingId int) error {
	return s.followRepo.UnfollowUser(ctx, followerId, followingId)
}
