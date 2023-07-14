package resthandler

import (
	"net/http"
	"strconv"
)

// UserBlock - UserBlock
func (s Handler) UserBlock(w http.ResponseWriter, r *http.Request) {
	userID := r.FormValue("user_id") // userID will be "" if parameter is not set
	if userID == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	index, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	index, err = s.user.Block(index)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	s.log.Info(index)
	w.WriteHeader(http.StatusOK)
}
