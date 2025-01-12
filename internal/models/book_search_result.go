package models

type metadata struct {
	Total int `json:"total"`
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type paginatedResponse struct {
	Data     interface{} `json:"data"`
	Metadata metadata    `json:"metadata"`
}

type OpenLibraryBook struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	ISBN          string `json:"isbn"`
	Publisher     string `json:"publisher"`
	PublishedYear int    `json:"published_year"`
	CoverURL      string `json:"cover_url"`
	Description   string `json:"description"`
	PageCount     string `json:"page_count"`
}
