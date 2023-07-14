package resthandler

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ViewDataToken struct {
	Tokens map[string]string
}

// TokenList -  the page that displays all tokens user
func (s Handler) TokenList(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userId")

	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		s.log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles(s.config.TemplatePathToken)
	if err != nil {
		s.log.Errorf("Parse failed: %s", err)
		http.Error(w, "Error loading user list page", http.StatusInternalServerError)
		return
	}

	tokens := make(map[string]string)

	tokenList, err := s.token.GetList(id)
	if err != nil {
		s.log.Error(err)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	for _, token := range tokenList {
		tokens[token.AccessToken] = token.EndDateAt.String()
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	data := ViewDataToken{Tokens: tokens}
	err = tmpl.Execute(w, data)
	if err != nil {
		s.log.Errorf("Execution failed: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
