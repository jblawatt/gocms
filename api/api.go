package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BlogList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func BlogDetail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}