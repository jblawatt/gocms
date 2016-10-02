package views

import (
	"../pages"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
)

// Index -- Main View
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

// BlogIndex -- Main view from blog
func BlogIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "blog/index.tmpl", gin.H{})
}

// BlogDetail -- Detail page for a blog
func BlogDetail(c *gin.Context) {
	title := c.Param("title")
	dbase := c.MustGet("dbase").(*sqlx.DB)
	page, ok := pages.GetPage(dbase, title)
	if ok != nil {
		c.HTML(http.StatusNotFound, "not_found.tmpl", gin.H{"title": title})
	} else {
		c.HTML(http.StatusOK, "blog/detail.tmpl", gin.H{"page": page})
	}
}

// BlogEdit
func BlogEdit(c *gin.Context) {
	c.HTML(http.StatusOK, "blog_edit.tmpl", gin.H{})
}

// PageDetail -- Some other Page
func PageDetail(c *gin.Context) {
	title := c.Param("title")
	c.HTML(http.StatusOK, "page.tmpl", gin.H{"title": title})
}
