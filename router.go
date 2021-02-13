package giorouter

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type Route interface {
	Layout(gtx layout.Context, th *material.Theme)
}

type Routes map[string]Route

type Router struct {
	C      chan byte
	Routes Routes
	Stack  []string
	th     *material.Theme
}

func NewRouter(th *material.Theme) *Router {
	return &Router{
		C:      make(chan byte, 1),
		Routes: make(Routes),
		th:     th,
	}
}

func (r *Router) SetRoutes(routes Routes, initialRoute string) {
	r.Routes = routes
	r.Push(initialRoute)
}

func (r *Router) Push(route string) {
	r.Stack = append(r.Stack, route)
	r.Redraw()
}

func (r *Router) Top() string {
	l := len(r.Stack)
	if l == 0 {
		return ""
	}
	return r.Stack[l-1]
}

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

func (r *Router) Redraw() {
	r.C <- 0
}

func (r *Router) CanPop() bool {
	return len(r.Stack) > 1
}

func (r *Router) Layout(gtx layout.Context) {
	r.Routes[r.Top()].Layout(gtx, r.th)
}
