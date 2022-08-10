package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gabezeck/test-api/internal/api/types"
	"github.com/gabezeck/test-api/internal/deps"
	"github.com/gabezeck/test-api/internal/manager/reddit"
	"github.com/gorilla/mux"
)

type ApiI interface {
}

type API struct {
	Deps *deps.Deps
	RMgr *reddit.RMgr
}

func New(deps *deps.Deps) *API {
	return &API{
		Deps: deps,
		RMgr: reddit.New(deps),
	}
}

func RegisterRoutes(r *mux.Router, deps *deps.Deps) {
	api := New(deps)

	r.HandleFunc("/", HomeHandler)
	r.
		HandleFunc("/songs-from-posts", func(w http.ResponseWriter, r *http.Request) { SongsFromPostsHandler(w, r, api) }).
		Methods("GET")

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func SongsFromPostsHandler(w http.ResponseWriter, r *http.Request, api *API) {
	query := r.URL.Query()

	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil {
		limit = 10
		api.Deps.Logger.Debug("Failed to convert limit string to int. Defaulting to 10 posts.")
	}

	songs, err := api.RMgr.GetSongsFromPosts(query.Get("sub"), limit, query.Get("time"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	output, err := json.Marshal(&types.SongsResponse{
		Songs: songs,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
