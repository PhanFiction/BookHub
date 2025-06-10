package types

import "net/http"

type BookDetails struct {
	ID          int
	Title       string
	Author      string
	Pages       int
	Publisher   string
	ISBN        string
	Description string
	PublishedAt string
	Genre       string
	CoverImg    string
}

type Data struct {
	BookData      BookDetails
	Message       string
	Success       bool
	Authenticated bool
}

type User struct {
	ID       int
	Username string
	Name     string
	Email    string
	Password string
	Avatar   string
	GoogleID string
}

// Type function that takes an http.HandlerFunc and returns another http.HandlerFunc.
type Middleware func(http.HandlerFunc) http.HandlerFunc
