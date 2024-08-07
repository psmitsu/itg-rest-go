package model

import (
	"database/sql"
	"fmt"
	"time"
)

type NotesRepoSql struct {
  Db *sql.DB
}

func MakeNotesRepoSql(db *sql.DB) *NotesRepoSql {
  return &NotesRepoSql{Db: db}
}

func (r *NotesRepoSql) Sync() error {
  _, err := r.Db.Exec("CREATE TABLE IF NOT EXISTS notes (id SERIAL PRIMARY KEY, text TEXT NOT NULL, dt TIMESTAMPTZ NOT NULL)")

  if err != nil {
    return err
  }

  return nil
}

func (r *NotesRepoSql) Create(input NoteInput) (*Note, error) {
  result, err := r.Db.Exec("INSERT INTO notes(text, dt) VALUES (?, ?)", input.Text, input.Dt.Format(time.RFC3339))

  if err != nil {
    return nil, err
  }

  id, err := result.LastInsertId()
  if err != nil {
    return nil, err
  }

  return &Note{
    Id: id,
    Text: input.Text,
    Dt: input.Dt,
  }, nil
}

func (r *NotesRepoSql) GetOne(id int64) (*Note, error) {
  var note Note

  row := r.Db.QueryRow("SELECT * FROM notes WHERE id = ?", id)
  if err := row.Scan(&note.Id, &note.Text, &note.Dt); err != nil {
    return nil, err
  }

  return &note, nil
}

func (r *NotesRepoSql) GetMany() ([]Note, error) {
  var notes []Note

  rows, err := r.Db.Query("SELECT * FROM notes")
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  for rows.Next() {
    var note Note
    if err := rows.Scan(&note.Id, &note.Text, &note.Dt); err != nil {
      return nil, err
    }
    notes = append(notes, note)
  }

  if err := rows.Err(); err != nil {
    return nil, err
  }

  return notes, nil
}

func (r *NotesRepoSql) Update(id int64, input NoteInput) (*Note, error) {
  result, err := r.Db.Exec("UPDATE notes SET text=?, dt=? WHERE id=?", id, input.Text, input.Dt)

  if err != nil {
    return nil, err
  }

  naffected, err := result.RowsAffected()
  if err != nil {
    return nil, err
  }

  if naffected == 0 {
    return nil, fmt.Errorf("not found")
  }

  return &Note {
    Id: id,
    Text: input.Text,
    Dt: input.Dt,
  }, nil
}

func (r *NotesRepoSql) Delete(id int64) error {
  result, err := r.Db.Exec("DELETE FROM notes WHERE id=?", id)
  if err != nil {
    return err
  }
  naffected, err := result.RowsAffected()
  if err != nil {
    return err
  }
  if naffected == 0 {
    return fmt.Errorf("not found")
  }
  return nil
}
