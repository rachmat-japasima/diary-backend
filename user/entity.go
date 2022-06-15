package user

import "time"

type User struct {
	ID		     int
	Name         string
	Gender       string
	Email        string
	PasswordHash string
	FileAvatar   string
	Role         string
	CreatedAt    time.Time
	UpdatedAt	 time.Time
}