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

type PostService struct {
	postRepo *repository.PostRepository
}

func NewPostService(postRepo *repository.PostRepository) *PostService {
	return &PostService{
		postRepo: postRepo,
	}
}

func (s *PostService) CreatePost(ctx context.Context, post model.Post, file *multipart.FileHeader) (int, error) {
	if file != nil {
		if file.Size > 5*1024*1024 {
			return 0, errors.New("file size exceeds 5MB limit")
		}

		imageUrl, err := s.saveUploadedFile(file)
		if err != nil {
			return 0, err
		}
		post.ImageUrl = imageUrl
	}
	return s.postRepo.CreatePost(ctx, post)
}

func (s *PostService) saveUploadedFile(file *multipart.FileHeader) (string, error) {
	ext := filepath.Ext(file.Filename)
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" {
		return "", errors.New("invalid file type")
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	uploadDir := "./uploads/posts"
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

	return "/uploads/posts/" + filename, nil
}

func (s *PostService) GetPostById(ctx context.Context, id int) (dto.PostResponse, error) {
	return s.postRepo.GetPostById(ctx, id)
}

func (s *PostService) GetFeed(ctx context.Context, userId int) ([]dto.PostResponse, error) {
	return s.postRepo.GetFeed(ctx, userId)
}
