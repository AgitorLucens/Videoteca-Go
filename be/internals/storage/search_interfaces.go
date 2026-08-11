package storage

type SearchRepository interface {
	SearchMovies(query string) ([]MovieSerie, error)
	SearchActors(query string) ([]Actor, error)
}

type SearchResult struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Type   string `json:"type"`
	Photo  string `json:"photo,omitempty"`
	MSType string `json:"msType,omitempty"`
}
