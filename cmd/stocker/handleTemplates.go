package main

import (
	// "html/template"
	// "net/http"
	// "path/filepath"
	//"go-stocker/templates"
	"fmt"
	"net/http"
	"os"
	//"context"
)

type Page struct {
	Title string
	Body  []byte
}

// handling persistent storage
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := LoadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func editHandler(w http.ResponseWriter, r http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := LoadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	fmt.Fprintf(w, "<h1>Editing %s</h1>"+
	"<form action=\"/save/%s\" method=\"POST\">"+
	"<textarea name=\"body\">%s</textarea><br>"+
	"<input type=\"submit\" value=\"Save\">"+
	"</form>",
	p.Title, p.Title, p.Body)
}

// func LoginButtonLandingPage() templ.Component {
// 	//templates.RegisterButton("Register").Render(context.Background(), os.Stdout)
// 	//return templates.LoginButton("Login")
// 	return templ.ComponentGroup(
// 		templates.LoginButton("Login"),
// 		templates.RegisterButton("Register"),
// 	)
// }
// // loadTemplates parses templates once and organizes them
// func LoadTemplates() {
// 	templates = make(map[string]*template.Template)

// 	pages := []string{"home", "about"}

// 	for _, page := range pages {
// 		tmpl := template.Must(template.ParseFiles(
// 			//filepath.Join("templates", "layouts", "base.html"),
// 			filepath.Join("../../templates/layouts/base.html"),
// 			filepath.Join("../../templates/pages/home.html"),

// 			//filepath.Join("templates", "pages", page+".html"),
// 		))
// 		templates[page] = tmpl
// 	}
// }

// func HomeHandler(w http.ResponseWriter, r *http.Request) {
// 	data := struct {
// 		Title    string
// 		Username string
// 	}{
// 		Title:    "Home",
// 		Username: "John Doe",
// 	}
// 	RenderTemplate(w, "home", data)
// }

// func AboutHandler(w http.ResponseWriter, r *http.Request) {
// 	data := struct {
// 		Title string
// 	}{
// 		Title: "About",
// 	}
// 	RenderTemplate(w, "about", data)
// }

// // renderTemplate renders the chosen page
// func RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
// 	tmpl, ok := templates[name]
// 	if !ok {
// 		http.Error(w, "The page does not exist.", http.StatusNotFound)
// 		return
// 	}
// 	err := tmpl.Execute(w, data)
// 	if err != nil {
// 		http.Error(w, "Template execution error.", http.StatusInternalServerError)
// 	}
// }
