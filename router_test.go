package giorouter

import "testing"

const (
	initialRoute = "home"
	gotoRoute    = "test"
	dummyRoute   = "dummy"
)

func newRouter() *Router {
	r := NewRouter(nil)
	r.SetRoutes(Routes{
		gotoRoute:    nil,
		initialRoute: nil,
	}, initialRoute)

	go func() {
		for range r.C {
		}
	}()
	return r
}

func TestInitialRoute(t *testing.T) {
	r := newRouter()
	if r.Top() != initialRoute {
		t.Errorf("Invalid current route, expected %s got %s", initialRoute, r.Top())
	}
}

func TestPushRoute(t *testing.T) {
	r := newRouter()
	r.Push(gotoRoute)

	top := r.Top()
	if top != gotoRoute {
		t.Errorf("Invalid top route, expected %s got %s", gotoRoute, top)
	}
}

func TestPopRoute(t *testing.T) {
	r := newRouter()

	if len(r.Stack) != 1 {
		t.Errorf("Invalid stack's length, expected %d got %d", 1, len(r.Stack))
	}

	pop := r.Pop()
	if pop != "" {
		t.Errorf("Invalid pop, expected %s got %s", "''", pop)
	}
	if len(r.Stack) != 1 {
		t.Errorf("Invalid stack's length, expected %d got %d", 1, len(r.Stack))
	}

	r.Push(gotoRoute)
	if len(r.Stack) != 2 {
		t.Errorf("Invalid stack's length, expected %d got %d", 2, len(r.Stack))
	}

	pop = r.Pop()
	if pop == "" {
		t.Errorf("Invalid pop, expected %s got %s", "'something'", pop)
	}
}

func TestStackEmpty(t *testing.T) {
	r := newRouter()
	r.Stack = r.Stack[:0]
	top := r.Top()
	if top != "" {
		t.Errorf("Invalid top, expected %s got %s", "''", top)
	}
}
