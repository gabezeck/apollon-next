package api

import (
	"net/http"

	"github.com/gabezeck/test-api/internal/deps"
	"github.com/gabezeck/test-api/internal/manager/lastfm"
	"github.com/gabezeck/test-api/internal/manager/reddit"
	"github.com/gorilla/mux"
)

type ApiI interface {
}

type API struct {
	Deps *deps.Deps
	LMgr *lastfm.LMgr
	RMgr *reddit.RMgr
}

func New(deps *deps.Deps) *API {
	return &API{
		Deps: deps,
		LMgr: lastfm.New(deps),
		RMgr: reddit.New(deps),
	}
}

func RegisterRoutes(r *mux.Router, deps *deps.Deps) {
	api := New(deps)

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/songs-from-posts", func(w http.ResponseWriter, r *http.Request) {
		SongsFromPostsHandler(w, r, api)
	}).Methods("GET")

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
