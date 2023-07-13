package resthandler

import (
	"html/template"
	"net/http"
)

type ViewData struct {
	Users map[int64]string
}

// IndexHandler - the page that displays all users
func (s Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(s.config.TemplatePath)
	if err != nil {
		s.log.Errorf("Parse failed: %s", err)
		http.Error(w, "Error loading index page", http.StatusInternalServerError)
		return
	}

	users := make(map[int64]string)

	usersDb, err := s.user.GetAllUsers()
	if err != nil {
		s.log.Errorf("Execution failed: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, user := range usersDb {
		users[user.ID] = user.Username
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	data := ViewData{Users: users}
	err = tmpl.Execute(w, data)
	if err != nil {
		s.log.Errorf("Execution failed: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
