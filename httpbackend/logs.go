package httpbackend

import (
	"encoding/json"
	"net/http"

	"github.com/khades/servbot/models"

	"github.com/khades/servbot/repos"

	"goji.io/pat"
)

func logs(w http.ResponseWriter, r *http.Request, s *models.HTTPSession) {
	channel := pat.Param(r, "channel")
	user := pat.Param(r, "user")
	if channel == "" || user == "" {
		http.Error(w, "URL is not valid", http.StatusBadRequest)
		return
	}
	userLogs, error := repos.GetUserMessageHistory(&user, &channel)
	if error != nil {
		http.Error(w, error.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(*userLogs)
}
