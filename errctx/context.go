package errctx

import "github.com/tvarney/maputil/mpath"

// Context is an error context.
type Context struct {
	Path    *mpath.Path
	Handler ErrorHandler

	errCount int
	lastErr  error
}

// New returns a new context.
func New(handlers ...ErrorHandler) *Context {
	switch len(handlers) {
	case 0:
		return &Context{
			Path:    mpath.New(mpath.DotNotation{}),
			Handler: nil,
		}
	case 1:
		return &Context{
			Path:    mpath.New(mpath.DotNotation{}),
			Handler: handlers[0],
		}
	}
	return &Context{
		Path:    mpath.New(mpath.DotNotation{}),
		Handler: &MultiHandler{Handlers: handlers},
	}
}

// ErrorCount returns the total count of errors this context has handled.
func (ctx *Context) ErrorCount() int {
	return ctx.errCount
}

// LastError returns the last error that this context handled.
func (ctx *Context) LastError() error {
	return ctx.lastErr
}

// Reset resets the context error count and last error values.
func (ctx *Context) Reset() {
	ctx.lastErr = nil
	ctx.errCount = 0
}

// Error handles an error.
func (ctx *Context) Error(err error) {
	if err == nil {
		return
	}

	ctx.errCount++
	ctx.lastErr = err
	if ctx.Handler != nil {
		ctx.Handler.Add(ctx.Path, err)
	}
}

// ErrorWith handles an error for the given new element.
func (ctx *Context) ErrorWith(err error, elem mpath.Element) {
	if err == nil {
		return
	}

	ctx.errCount++
	ctx.lastErr = err
	if ctx.Handler != nil {
		ctx.Path.Add(elem)
		ctx.Handler.Add(ctx.Path, err)
		ctx.Path.Pop()
	}
}

// ErrorWithKey handles an error for the given key.
func (ctx *Context) ErrorWithKey(err error, key string) {
	ctx.ErrorWith(err, mpath.Key(key))
}

// ErrorWithIndex handles an error for the given index.
func (ctx *Context) ErrorWithIndex(err error, idx int) {
	ctx.ErrorWith(err, mpath.Index(idx))
}
