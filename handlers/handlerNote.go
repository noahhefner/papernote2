package handlers

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"

	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"
)

type viewContext struct {
	Filename string
	Content  template.HTML
}

func HandleGetEditor(w http.ResponseWriter, r *http.Request) {

	filename := r.FormValue("filename")

	path := path.Join("notes", filename)

	content, err := os.ReadFile(path)

	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t, err := template.ParseFiles("templates/pages/editor.html")
	if err != nil {

		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, string(content))
	if err != nil {

		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func HandleGetRendered(w http.ResponseWriter, r *http.Request) {

	filename := r.FormValue("filename")

	path := path.Join("notes", filename)

	content, err := os.ReadFile(path)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	parsed := p.Parse(content)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	templateContext := viewContext{
		Filename: filename,
		Content:  template.HTML(markdown.Render(parsed, renderer)),
	}

	t, err := template.ParseFiles("templates/pages/view.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, templateContext)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
