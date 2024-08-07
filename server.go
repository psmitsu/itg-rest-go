package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/psmitsu/itglobal-go-example/controller"
	"github.com/psmitsu/itglobal-go-example/model"
)

func setupRouter(db *sql.DB) *gin.Engine {
  var (
    router = gin.Default()
    repo = &model.NotesRepoSql{Db: db}
    cont = controller.Controller{Repo: repo}
  )

  cont.SetupRoutes(router)
  // router.Run("localhost:8088")

  return router;
}

func main() {
  db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
  if err != nil {
    log.Fatal("could not open db", err)
  }
  defer db.Close()

  router := setupRouter(db)
  router.Run("localhost:8088")
}
