package models

import "time"

type CreateBook struct {
	Id        string    `json:"id"`
	Title     string    `json:"title" binding:"required"`
	Author    string    `json:"author" binding:"required"`
	ISBN      string    `json:"isbn"`
	Status    string    `json:"status" binding:"required,oneof=read unread"`
	Rating    float64   `json:"rating" binding:"min=0,max=5"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateBook struct {
	Title  string  `json:"title" binding:"required"`
	Author string  `json:"author" binding:"required"`
	ISBN   string  `json:"isbn"`
	Status string  `json:"status" binding:"required, oneof=read unread"`
	Rating float64 `json:"rating" binding:"min=0,max=5"`
	Notes  string  `json:"notes"`
}

type UpdateStatus struct {
	Status string `json:"status" binding:"required, oneof=read unread"`
}

type UpdateRating struct {
	Rating float64 `json:"rating" binding:"min=0,max=5"`
}
