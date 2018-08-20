package api

import (
	"fmt"
	"io"
	"text/tabwriter"
)

// Endpoints is a list of Endpoints, which adds helper methods to the list.
type Endpoints []Endpoint

// Print will write the server endpoints to an io.Writer, which can be used for debugging purposes.
// A common usage is to write to stdout (`endpoints.Print(os.Stdout)`)
func (e Endpoints) Print(writer io.Writer) {
	fmt.Fprint(writer, "Server endpoints:")
	w := tabwriter.NewWriter(writer, 10, 0, 2, ' ', 0)
	for _, endpoint := range e {
		fmt.Fprintln(w, fmt.Sprintf("%s\t%s\t%s", endpoint.Path(), endpoint.Method(), endpoint.HandlerName()))
	}
	w.Flush()
}
