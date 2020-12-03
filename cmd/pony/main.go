package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	const t = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		<h2>{{.Title}}</h2>
		{{.Content}}
	</body>
</html>
`

	tmpl := template.Must(template.New("rootTemplate").Parse(t))

	http.HandleFunc("/", rootHandler(tmpl))
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func rootHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Title   string
			Content string
		}{
			Title:   "Pony",
			Content: "Minimalistic static site generator in Go",
		}
		if err := tmpl.Execute(w, data); err != nil {
			log.Println(err)
		}
	}
}

// Commands:
// build 	- builds static pages.
// run 		- runs preview.
// verify - verifies that meta matches templates and vice versa.
//					outputs: found not used meta key, found key in template without meta.
//
// Flags:
// strict - meta must match template, used by [build|run], fatal when not matching.
