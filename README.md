# pongoal
Demo of a controller for use of [Pongo2](https://github.com/flosch/pongo2) template engine
with your [Goal Toolkit](https://github.com/colegion/goal) based application.

## Use
Assuming that you are starting with the [automatically generated app](https://github.com/colegion/goal/tree/master/internal/skeleton),
update `Controllers` type in your `controllers/init.go`:
```diff
import (
+	"github.com/alkchr/pongoal"
+
	"github.com/colegion/contrib/controllers/requests"
	"github.com/colegion/contrib/controllers/sessions"
-	"github.com/colegion/contrib/controllers/templates"
)
...
type Controllers struct {
	*requests.Requests
-	*templates.Templates
+	*pongoal.Pongo2
	*sessions.Sessions
}
```

It will make available `Render` and `RenderTemplate` methods inside your app.
You may use them as follows:
```go
// App is a sample controller.
type App struct {
	*Controllers
}

// Index is an action that renders a home page.
func (c *App) Index() http.Handler {
	c.Context["someKey1"] = "someValue1"
	...
	c.Context["someKeyN"] = "someValueN"
	return c.Render()
}
```
Pongo's ["first impression template"](./example/views/App/Index.html)'s
rendering is showed [here](./example/controllers/app.go#L21).
