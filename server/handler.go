package server

import (
	"WebApplication/objects"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Handler struct {
}

func LoadPage(title string) (*objects.Page, error) {
	fileName := title + ".txt"
	body, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return &objects.Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	fmt.Println(title)
	//p, _ := LoadPage(title)
	//fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", "title", "boy")
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := LoadPage(title)
	if err != nil {
		p = &objects.Page{Title: title}
	}
	fmt.Fprintf(w, "<h1>Editing%s<h1>", p.Title)
}

type Request interface {
	handle()
}

type View string

type Edit string

func (v View) handle() {
	fmt.Println("view")
	http.HandleFunc("/view", viewHandler)
	s := http.ListenAndServe(":808", nil)
	fmt.Println(s)
}

func (e Edit) handle() {
	fmt.Println("edit")
	http.HandleFunc("/edit", editHandler)
	log.Fatal(http.ListenAndServe(":808", nil))
}

func (h Handler) HandleClientRequest(request Request) {
	request.handle()
}
