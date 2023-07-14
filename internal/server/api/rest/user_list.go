package resthandler

import (
	"html/template"
	"net/http"
)

type ViewDataUser struct {
	Users map[string]string
}

// UserListHandler - the page that displays all users
func (s Handler) UserListHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(s.config.TemplatePathUser)
	if err != nil {
		s.log.Errorf("Parse failed: %s", err)
		http.Error(w, "Error loading user list page", http.StatusInternalServerError)
		return
	}

	users := make(map[string]string)

	usersDb, err := s.user.UserList()
	if err != nil {
		s.log.Errorf("Execution failed: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, user := range usersDb {
		users[user.Username] = user.DeletedAt.String()
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	data := ViewDataUser{Users: users}
	err = tmpl.Execute(w, data)
	if err != nil {
		s.log.Errorf("Execution failed: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
