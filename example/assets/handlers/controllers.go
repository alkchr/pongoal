// Package handlers is generated automatically by goal toolkit.
// Please, do not edit it manually.
package handlers

import (
	"net/http"
	"net/url"

	c1 "github.com/alkchr/pongoal/example/assets/handlers/github.com/alkchr/pongoal"
	c0 "github.com/alkchr/pongoal/example/assets/handlers/github.com/colegion/contrib/controllers/requests"
	c2 "github.com/alkchr/pongoal/example/assets/handlers/github.com/colegion/contrib/controllers/sessions"
	contr "github.com/alkchr/pongoal/example/controllers"

	"github.com/colegion/goal/strconv"
)

// Controllers is an insance of tControllers that is automatically generated from Controllers controller
// being found at "github.com/alkchr/pongoal/example/controllers/init.go",
// and contains methods to be used as handler functions.
//
// Controllers is a struct that should be embedded into every controller
// of your app to make methods and fields provided by standard controllers available.
var Controllers tControllers

// context stores names of all controllers and packages of the app.
var context = url.Values{}

// tControllers is a type with handler methods of Controllers controller.
type tControllers struct {
}

// New allocates (github.com/alkchr/pongoal/example/controllers).Controllers controller,
// initializes its parents; then returns the controller.
func (t tControllers) New(w http.ResponseWriter, r *http.Request, ctr, act string) *contr.Controllers {
	c := &contr.Controllers{}
	c.Requests = c0.Requests.New(w, r, ctr, act)
	c.Pongo2 = c1.Pongo2.New(w, r, ctr, act)
	c.Sessions = c2.Sessions.New(w, r, ctr, act)
	return c
}

// Before is a method that is started by every handler function at the very beginning
// of their execution phase no matter what.
func (t tControllers) Before(c *contr.Controllers, w http.ResponseWriter, r *http.Request) http.Handler {
	// Execute magic Before actions of embedded controllers.
	if h := c0.Requests.Before(c.Requests, w, r); h != nil {
		return h
	}

	if h := c1.Pongo2.Before(c.Pongo2, w, r); h != nil {
		return h
	}

	if h := c2.Sessions.Before(c.Sessions, w, r); h != nil {
		return h
	}

	// Call magic Before action of (github.com/alkchr/pongoal/example/controllers).Before.
	if h := c.Before(); h != nil {
		return h
	}

	return nil
}

// After is a method that is started by every handler function at the very end
// of their execution phase no matter what.
func (t tControllers) After(c *contr.Controllers, w http.ResponseWriter, r *http.Request) (h http.Handler) {

	// Execute magic After methods of embedded controllers.

	if h = c0.Requests.After(c.Requests, w, r); h != nil {
		return h
	}

	if h = c1.Pongo2.After(c.Pongo2, w, r); h != nil {
		return h
	}

	if h = c2.Sessions.After(c.Sessions, w, r); h != nil {
		return h
	}

	return
}

// Init is used to initialize controllers of "github.com/alkchr/pongoal/example/controllers"
// and its parents.
func Init() {

	initApp()

	initControllers()

	contr.Init(context)

}

func initControllers() {

	c0.Init()

	c1.Init()

	c2.Init()

}

func init() {
	_ = strconv.MeaningOfLife
}
