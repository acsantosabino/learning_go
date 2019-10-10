package main

import (
    "page"
    "regexp"
	"log"
    "net/http"
    "io"
    "html/template"
    "bytes"
)

var validPath = regexp.MustCompile("^/(edit|save|view)?/?([a-zA-Z0-9]+)$")
var pageList = regexp.MustCompile("(\\[PagesList\\])")
var fileName = regexp.MustCompile("^([a-zA-Z0-9]+).txt$")
var filepath = "../tmpl/"
var templates = template.Must(template.ParseFiles(filepath+"edit.html", filepath+"view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *page.Page) {
    html := new(bytes.Buffer)
	err := templates.ExecuteTemplate(html, tmpl + ".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    finalResponse := string(pageList.ReplaceAllFunc(html.Bytes(),func(s []byte) []byte{
        return interPages()
    }))
	_, err = io.WriteString(w, finalResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
    p, err := page.LoadPage(title)
    if err != nil {
        http.Redirect(w, r, "/edit/"+title, http.StatusFound)
        return
    }
    renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
    p, err := page.LoadPage(title)
    if err != nil {
        p = &page.Page{Title: title}
    }
    renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
    body := r.FormValue("body")
    p := &page.Page{Title: title, Body: []byte(body)}
    err := p.Save()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func rootHandler(w http.ResponseWriter, r *http.Request, title string) {
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        m := validPath.FindStringSubmatch(r.URL.Path)
        if m == nil {
            http.NotFound(w, r)
            return
        }
        fn(w, r, m[2])
    }
}

func interPages() []byte{
    pages, _ := page.AllPages()
    var list string

    for i := 0; i < len(pages); i++ {
        m := fileName.FindStringSubmatch(pages[i])
        if m != nil {
            list += "<li><a href=\"/view/"+m[1]+"\">"+m[1]+"</a></li>"
        }
    }
    return []byte(list)
}

func main() {
    page.CheckPagesPath()
    http.HandleFunc("/", makeHandler(rootHandler))
    http.HandleFunc("/view/", makeHandler(viewHandler))
    http.HandleFunc("/edit/", makeHandler(editHandler))
    http.HandleFunc("/save/", makeHandler(saveHandler))

    log.Fatal(http.ListenAndServe(":8080", nil))
}