package cueparser

type CueSheet struct {
	Title     string
	Performer string
	Files     []File
}

type File struct {
	Name   string
	Type   string // usually "WAVE"
	Tracks []Track
}

type Track struct {
	Number    int
	Title     string
	Performer string
	Indexes   []Index
}

type Index struct {
	Number int
	Time   string // MM:SS:FF
}
