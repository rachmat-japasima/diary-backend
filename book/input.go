package book

import "diary/user"

type GetBookDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateBookInput struct {
	Title      string `json:"title" binding:"required"`
	Writer     string `json:"writer" binding:"required"`
	Pages      int    `json:"pages" binding:"required"`
	Synopsis   string `json:"synopsis" binding:"required"`
	User       user.User
	Category   []string `json:"category" binding:"required"`
}

type GetBookStatusInput struct {
	ID int `uri:"id" binding:"required"`
	Status string `uri:"status" binding:"required"`
}