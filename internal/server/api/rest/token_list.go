package resthandler

import (
	"html/template"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/vasiliyantufev/gophkeeper/internal/client/storage/layouts"
)

type ViewDataToken struct {
	Tokens map[string]string
}

// TokenList -  the page that displays all tokens user
func (s Handler) TokenList(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	tmpl, err := template.ParseFiles(s.config.TemplatePathToken)
	if err != nil {
		s.log.Errorf("Parse failed: %s", err)
		http.Error(w, "Error loading user list page", http.StatusInternalServerError)
		return
	}

	tokens := make(map[string]string)

	userID, err := s.user.GetUserID(username)
	if err != nil {
		s.log.Error(err)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	tokenList, err := s.token.GetList(userID)
	if err != nil {
		s.log.Error(err)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	for _, token := range tokenList {
		now := time.Now().Format(layouts.LayoutDateAndTime.ToString())
		end := token.EndDateAt.Format(layouts.LayoutDateAndTime.ToString())
		check := now > end

		s.log.Debug(now)
		s.log.Debug(end)
		s.log.Debug(check)

		if check {
			tokens[token.AccessToken] = "Block"
		} else {
			tokens[token.AccessToken] = "Active"
		}
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
