package types

type Song struct {
	Artist   string `json:"artist"`
	Title    string `json:"title"`
	Metadata string `json:"metadata"`
	Year     int    `json:"year"`
}
