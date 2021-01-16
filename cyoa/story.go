package cyoa

import (
	"encoding/json"
	"io"
	"net/http"
	"html/template"
	"log"
)

var tpl *template.Template

const baseTemplate string = `
<!DOCTYPE html>
<html lang="en" dir="ltr">
    <head>
        <meta charset="utf-8">
        <title>Choose your own adventure</title>
    </head>
    <body>
        <h1>{{ .Title }}</h1>
        {{ range .Paragraphs }}
            <p>{{ . }}</p>
        {{ end }}
        <ul>
            {{ range .Options }}
                <a href="/{{ .Chapter }}"><li>{{ .Text }}</li></a>
            {{ end }}
        </ul>
    </body>
</html>
`

func init(){
	tpl = template.Must(template.New("").Parse(baseTemplate))
}

// JSONStoryParser func
func JSONStoryParser(r io.Reader) (Story, error) {
	decoder := json.NewDecoder(r)
	story := Story{}
	if err := decoder.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]

	if chapter, ok := h.s[path]; ok {
		if err := tpl.Execute(w, chapter); err != nil {
			log.Printf("%v", err)
			http.Error(w, "Bad shit happened", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "You're lost", http.StatusNotFound)
}

// Story alias
type Story map[string]Chapter

// Chapter struct
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

// Option struct
type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
