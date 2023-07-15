package resthandler

import (
	"net/http"
)

// UserBlock - UserBlock
func (s Handler) UserBlock(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username") // username will be "" if parameter is not set
	if username == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	index, err := s.user.Block(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	s.log.Info(index)
	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
