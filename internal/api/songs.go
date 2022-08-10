package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gabezeck/test-api/internal/api/types"
)

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

	var songsRes []*types.Song
	for _, song := range songs {
		track, err := api.LMgr.GetSongDetails(song)
		if err != nil {
			api.Deps.Logger.Errorf("\nFailed to get track info: %v", err)
		}

		songsRes = append(songsRes, &types.Song{
			TrackInfo: song,
			MetaData:  track,
		})
	}

	output, err := json.Marshal(&types.SongsResponse{
		Songs: songsRes,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
