package main

import (
	"WebApplication/objects"
	"WebApplication/server"
	"fmt"
)

func main() {
	page := &objects.Page{Title: "hi", Body: []byte("This is a sample page")}
	page.Save()
	p, err := server.LoadPage("hi")
	fmt.Println(err)
	fmt.Println(string(p.Body))
	handler := new(server.Handler)
	handler.HandleClientRequest(server.View(""))
	handler.HandleClientRequest(server.Edit(""))
}
