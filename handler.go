package pongoal

import (
	"net/http"

	"github.com/flosch/pongo2"
)

// Handler implements http.Handler interface.
type Handler struct {
	context  pongo2.Context // Variables to be passed to the template.
	template string         // Path to the template to be rendered.
	status   int            // Expected status code of the response.
}

// Apply writes to response the result received from action.
func (t *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
