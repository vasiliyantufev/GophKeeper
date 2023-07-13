package resthandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// TokenGetList - TokenGetList
func (s Handler) TokenGetList(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userId")

	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		s.log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tokenList, err := s.token.GetList(id)
	if err != nil {
		s.log.Error(err)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	resp, err := json.Marshal(tokenList)
	if err != nil {
		s.log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
