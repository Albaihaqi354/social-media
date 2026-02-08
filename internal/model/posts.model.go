package model

import "time"

type Post struct {
	Id        int       `db:"id"`
	UserId    int       `db:"user_id"`
	Content   string    `db:"content"`
	ImageUrl  string    `db:"image_url"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Follow struct {
	Id          int       `db:"id"`
	FollowerId  int       `db:"follower_id"`
	FollowingId int       `db:"following_id"`
	CreatedAt   time.Time `db:"created_at"`
}

type Like struct {
	Id        int       `db:"id"`
	PostId    int       `db:"post_id"`
	UserId    int       `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
}

type Comment struct {
	Id        int       `db:"id"`
	PostId    int       `db:"post_id"`
	UserId    int       `db:"user_id"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
