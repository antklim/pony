package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newRunCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run server to preview pages",
		RunE:  runHandler,
	}

	addStrictFlag(cmd.Flags())

	return cmd
}

func runHandler(cmd *cobra.Command, args []string) error {
	fmt.Println("run handler >>>")
	// addr := ":9000"
	// router := &internal.Router{}

	// if _, err := os.Stat(meta); err != nil {
	// 	return errors.Wrap(err, "metadata file read failed")
	// }

	// if _, err := os.Stat(tmpl); err != nil {
	// 	return errors.Wrap(err, "template file read failed")
	// }

	// if err := router.LoadMeta(meta); err != nil {
	// 	return err
	// }

	// if err := router.LoadTemplate(tmpl); err != nil {
	// 	return err
	// }

	// mux := http.NewServeMux()

	// for route, handler := range router.PreviewRoutes() {
	// 	mux.Handle(route, handler)
	// }

	// s := &http.Server{
	// 	Addr:           addr,
	// 	Handler:        mux,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }

	// log.Printf("pony preview is listening at %s", addr)
	// log.Fatal(s.ListenAndServe())
	return nil
}
