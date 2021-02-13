# Giorouter

Giorouter is package meant to simplify routing in a Gioui app.

It provides a simple API to use:
- `Push(route string)`
- `Pop()`
- `CanPop() bool`
`

## Usage

```go
type Home struct {
    Router *giorouter.Router
    click *widget.Clickable
}

func (h *Home) Layout(gtx layout.Context, th *material.Theme) {
    if h.click.Clicked() {
        h.Router.Push("about")
    }
    // Draw your page...
}

type About struct {
    Router *giorouter.Router
}

func (a *About) Layout(gtx layout.Context, th *material.Theme) {
    if a.click.Clicked() {
        a.Router.Pop()
    }
    // Draw your page...
}

func main() {
  	th := material.NewTheme(gofont.Collection())
	router := giorouter.NewRouter(th)

    home := Home{Router: router}
    about := About{Router: router}

    router.SetRoutes(giorouter.Routes{
        "home": &home,
        "about": &about,
    }, "home")

    go func() {
		w := app.NewWindow()
		var ops op.Ops

		for {
			select {
			case e := <-w.Events():
				switch e := e.(type) {
				case system.FrameEvent:
					gtx := layout.NewContext(&ops, e)
					router.Layout(gtx, e) // Draw the current page
					e.Frame(gtx.Ops)
				case system.DestroyEvent:
					os.Exit(0)
				}
			case <-router.C:
                // Redraw the screen when the route is changing
				w.Invalidate()
			}
		}
    }()
    app.Main()
}
```
