package cyoa

import (
    "encoding/json"
    "io"
)

func JsonStoryParser(r io.Reader) (Story, error) {
        decoder := json.NewDecoder(r)
        story := Story{}
        if err := decoder.Decode(&story); err != nil {
            return nil, err
        }
        return story, nil
}

type Story map[string]Chapter

type Chapter struct {
    Title string `json:"title"`
    Paragraphs []string `json:"story"`
    Options []Option `json:"options"`
}

type Option struct {
    Text string `json:"text"`
    Chapter string `json:"arc"`
}
