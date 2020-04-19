package models

type Repository struct {
	ID          string   `json:"id,omitempty"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	URL         string   `json:"url"`
	Language    string   `json:"language"`
	Tags        []string `json:"tags,omitempty"`
}

type UserRepository struct {
	UserID       string
	RepositoryID string
	Tags         []string
}
