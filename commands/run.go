package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newRunCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run server to preview pages",
		Run:   runHandler,
	}

	addStrictFlag(cmd.Flags())

	return cmd
}

func runHandler(cmd *cobra.Command, args []string) {
	fmt.Println("run >>>")
}

// const t = `
// <!DOCTYPE html>
// <html>
// 	<head>
// 		<meta charset="UTF-8">
// 		<title>{{.Title}}</title>
// 	</head>
// 	<body>
// 		<h2>{{.Title}}</h2>
// 		{{.Content}}
// 	</body>
// </html>
// `

// 	tmpl := template.Must(template.New("rootTemplate").Parse(t))

// 	http.HandleFunc("/", rootHandler(tmpl))
// 	log.Fatal(http.ListenAndServe(":9000", nil))

// func rootHandler(tmpl *template.Template) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		data := struct {
// 			Title   string
// 			Content string
// 		}{
// 			Title:   "Pony",
// 			Content: "Minimalistic static site generator in Go",
// 		}
// 		if err := tmpl.Execute(w, data); err != nil {
// 			log.Println(err)
// 		}
// 	}
// }
