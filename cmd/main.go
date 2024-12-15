package main

import (
	//"net/http"

	//"github.com/gin-gonic/gin"
	"github.com/laysaalves/todolist-go/internal"
)

func main() {
  /* r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.Run()*/

  internal.ConnDB()
}