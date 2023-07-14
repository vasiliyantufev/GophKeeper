package resthandler

import (
	"net/http"
	"strconv"
)

// UserUnblock - UserUnblock
func (s Handler) UserUnblock(w http.ResponseWriter, r *http.Request) {
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
	index, err = s.user.Unblock(index)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	s.log.Info(index)
	w.WriteHeader(http.StatusOK)
}
