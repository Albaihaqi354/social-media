package model

import "time"

type Users struct {
	Id        int       `db:"id"`
	AccountId int       `db:"account_id"`
	Name      string    `db:"name"`
	AvatarUrl string    `db:"avatar_url"`
	Bio       string    `db:"bio"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"update_at"`
}

type Account struct {
	Id        int       `db:"id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"update_at"`
}
