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
	_ "github.com/lib/pq"
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
	e.GET("/dbtest", func(c echo.Context) error {
		r := TableTest()
		return c.String(http.StatusOK, r)
	})
	e.Logger.Fatal(e.Start(":8080"))
}

func TableTest() string {
	dsn := os.Getenv("DSN")
	fmt.Printf("DSN=%s\n", dsn)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbmap.AddTableWithName(Comment{}, "comments")
	err = dbmap.CreateTablesIfNotExists()

	err = dbmap.Insert(&Comment{Text: "こんにちは"})
	if err != nil {
		log.Fatal(err)
		return fmt.Sprintf("Error :%v", err)
	}

	return "success"
}
