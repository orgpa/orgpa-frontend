package database

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Notes model
type Notes struct {
	ID       bson.ObjectId `json:"ID"`
	Title    string        `json:"Title"`
	Content  string        `json:"Content"`
	LastEdit time.Time     `json:"LastEdit"`
}

// NotesString copy of the note model with the ID in string (for rendering purpose)
type NotesString struct {
	ID       string    `json:"ID"`
	Title    string    `json:"Title"`
	Content  string    `json:"Content"`
	LastEdit time.Time `json:"LastEdit"`
}
