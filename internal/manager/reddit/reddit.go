package reddit

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/gabezeck/test-api/internal/deps"
	"github.com/gabezeck/test-api/internal/manager/reddit/types"

	"github.com/thecsw/mira"
	"go.uber.org/zap"
)

type Mgr interface {
	GetSongsFromPosts(sub string, limit int, time string) ([]*types.Song, error)
}

type RMgr struct {
	RClient *mira.Reddit
	Logger  *zap.SugaredLogger
}

func New(d *deps.Deps) *RMgr {
	return &RMgr{
		RClient: d.RClient,
		Logger:  d.Logger,
	}
}

func (r *RMgr) GetSongsFromPosts(sub string, limit int, time string) ([]*types.Song, error) {
	posts, err := getPosts(r.RClient, sub, limit, time)
	if err != nil {
		r.Logger.Errorf("Failed to retrieve posts: %v", err)
		return nil, err
	}

	if len(posts) == 0 {
		return nil, errors.New("no posts found")
	}

	songs, err := parseSongData(posts, r.Logger)
	if err != nil {
		return nil, errors.New("failed to parse song information from post titles")
	}

	return songs, nil
}

func getPosts(client *mira.Reddit, sub string, limit int, time string) ([]mira.PostListingChild, error) {
	sort := "top"
	subs, err := client.GetSubredditPosts(sub, sort, time, limit)

	if err != nil {
		return nil, err
	}

	return subs, nil
}

func parseSongData(posts []mira.PostListingChild, log *zap.SugaredLogger) ([]*types.Song, error) {
	//Matches songs in Reddit post titles, e.g.
	// Song Title - Artist [Location, Genres, Other Metadata] (Year)
	re := regexp.MustCompile(`^(.+) \- (.+) \[(.+)\] \(([0-9]+)\)`)

	var songs []*types.Song
	for _, post := range posts {
		matches := re.FindAllStringSubmatch(post.GetTitle(), -1)

		if len(matches) > 0 && len(matches[0]) >= 2 {
			artist := matches[0][1]
			title := matches[0][2]
			year, err := strconv.Atoi(matches[0][4])
			if err != nil {
				log.Debugf("Failed to convert year string to int: %v", err)
			}

			if artist != "" && title != "" {
				songs = append(songs, &types.Song{
					Artist:   matches[0][1],
					Title:    matches[0][2],
					Metadata: matches[0][3],
					Year:     year,
				})
			}
		}
	}

	return songs, nil
}
