package model

import (
	"cmp"
	"fmt"
	"slices"
	// "time"
)

type NotesRepoMem struct {
  storage []Note
}

func MakeNotesRepoMem() *NotesRepoMem {
  return &NotesRepoMem{storage: make([]Note, 0) }
}

func (r *NotesRepoMem) Create(noteInput NoteInput) (*Note, error) {
  // if (noteInput.Text == nil || noteInput.Dt == nil) {
  //   return nil, fmt.Errorf("wrong input")
  // }

  var maxId int64

  if (len(r.storage) == 0) {
    maxId = 0;
  } else {
    maxId = slices.MaxFunc(r.storage, func(a, b Note) int {
      return cmp.Compare(a.Id, b.Id)
    }).Id
  }

  note := Note{
    Id: maxId + 1,
    Text: noteInput.Text,
    Dt: noteInput.Dt,
  }

  r.storage = append(r.storage, note) 
  return &note, nil
}

func (r *NotesRepoMem) GetOne(id int64) (*Note, error) {
  for _, item := range r.storage {
    if item.Id == id {
      return &item, nil
    }
  }
  return nil, fmt.Errorf("not found")
}

func (r *NotesRepoMem) GetMany() ([]Note, error) {
  return r.storage, nil
}

func (r *NotesRepoMem) Update(id int64, note NoteInput) (*Note, error) {
  for i, item := range r.storage {
    if item.Id == id {
      r.storage[i].Text = note.Text
      r.storage[i].Dt = note.Dt
      return &r.storage[i], nil
    }
  }
  return nil, fmt.Errorf("not found")
}

func (r *NotesRepoMem) Delete(id int64) error {
  initialLen := len(r.storage)
  r.storage = slices.DeleteFunc(r.storage, func(item Note) bool {
    return item.Id == id
  })
  if initialLen == len(r.storage) {
    return fmt.Errorf("not found")
  }
  return nil
}
