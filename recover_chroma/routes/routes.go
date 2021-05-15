package routes

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/ElMehdi19/gophercises/recover_chroma/utils"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Home page</h1>")
}

func PanicDemo(w http.ResponseWriter, r *http.Request) {
	utils.FuncThatPanics()
}

func SourceCodeHandler(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("path")
	lineNum, err := strconv.Atoi(r.FormValue("line"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	file, err := os.Open(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	lines := [][2]int{
		{lineNum, lineNum},
	}
	formatter := html.New(html.WithLineNumbers(true), html.TabWidth(2), html.HighlightLines(lines))
	style := styles.Get("github")

	lexer := lexers.Get("go")
	iterator, err := lexer.Tokenise(nil, buf.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("content-type", "text/html")
	formatter.Format(w, style, iterator)
	// quick.Highlight(w, buf.String(), "go", "html", "github")
}
