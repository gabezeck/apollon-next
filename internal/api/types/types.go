package types

import (
	LTypes "github.com/gabezeck/test-api/internal/manager/lastfm/types"
	RTypes "github.com/gabezeck/test-api/internal/manager/reddit/types"
)

type PostsRequest struct {
	Sub   string `json:"sub"`
	Limit int    `json:"limit"`
	Time  string `json:"time"`
}

type Song struct {
	TrackInfo *RTypes.Song        `json:"trackInfo"`
	MetaData  *LTypes.LastFMTrack `json:"metadata"`
}

type SongsResponse struct {
	Songs []*Song `json:"songs"`
}
