package database

import (
	"time"
)

// Notes model
type Notes struct {
	ID       int       `json:"ID"`
	Title    string    `json:"Title"`
	Content  string    `json:"Content"`
	LastEdit time.Time `json:"LastEdit"`
}

// NotesString copy of the note model with the ID in string (for rendering purpose)
type NotesString struct {
	ID       string    `json:"ID"`
	Title    string    `json:"Title"`
	Content  string    `json:"Content"`
	LastEdit time.Time `json:"LastEdit"`
}
