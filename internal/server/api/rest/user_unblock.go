package resthandler

import (
	"io"
	"net/http"
)

// UserUnblock - UserUnblock
func (s Handler) UserUnblock(w http.ResponseWriter, r *http.Request) {

	resp, err := io.ReadAll(r.Body)
	if err != nil {
		s.log.Error(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	s.log.Info(resp)
	w.WriteHeader(http.StatusOK)
}
