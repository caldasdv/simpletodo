package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	InitDatabase()
	defer DB.Close()
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.GET("/", func(ctx *gin.Context) {
		todos := ReadToDoList()
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"todos": todos,
		})
	})
	e.POST("/todos", func(ctx *gin.Context) {
		title := ctx.PostForm("title")
		status := ctx.PostForm("status")
		id, _ := CreateToDo(title, status)
		ctx.HTML(http.StatusOK, "task.html", gin.H{
			"title":  title,
			"status": status,
			"id":     id,
		})
	})
	e.DELETE("/todos/:id", func(ctx *gin.Context) {
		param := ctx.Param("id")
		id, _ := strconv.ParseInt(param, 10, 64)
		DeleteToDo(id)
	})
	e.Run(":8080")
}
