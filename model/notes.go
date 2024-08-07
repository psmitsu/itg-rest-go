package model

import (
	"time"
)

type Note struct {
  Id int64        `json:"id"`
  Text string   `json:"text"`
  Dt time.Time  `json:"dt"`
}

type NoteInput struct {
  Text string  `json:"text"`
  Dt time.Time `json:"dt"`
}

type INotesRepository interface {
  Create(note NoteInput) (*Note, error)
  GetOne(id int64) (*Note, error)
  GetMany() ([]Note, error)
  Update(id int64, note NoteInput) (*Note, error)
  Delete(id int64) error
}
