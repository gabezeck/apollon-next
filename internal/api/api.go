package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

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
	r.HandleFunc("/songs-from-posts", func(w http.ResponseWriter, r *http.Request) { SongsFromPostsHandler(w, r, api) })
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func SongsFromPostsHandler(w http.ResponseWriter, r *http.Request, api *API) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var req types.PostsRequest
	err = json.Unmarshal(b, &req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	songs, err := api.RMgr.GetSongsFromPosts(req.Sub, req.Limit, req.Time)
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

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
