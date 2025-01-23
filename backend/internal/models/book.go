package models

import "time"

type Metadata struct {
	Result      int `json:"results"`
	CurrentPage int `json:"current_page"`
	PageLimit   int `json:"page_limit"`
}

type PaginatedResponse struct {
	Status   int         `json:"status"`
	Metadata Metadata    `json:"metadata"`
	Data     interface{} `json:"data"`
}

type Book struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	ISBN      string    `json:"isbn"`
	Status    string    `json:"status"`
	Rating    float64   `json:"rating" binding:"min=0.0,max=5.0"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateBook struct {
	Id        string    `json:"id"`
	Title     string    `json:"title" binding:"required"`
	Author    string    `json:"author" binding:"required"`
	ISBN      string    `json:"isbn"`
	Status    string    `json:"status" binding:"required,oneof=read unread"`
	Rating    float64   `json:"rating" binding:"min=0.0,max=5.0"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateStatus struct {
	Status    string    `json:"status" binding:"required,oneof=read unread"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateRating struct {
	Rating    float64   `json:"rating" binding:"min=0,max=5"`
	UpdatedAt time.Time `json:"updated_at"`
}
