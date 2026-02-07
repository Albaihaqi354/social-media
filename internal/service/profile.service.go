package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/Albaihaqi354/FinalPhase3.git/internal/dto"
	"github.com/Albaihaqi354/FinalPhase3.git/internal/model"
	"github.com/Albaihaqi354/FinalPhase3.git/internal/repository"
)

type ProfileService struct {
	repo *repository.ProfileRepository
}

func NewProfileService(repo *repository.ProfileRepository) *ProfileService {
	return &ProfileService{
		repo: repo,
	}
}

func (s *ProfileService) GetProfile(ctx context.Context, accountId int) (dto.ProfileResponse, error) {
	return s.repo.GetProfileByAccountId(ctx, accountId)
}

func (s *ProfileService) UpdateProfile(ctx context.Context, accountId int, req dto.UpdateProfileRequest, file *multipart.FileHeader) (dto.ProfileResponse, error) {
	profile := model.Users{
		Name: req.Name,
		Bio:  req.Bio,
	}

	if file != nil {
		if file.Size > 2*1024*1024 {
			return dto.ProfileResponse{}, errors.New("file size exceeds 2MB limit")
		}

		avatarUrl, err := s.saveUploadedFile(file)
		if err != nil {
			return dto.ProfileResponse{}, err
		}
		profile.AvatarUrl = avatarUrl
	} else {
		existing, err := s.repo.GetProfileByAccountId(ctx, accountId)
		if err == nil {
			profile.AvatarUrl = existing.AvatarUrl
		}
	}

	err := s.repo.UpdateProfile(ctx, accountId, profile)
	if err != nil {
		return dto.ProfileResponse{}, err
	}

	return s.repo.GetProfileByAccountId(ctx, accountId)
}

func (s *ProfileService) saveUploadedFile(file *multipart.FileHeader) (string, error) {
	ext := filepath.Ext(file.Filename)
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		return "", errors.New("invalid file type")
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	uploadDir := "./uploads/avatars"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return "", err
	}

	filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
	dstPath := filepath.Join(uploadDir, filename)

	dst, err := os.Create(dstPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	return "/uploads/avatars/" + filename, nil
}
