package cyoa

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

func init() {
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

// WithTemplate function used to render a custom template
func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.t = t
	}
}

// NewHandler function to customize the default handler with HandlerOption
func NewHandler(s Story, opts ...HandlerOption) http.Handler {
	h := handler{s, tpl}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

func (h handler) chapterPath(r *http.Request) (*Chapter, error) {
	path := r.URL.Path

	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]

	chapter, ok := h.s[path]
	if !ok {
		return nil, fmt.Errorf(fmt.Sprintf("Chapter `%s` doesn't exist in story", path))
	}
	return &chapter, nil

}

func logRequest(r *http.Request) {
	path := r.URL.Path
	method := r.Method
	remote := r.RemoteAddr
	log.Printf("%s %s FROM %s", method, path, remote)
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	chapter, err := h.chapterPath(r)
	if err != nil {
		log.Printf("%v", err)
		http.Error(w, fmt.Sprintf("%v", err), http.StatusNotFound)
		return
	}

	if err := h.t.Execute(w, *chapter); err != nil {
		log.Printf("%v", err)
		http.Error(w, "Bad shit happened", http.StatusInternalServerError)
	}
}
