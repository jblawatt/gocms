package main

import (
	"./api"
	"./views"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

// http://jmoiron.github.io/sqlx/

// DatabaseMiddleware adds a database instance to the context for using
// it in the different views.
func DatabaseMiddleware(filename string) gin.HandlerFunc {
	return func(c *gin.Context) {
		dbase, err := sqlx.Open("sqlite3", filename)
		if err != nil {
			panic(err)
		}
		c.Set("dbase", dbase)
		c.Next()
	}
}

var schema = `

CREATE TABLE IF NOT EXISTS pages (
    id INTEGER PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    lang VARCHAR(2) NOT NULL DEFAULT 'de',
    date DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS tags (
    page_id INTEGER,
    slug VARCHAR(100),
    FOREIGN KEY(page_id) REFERENCES pages(id)
);

CREATE UNIQUE INDEX IF NOT EXISTS unique_tags_page_id_slug ON tags(page_id, slug);

CREATE TABLE IF NOT EXISTS content (
    id INTEGER PRIMARY KEY,
    page_id INTEGER,
    sort INTEGER,
    content TEXT,
    FOREIGN KEY(page_id) REFERENCES page(id)
);

`

// InitializeDatabase opens the database an ensures that the schmea is applied.
func InitializeDatabase(filename string) {
	dbase, errOpen := sqlx.Open("sqlite3", filename)
	if errOpen != nil {
		panic(errOpen)
	}
	defer dbase.Close()
	_, errExec := dbase.Exec(schema)
	if errExec != nil {
		panic(errExec)
	}
}

func main() {

	const DATABASE string = "data.sqlite3"

	InitializeDatabase(DATABASE)

	engine := gin.Default()
	engine.Use(DatabaseMiddleware(DATABASE))
	engine.LoadHTMLGlob("templates/**/*")

	engine.GET("/", views.Index)
	engine.GET("/blog", views.BlogIndex)
	engine.GET("/blog/:title", views.BlogDetail)
	// engine.GET("/blog/:title/edit", views.BlogEdit)

	// engine.GET("/:title", views.PageDetail)

	apiGroup := engine.Group("/api")
	{
		apiGroup.GET("/blog", api.BlogList)
		apiGroup.GET("/blog/:title", api.BlogDetail)
	}

	ip := os.Getenv("IP")
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = ":3000"
	}

	engine.Run(ip + port)

}
