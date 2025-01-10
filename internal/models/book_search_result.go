package models

type BookSearchResult struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	ISBN          string `json:"isbn"`
	Publisher     string `json:"publisher"`
	PublishedYear int    `json:"published_year"`
	CoverURL      string `json:"cover_url"`
}
