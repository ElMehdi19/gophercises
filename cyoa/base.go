package cyoa

import (
	"html/template"
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

const AltTemplate string = `
<!DOCTYPE html>
<html lang="en" dir="ltr">
    <head>
        <meta charset="utf-8">
        <title>Choose your own adventure</title>
    </head>
    <body>
		<h1>Chapter is about: <span>{{ .Title }}</span></h1>
    </body>
</html>
`
