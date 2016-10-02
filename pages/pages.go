package pages

import (
	"github.com/jmoiron/sqlx"
)

type Content struct {
	Id   string
	Sort int
	Text string
}

type Page struct {
	Id      string `db:"id"`
	Title   string `db:"title"`
	Date    string "db:date"
	Tags    []string
	Lang    string `db:lang`
	Content []Content
}

func GetPage(dbase *sqlx.DB, title string) (*Page, error) {
	page := Page{}
	err := dbase.Get(&page, "SELECT id, title, date, lang FROM pages WHERE title = ?", title)
	if err != nil {
		return nil, err
	}
	return &page, nil

}

func CreatePage(dbase *sqlx.DB, page *Page) {

}

func DeletePage(dbase *sqlx.DB, page *Page) {

}

func UpdatePage(dbase *sqlx.DB, page *Page) {

}
