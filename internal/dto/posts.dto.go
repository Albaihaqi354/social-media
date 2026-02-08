package dto

import "time"

type CreatePostRequest struct {
	Content string `form:"content"`
}

type PostResponse struct {
	Id            int       `json:"id"`
	Content       string    `json:"content"`
	ImageUrl      string    `json:"image_url"`
	UserId        int       `json:"user_id"`
	UserName      string    `json:"user_name"`
	UserAvatarUrl string    `json:"user_avatar_url"`
	LikeCount     int       `json:"like_count"`
	CommentCount  int       `json:"comment_count"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
