package types

type Album struct {
	Position string `json:"position"`
	Artist   string `json:"artist"`
	Title    string `json:"title"`
	Mbid     string `json:"mbid"`
	Url      string `json:"url"`
	Images   Images `json:"image"`
}

type Artist struct {
	Name string `json:"name"`
	Mbid string `json:"mbid"`
	Url  string `json:"url"`
}

type Images []Image

type Image struct {
	Size string `json:"size"`
	Url  string `json:"url"`
}

type LastFMTrack struct {
	Id            string  `json:"id"`
	Name          string  `json:"name"`
	Mbid          string  `json:"mbid"`
	Url           string  `json:"url"`
	Duration      string  `json:"duration"`
	PlayCount     string  `json:"playcount"`
	UserPlayCount string  `json:"userplaycount"`
	UserLoved     string  `json:"userloved"`
	Artist        Artist  `json:"artist"`
	Album         Album   `json:"album"`
	TopTags       TopTags `json:"toptags>tag"`
	Wiki          Wiki    `json:"wiki"`
}

type Tags struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type TopTags []Tags

type Wiki struct {
	Published string `json:"published"`
	Summary   string `json:"summary"`
	Content   string `json:"content"`
}
