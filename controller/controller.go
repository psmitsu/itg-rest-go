package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/psmitsu/itglobal-go-example/model"
)

// Don't need that, it is for clarity/future
type IController interface {
  SetupRoutes(routes gin.IRouter)
}

type Controller struct {
  Repo model.INotesRepository
}

func (cont *Controller) SetupRoutes(router gin.IRouter) {
  router.POST("/notes", cont.postNote)
  router.GET("/notes", cont.getNotes)
  router.GET("/notes/:id", cont.getNoteById)
  router.PATCH("/notes/:id", cont.patchNote)
  router.DELETE("/notes/:id", cont.deleteNote)
}

func (cont *Controller) postNote(c *gin.Context) {
  var (
    noteInput model.NoteInput
    note *model.Note
    err error
  )

  if err = c.BindJSON(&noteInput); err != nil {
    // TODO: handle error
    fmt.Println(err)
    return
  }

  note, err = cont.Repo.Create(noteInput)

  if err != nil {
    // TODO: handle error
    fmt.Println(err)
    return
  }

  c.JSON(http.StatusCreated, note)
}

func (cont *Controller) getNoteById(c *gin.Context) {
  var (
    id int
    note *model.Note
    err error
  )

  id, err = strconv.Atoi(c.Param("id"))

  if err != nil {
    // TODO: handle error
    fmt.Println(err)
    c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
    return
  }

  note, err = cont.Repo.GetOne(int64(id))

  if err != nil {
    // TODO: handle error
    fmt.Println(err)
    return
  }

  c.JSON(http.StatusOK, note)
}

func (cont *Controller) getNotes(c *gin.Context) {
  notes, err := cont.Repo.GetMany()

  if err != nil {
    // TODO: handle error
    fmt.Println(err)
    return
  }

  c.JSON(http.StatusOK, notes)
}

func (cont *Controller) patchNote(c *gin.Context) {
  var (
    id int
    noteInput model.NoteInput
    updatedNote *model.Note
    err error
  )

  id, err = strconv.Atoi(c.Param("id"))
  if err != nil {
    // TODO: handle error
    fmt.Println("id convert error", err)
    return
  }

  if err = c.BindJSON(&noteInput); err != nil {
    // TODO: handle error
    fmt.Println("bind json error", err)
    return
  }

  updatedNote, err = cont.Repo.Update(int64(id), noteInput)
  if err != nil {
    // TODO: handle error
    fmt.Println("update error", err)
    return
  }

  c.JSON(http.StatusOK, updatedNote)
}

func (cont *Controller) deleteNote(c *gin.Context) {
  var (
    id int
    err error
  )

  id, err = strconv.Atoi(c.Param("id"))

  if err != nil {
    // TODO: handle error
    fmt.Println(err)
    return
  }

  err = cont.Repo.Delete(int64(id))

  if err != nil {
    // TODO: handle error
    fmt.Println(err)
    return
  }

  c.JSON(http.StatusOK, gin.H{"message" : "deleted"})
}
