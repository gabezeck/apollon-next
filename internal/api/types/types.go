package types

import (
	RTypes "github.com/gabezeck/test-api/internal/manager/reddit/types"
)

type PostsRequest struct {
	Sub   string `json:"sub"`
	Limit int    `json:"limit"`
	Time  string `json:"time"`
}

type SongsResponse struct {
	Songs []*RTypes.Song `json:"songs"`
}
