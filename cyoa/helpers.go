package cyoa

import (
    "encoding/json"
    "html/template"
    "net/http"
    "log"
    "io"
)


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
