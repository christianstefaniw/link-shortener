package models

type ShortenedURL struct {
	FullURL  string `json:”title,omitempty”`
	ShortURL string `json:”title,omitempty”`
}

type URLS struct {
	Urls []*ShortenedURL
}