package cyoa


type handler struct {
	s Story
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
