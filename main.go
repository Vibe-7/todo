package main

import (
	"todo/http"
	"todo/todo1"
)

func main() {
	todoList := todo1.NewList()
	HttpHandlers := http.NewHttpHandlers(todoList)
	HttpServer := http.NewHTTPServer(HttpHandlers)

	HttpServer.StartServer()
}
