package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-gorp/gorp"
	"github.com/labstack/echo"
)

type Comment struct {
	Id      int64     `db:"id,primarykey,autoincrement"`
	Name    string    `db:"name,notnull,default:'名無し',size:200"`
	Text    string    `db:"text,notnull,size:400"`
	Created time.Time `db:"created,notnull"`
	Updated time.Time `db:"updated,notnull"`
}

func main() {
	fmt.Println("Hello go")

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Go")
	})
	e.Logger.Fatal(e.Start(":8080"))
}

func tableTest() {
	dsn := os.Getenv("DSN")
	var err error
	db, _ := sql.Open("postgres", dsn)
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbmap.AddTableWithName(Comment{}, "comments")
	err = dbmap.CreateTablesIfNotExists()

	err = dbmap.Insert(&Comment{Text: "こんにちは"})
	if err != nil {
		log.Fatal(err)
	}
}
