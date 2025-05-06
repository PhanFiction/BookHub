package types

import "net/http"

type BookDetails struct {
	Title       string
	Author      string
	Pages       int
	Publisher   string
	ISBN        string
	Description string
	PublishedAt string
}

type Data struct {
	BookData      BookDetails
	Message       string
	Success       bool
	Authenticated bool
}

// Type function that takes an http.HandlerFunc and returns another http.HandlerFunc.
type Middleware func(http.HandlerFunc) http.HandlerFunc
