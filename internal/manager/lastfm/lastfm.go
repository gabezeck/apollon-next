package lastfm

import (
	"errors"

	"github.com/gabezeck/test-api/internal/deps"
	LTypes "github.com/gabezeck/test-api/internal/manager/lastfm/types"
	RTypes "github.com/gabezeck/test-api/internal/manager/reddit/types"
	"github.com/shkh/lastfm-go/lastfm"
	"go.uber.org/zap"
)

type Mgr interface {
	GetSongDetails(song *RTypes.Song) (lastfm.TrackGetInfo, error)
}

type LMgr struct {
	LClient *lastfm.Api
	Logger  *zap.SugaredLogger
}

func New(d *deps.Deps) *LMgr {
	return &LMgr{
		LClient: d.LClient,
		Logger:  d.Logger,
	}
}

func (r *LMgr) GetSongDetails(song *RTypes.Song) (*LTypes.LastFMTrack, error) {
	query := lastfm.P{"artist": song.Artist, "track": song.Title}
	result, err := r.LClient.Track.GetInfo(query)
	if err != nil {
		return nil, errors.New("failed to retrieve track information from Last.fm")
	}

	var images LTypes.Images
	for _, image := range result.Album.Images {
		images = append(images, LTypes.Image{
			Size: image.Size,
			Url:  image.Url,
		})
	}

	var tags LTypes.TopTags
	for _, tag := range result.TopTags {
		tags = append(tags, LTypes.Tags{
			Name: tag.Name,
			Url:  tag.Url,
		})
	}

	track := &LTypes.LastFMTrack{
		Id:            result.Id,
		Name:          result.Name,
		Mbid:          result.Mbid,
		Url:           result.Url,
		Duration:      result.Duration,
		PlayCount:     result.PlayCount,
		UserPlayCount: result.UserPlayCount,
		UserLoved:     result.UserLoved,
		Artist: LTypes.Artist{
			Name: result.Artist.Name,
			Mbid: result.Artist.Mbid,
			Url:  result.Artist.Url,
		},
		Album: LTypes.Album{
			Position: result.Album.Position,
			Artist:   result.Album.Artist,
			Title:    result.Album.Title,
			Mbid:     result.Album.Mbid,
			Url:      result.Album.Url,
			Images:   images,
		},
		TopTags: tags,
		Wiki: LTypes.Wiki{
			Published: result.Wiki.Published,
			Summary:   result.Wiki.Summary,
			Content:   result.Wiki.Content,
		},
	}

	return track, nil
}
