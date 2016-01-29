// Package pongoal is an abstraction around Pongo2
// for use with Goal Toolkit.
package pongoal

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/flosch/pongo2"
)

var (
	views    = flag.String("pongoal:path", "./views/", "path to the directory with views")
	contType = flag.String("pongoal:content.type", "text/html; charset=utf-8", "Content-Type header's value")
	defTpl   = flag.String("pongoal:default.pattern", "%v/%v.html", "default template's name pattern")
)

// Pongo2 is a controller that provides support of pongo2
// templates to your Goal Toolkit based application.
type Pongo2 struct {
	// Context is used for passing variables to templates.
	Context pongo2.Context

	// StatusCode is a status code that will be returned when rendering.
	// If not specified explicitly, 200 will be used.
	StatusCode int

	// defTpl is a name of the template that will be rendered
	// if no other templates are specified explicitly.
	defTpl string
}

// Initially allocates Context and prepares the name of template to render
// if it isn't specified explicitly.
func (c *Pongo2) Initially(w http.ResponseWriter, r *http.Request, a []string) bool {
	// Set the default template to render.
	c.defTpl = fmt.Sprintf(*defTpl, a[0], a[1])

	// Allocate a new context.
	c.Context = pongo2.Context{}
	return false
}

// RenderTemplate is an action that gets a path to template
// and renders it using data from Context.
func (c *Pongo2) RenderTemplate(templatePath string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Prepare status of the response.
		if c.StatusCode == 0 {
			c.StatusCode = http.StatusOK
		}

		// If requested template exists, render it.
		if t, ok := Templates[templatePath]; ok {
			// Write necessary headers.
			w.Header().Set("Content-Type", *contType)
			w.WriteHeader(c.StatusCode)

			// Execute requested template.
			if err := t.ExecuteWriter(c.Context, w); err != nil {
				http.Error(
					w,
					fmt.Sprintf(`Failed to execute "%s". Error: "%v".`, templatePath, err),
					http.StatusInternalServerError,
				)
			}
			return
		}

		// Otherwise, show an internal server error.
		http.Error(w, fmt.Sprintf(`Template "%s" does not exist.`, templatePath), http.StatusInternalServerError)
	})
}

// Render is an equivalent of the following:
//	RenderTemplate(CurrentController + "/" + CurrentAction + ".html")
// The default path pattern may be overriden by adding the following
// line to your configuration file:
//	[templates]
//	default.pattern = %s/%s.tpl
func (c *Pongo2) Render() http.Handler {
	return c.RenderTemplate(c.defTpl)
}

// Init triggers loading of templates.
func Init(_ url.Values) {
	// Scan requested views directory and its subdirectories for template files.
	// Ignore errors.
	dir := filepath.FromSlash(*views)
	filepath.Walk(dir, func(p string, info os.FileInfo, err error) error {
		// If there is an error or current object is a directory, do nothing.
		if err != nil || info.IsDir() {
			return nil
		}

		// Get current path without the dir at the beginning.
		// E.g. if we're scanning "views/app/index.html", extract
		// the "app/index.html" part only.
		rel, _ := filepath.Rel(dir, p)
		relNorm := filepath.ToSlash(rel)

		// Try to parse the template.
		log.Printf(`Parsing template "%s" => "%s"...`, relNorm, p)
		t, err := pongo2.FromFile(p)
		if err != nil {
			log.Printf(`Failed to parse "%s" => "%s". Error: "%v".`, relNorm, p, err)
		}
		Templates[relNorm] = t
		return nil
	})
}

// Templates stores Pongo2 templates along with their names.
var Templates = map[string]*pongo2.Template{}
