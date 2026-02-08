package dto

import "time"

type CommentRequest struct {
	Content string `json:"content" binding:"required"`
}

type CommentResponse struct {
	Id            int       `json:"id"`
	PostId        int       `json:"post_id"`
	UserId        int       `json:"user_id"`
	UserName      string    `json:"user_name"`
	UserAvatarUrl string    `json:"user_avatar_url"`
	Content       string    `json:"content"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
