package errctx

import (
	"fmt"
	"io"

	"github.com/tvarney/maputil/mpath"
)

// ErrorHandler is an interface used by the Context type used to handle errors.
type ErrorHandler interface {
	Add(*mpath.Path, error)
}

// ErrorDiscarder is an ErrorHandler which discards the error.
type ErrorDiscarder struct{}

// Add does nothing with the given error.
func (h ErrorDiscarder) Add(p *mpath.Path, err error) {}

// ErrorPrinter is an ErrorHandler which prints errors to an io.Writer.
type ErrorPrinter struct {
	Stream io.Writer
}

// Add writes the given error to the internal stream.
func (h *ErrorPrinter) Add(p *mpath.Path, err error) {
	fmt.Fprintf(h.Stream, "%s: %s\n", p.String(), err.Error())
}

// MultiHandler is an ErrorHandler which allows grouping multiple
// ErrorHandlers.
type MultiHandler struct {
	Handlers []ErrorHandler
}

// NewMultiHandler returns a new MultiHandler with the given set of handlers.
func NewMultiHandler(handlers ...ErrorHandler) *MultiHandler {
	return &MultiHandler{
		Handlers: handlers,
	}
}

// Add dispatches the given error to all child handlers.
func (m *MultiHandler) Add(p *mpath.Path, err error) {
	for _, h := range m.Handlers {
		h.Add(p, err)
	}
}
