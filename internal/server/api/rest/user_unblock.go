package resthandler

import (
	"net/http"
)

// UserUnblock - UserUnblock
func (s Handler) UserUnblock(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username") // username will be "" if parameter is not set
	if username == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	index, err := s.user.Unblock(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	s.log.Info(index)
	w.WriteHeader(http.StatusOK)
}
