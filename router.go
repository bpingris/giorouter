package giorouter

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

// Route
type Route interface {
	Layout(gtx layout.Context) layout.Dimensions
}

// Routes
type Routes map[string]Route

// Router holds the precious information of the router
type Router struct {
	C      chan byte
	Routes Routes
	Stack  []string
	Th     *material.Theme
}

// NewRouter returns an instance of a Router
func NewRouter(th *material.Theme) Router {
	return Router{
		C:      make(chan byte, 1),
		Routes: make(Routes),
		Th:     th,
	}
}

// SetRoutes links the provided paths to the Route
func (r *Router) SetRoutes(routes Routes, initialRoute string) {
	r.Routes = routes
	r.Push(initialRoute)
}

// Push add a route to the stack and ask for a redraw
func (r *Router) Push(route string) {
	if _, ok := r.Routes[route]; ok {
		r.Stack = append(r.Stack, route)
		r.Redraw()
	}
}

// Top returns the current route
func (r Router) Top() string {
	l := len(r.Stack)
	if l == 0 {
		return ""
	}
	return r.Stack[l-1]
}

// Pop pops a route out of the Routes stack and ask for a redraw (basically go back)
func (r *Router) Pop() string {
	if !r.CanPop() {
		return ""
	}
	l := len(r.Stack)
	top := r.Stack[l-1]
	r.Stack = r.Stack[:l-1]
	r.Redraw()
	return top
}

// Redraw send an event through the Router'c Channel to ask for an immediate redraw
func (r *Router) Redraw() {
	r.C <- 0
}

// CanPop verifies if a previous route is available
func (r *Router) CanPop() bool {
	return len(r.Stack) > 1
}

// Layout draw the current route
func (r *Router) Layout(gtx layout.Context) layout.Dimensions {
	return r.Routes[r.Top()].Layout(gtx)
}
