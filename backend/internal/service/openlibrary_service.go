package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Bantamlak12/personal_book_library_manager/internal/models"
)

type OpenLibraryBook struct {
	Title  string   `json:"title"`
	Author []string `json:"author_name"`
	ISBN   []string `json:"isbn"`
}

type OpenLibrarySearchResponse struct {
	Metadata models.Metadata   `json:"metadata"`
	NumFound int64             `json:"numFound"`
	Docs     []OpenLibraryBook `json:"docs"`
}

func constructURL(page, limit int, title, author, isbn string) (string, error) {
	baseURL := "https://openlibrary.org/search.json"

	// Handle invalid page and page limit vlaues
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 20 {
		limit = 12
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Create URL
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	// Add query parameters
	query := u.Query()
	query.Set("limit", fmt.Sprintf("%d", limit))
	query.Set("offset", fmt.Sprintf("%d", offset))

	if title != "" {
		query.Set("title", title)
	}
	if author != "" {
		query.Set("author", author)
	}
	if isbn != "" {
		query.Set("isbn", isbn)
	}

	// Encode the query parameters to the URL
	u.RawQuery = query.Encode()

	return u.String(), nil
}

func (s *bookService) SearchFromOpenLibrary(page, limit int, title, author, isbn string) (*OpenLibrarySearchResponse, error) {
	u, err := constructURL(page, limit, title, author, isbn)
	if err != nil {
		return nil, fmt.Errorf("error constructing URL: %w", err)
	}

	// Fetch data from the API
	response, err := http.Get(u)
	if err != nil {
		return nil, fmt.Errorf("error fetching data from OpenLibrary: %w", err)
	}
	defer response.Body.Close()

	// Check the status
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non OK HTTP status: %s", response.Status)
	}

	// Decode JSON response
	var rawResponse OpenLibrarySearchResponse
	err = json.NewDecoder(response.Body).Decode(&rawResponse)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %w", err)
	}
	// Calculat the total page
	totalPage := (rawResponse.NumFound + int64(limit) - 1) / int64(limit)

	// Map the books to the model
	return &OpenLibrarySearchResponse{
		Metadata: models.Metadata{
			Result:      int(totalPage),
			CurrentPage: page,
			PageLimit:   limit,
		},
		NumFound: rawResponse.NumFound,
		Docs:     rawResponse.Docs,
	}, nil
}
