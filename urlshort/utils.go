package urlshort

import (
	"encoding/json"
	"net/http"
)

// MapHanlder function to redirect a path to a url
func MapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.RequestURI
		if dest, ok := pathToUrls[path]; ok == true {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// JSONHandler function to parse paths and links from json file and passes them to MapHandler
func JSONHandler(jsonStr []byte, fallback http.Handler) (http.HandlerFunc, error) {
	paths := []pathToUrl{}
	if err := json.Unmarshal(jsonStr, &paths); err != nil {
		return nil, err
	}
	pathToUrls := map[string]string{}
	for _, pu := range paths {
		pathToUrls[pu.Path] = pu.Dest
	}

	return MapHandler(pathToUrls, fallback), nil
}

type pathToUrl struct {
	Path string `json:"path"`
	Dest string `json:"dest"`
}
