package internal

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/labstack/gommon/log"
)

var teamMembers = []string{
	"Felix Kühner",
	"Alexander Karuth",
	"David Duck",
	"Alexander Zimmermann",
	"Assaad Naes",
	"Dominik Hermann",
	"Marten Bichel",
	"Moritz Sanden",
	"Hannes Harnisch",
}

type Api struct {
	logger *log.Logger
}

func NewApi() *Api {
	return &Api{
		logger: log.New("api"),
	}
}

func (a *Api) JsonWriter(w http.ResponseWriter, status int, obj interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(obj); err != nil {
		a.logger.Error("failed to write response: %s", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
	}
}

func (a *Api) HandleGetMember(w http.ResponseWriter, r *http.Request) {
	memberName := r.URL.Query().Get("member")

	for _, member := range teamMembers {
		if strings.EqualFold(member, memberName) {
			a.JsonWriter(w, http.StatusOK, memberName+" is a member of the team")
			return
		}
	}
	http.Error(w, memberName+" isn't member of the team", http.StatusBadRequest)
}
