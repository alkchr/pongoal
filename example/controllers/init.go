package controllers

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/alkchr/pongoal"

	"github.com/colegion/contrib/controllers/requests"
	"github.com/colegion/contrib/controllers/sessions"
)

var (
	pagesPath = flag.String("static.pages.path", "./static/pages", "path to a dir with static pages")
)

// Controllers is a struct that should be embedded into every controller
// of your app to make methods and fields provided by standard controllers available.
type Controllers struct {
	*requests.Requests
	*pongoal.Pongo2
	*sessions.Sessions
}

// Before is a magic action that is executed on every request
// before any other action.
//
// Only structures with at least one action are treated as controllers.
// So, do not delete this method.
func (c *Controllers) Before() http.Handler {
	return nil
}

// Init will create a simple static page.
func Init(_ url.Values) {
	for tplName, destFile := range map[string]string{
		"App/Index.html": "sample.html",
	} {
		err := renderTplToFile("App/Index.html", "sample.html", map[string]interface{}{
			"key1": "value1",
			"smth": "...",
			"keyN": "valueN",
		})
		if err != nil {
			log.Panicf(`Failed to render "%s" to "%s". Error: %v.`, tplName, destFile, err)
		}
	}
}

// renderTplToFile renders the requested template to the requested
// file. It returns an error if parameters are not correct.
// It may be safely used inside Init function of this controller.
func renderTplToFile(tplName, fileName string, c map[string]interface{}) error {
	// Make sure the template does exist.
	t, ok := pongoal.Templates[tplName]
	if !ok {
		return fmt.Errorf(`Template "%s" does not exist.`, tplName)
	}

	// Try to create the requested file.
	f, err := os.Create(filepath.Join(*pagesPath, fileName))
	if err != nil {
		return err
	}
	defer f.Close()

	// Render the template.
	res, err := t.ExecuteBytes(c)
	if err != nil {
		return err
	}

	// Write result to the file.
	_, err = f.Write(res)
	return err
}
