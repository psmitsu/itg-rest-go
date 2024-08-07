package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/psmitsu/itglobal-go-example/model"
)

func assertJSON(actual []byte, data interface{}, t *testing.T) {
  expected, err := json.Marshal(data)
  if err != nil {
    t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
  }

  if bytes.Compare(expected, actual) != 0 {
    t.Errorf("the expected json: %s is different from actual %s", expected, actual)
  }
}

func TestPost(t *testing.T) {
  // setup mock db
  db, mock, err := sqlmock.New()
  if err != nil {
    log.Fatal("error creating stub db", err)
  }
  defer db.Close()

  // setup input and expected output
  currentTime := time.Now()
  input := model.NoteInput{
    Text: "Test",
    Dt: currentTime,
  }
  inputJson, _ := json.Marshal(input)
  expected := model.Note{
    Id: 1,
    Text: "Test",
    Dt: currentTime,
  }

  // setup the application and response/request objects
  app := setupRouter(db)
  w := httptest.NewRecorder()
  req, _ := http.NewRequest("POST", "/notes", strings.NewReader(string(inputJson)))

  // setup db expectations
  mock.ExpectExec("^INSERT INTO notes*").
    WithArgs(input.Text, input.Dt.Format(time.RFC3339)).
    WillReturnResult(sqlmock.NewResult(1,1))

  // test request
  app.ServeHTTP(w, req)
  if w.Code != http.StatusCreated {
    t.Fatalf("expected status code to be %d, but got: %d", http.StatusCreated, w.Code)
  }

  // assert http response body
  assertJSON(w.Body.Bytes(), expected, t)

  // assert sql expectations
  if err := mock.ExpectationsWereMet(); err != nil {
    t.Errorf("there were unfulfilled expectations: %s", err)
  }
}

// func TestGetMany(t *testing.T)
// func TestPatch(t *testing.T)
// func TestDelete(t *testing.T)
