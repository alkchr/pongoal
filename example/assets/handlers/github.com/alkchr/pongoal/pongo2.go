// Package handlers is generated automatically by goal toolkit.
// Please, do not edit it manually.
package handlers

import (
	"net/http"
	"net/url"

	contr "github.com/alkchr/pongoal"

	"github.com/colegion/goal/strconv"
)

// Pongo2 is an insance of tPongo2 that is automatically generated from Pongo2 controller
// being found at "github.com/alkchr/pongoal/pongoal.go",
// and contains methods to be used as handler functions.
//
// Pongo2 is a controller that provides support of pongo2
// templates to your Goal Toolkit based application.
var Pongo2 tPongo2

// context stores names of all controllers and packages of the app.
var context = url.Values{}

// tPongo2 is a type with handler methods of Pongo2 controller.
type tPongo2 struct {
}

// New allocates (github.com/alkchr/pongoal).Pongo2 controller,
// then returns it.
func (t tPongo2) New(w http.ResponseWriter, r *http.Request, ctr, act string) *contr.Pongo2 {
	c := &contr.Pongo2{

		Action: act,

		Controller: ctr,
	}
	return c
}

// Before is a method that is started by every handler function at the very beginning
// of their execution phase no matter what.
func (t tPongo2) Before(c *contr.Pongo2, w http.ResponseWriter, r *http.Request) http.Handler {

	// Call magic Before action of (github.com/alkchr/pongoal).Before.
	if h := c.Before(); h != nil {
		return h
	}

	return nil
}

// After is a method that is started by every handler function at the very end
// of their execution phase no matter what.
func (t tPongo2) After(c *contr.Pongo2, w http.ResponseWriter, r *http.Request) (h http.Handler) {

	return
}

// RenderTemplate is a handler that was generated automatically.
// It calls Before, After methods, and RenderTemplate action found at
// github.com/alkchr/pongoal/pongoal.go
// in appropriate order.
//
// RenderTemplate is an action that gets a path to template
// and renders it using data from Context.
func (t tPongo2) RenderTemplate(w http.ResponseWriter, r *http.Request) {
	var h http.Handler
	c := Pongo2.New(w, r, "Pongo2", "RenderTemplate")
	defer func() {
		if h != nil {
			h.ServeHTTP(w, r)
		}
	}()
	defer Pongo2.After(c, w, r)
	if res := Pongo2.Before(c, w, r); res != nil {
		h = res
		return
	}
	if res := c.RenderTemplate(
		strconv.String(r.Form, "templatePath"),
	); res != nil {
		h = res
		return
	}
}

// Render is a handler that was generated automatically.
// It calls Before, After methods, and Render action found at
// github.com/alkchr/pongoal/pongoal.go
// in appropriate order.
//
// Render is an equivalent of the following:
//	RenderTemplate(CurrentController + "/" + CurrentAction + ".html")
// The default path pattern may be overriden by adding the following
// line to your configuration file:
//	[templates]
//	default.pattern = %s/%s.tpl
func (t tPongo2) Render(w http.ResponseWriter, r *http.Request) {
	var h http.Handler
	c := Pongo2.New(w, r, "Pongo2", "Render")
	defer func() {
		if h != nil {
			h.ServeHTTP(w, r)
		}
	}()
	defer Pongo2.After(c, w, r)
	if res := Pongo2.Before(c, w, r); res != nil {
		h = res
		return
	}
	if res := c.Render(); res != nil {
		h = res
		return
	}
}

// Init is used to initialize controllers of "github.com/alkchr/pongoal"
// and its parents.
func Init() {

	initPongo2()

	contr.Init(context)

}

func initPongo2() {

	context.Add("Pongo2", "RenderTemplate")

	context.Add("Pongo2", "Render")

}

func init() {
	_ = strconv.MeaningOfLife
}
