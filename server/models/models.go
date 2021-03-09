package models

type ShortenedURL struct {
	FullURL  string
	ShortURL string
}

type URLS struct {
	Urls []*ShortenedURL
}