package service

import (
	"context"

	"github.com/Albaihaqi354/FinalPhase3.git/internal/dto"
	"github.com/Albaihaqi354/FinalPhase3.git/internal/repository"
)

type InteractionService struct {
	interactionRepo *repository.InteractionRepository
}

func NewInteractionService(interactionRepo *repository.InteractionRepository) *InteractionService {
	return &InteractionService{
		interactionRepo: interactionRepo,
	}
}

func (s *InteractionService) LikePost(ctx context.Context, postId, userId int) error {
	return s.interactionRepo.LikePost(ctx, postId, userId)
}

func (s *InteractionService) UnlikePost(ctx context.Context, postId, userId int) error {
	return s.interactionRepo.UnlikePost(ctx, postId, userId)
}

func (s *InteractionService) CreateComment(ctx context.Context, postId, userId int, content string) (int, error) {
	return s.interactionRepo.CreateComment(ctx, postId, userId, content)
}

func (s *InteractionService) GetCommentsByPostId(ctx context.Context, postId int) ([]dto.CommentResponse, error) {
	return s.interactionRepo.GetCommentsByPostId(ctx, postId)
}
