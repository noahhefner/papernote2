package handlers

import (

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"

	"path"
	"os"
	"html/template"
	"io/ioutil"
	"net/http"
)

type editorContext struct {
	Content string
}

type renderedContentContext struct {
	Filename	string
	Content template.HTML
}

func HandleHome(w http.ResponseWriter, r *http.Request) {

	var fileNames []string

	files, err := ioutil.ReadDir("./notes")
	if err != nil {
		http.Error(w, "Failed to read notes.", http.StatusInternalServerError)
		return
	}

	for _, f := range files {
		fileNames = append(fileNames, f.Name())
	}

	t, err := template.ParseFiles("templates/pages/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, fileNames)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func HandleGetEditor (w http.ResponseWriter, r *http.Request) {

	filename := r.FormValue("filename")

	path := path.Join("notes", filename)

	content, err := os.ReadFile(path)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templateContext := editorContext {
		Content: string(content),
	}

  t, err := template.ParseFiles("templates/components/editor.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, templateContext)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}


}

func HandleGetContentsRendered(w http.ResponseWriter, r *http.Request) {
	
	filename := r.FormValue("filename")

	path := path.Join("notes", filename)

	content, err := os.ReadFile(path)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(content)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	templateContext := renderedContentContext {
		Filename: filename,
		Content: template.HTML(markdown.Render(doc, renderer)),
	}

	t, err := template.ParseFiles("templates/components/noteContent.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, templateContext)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

