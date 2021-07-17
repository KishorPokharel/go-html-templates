package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates map[string]*template.Template

func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	templates["allPosts"] = template.Must(template.ParseFiles("templates/layout.html", "templates/navbar.html", "templates/posts/all.html"))
	templates["getPost"] = template.Must(template.ParseFiles("templates/layout.html", "templates/navbar.html", "templates/posts/single.html"))
	log.Print("Templates parsed")
}

type post struct {
	Title string
	Body string
}

func getPost(w http.ResponseWriter, r *http.Request) {
	p := post{Title: "Single Post", Body: "Body of Single Post"}
	data := struct {
		Title string
		Post post
	} {
		Title: p.Title,
		Post: p,
	}
	tmpl := templates["getPost"]
	err := tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Print(err)
		return
	}
	log.Print("Template executed")
}

func allPosts(w http.ResponseWriter, r *http.Request) {
	posts := []post{
		{Title: "First Post", Body: "Body of First Post"},
		{Title: "Second Post", Body: "Body of Second Post"},
	}

	data := struct {
		Title string
		Posts []post
	} {
		Title: "All Posts",
		Posts: posts,
	}
	tmpl := templates["allPosts"]
	err := tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Print(err)
		return
	}
	log.Print("Template executed")
} 

func main() {
	http.HandleFunc("/posts", allPosts)
	http.HandleFunc("/posts/view", getPost)

	log.Print("Server Running...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}