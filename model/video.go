package model

// Video represents video model.
type Video struct {
	ID       string    `json:"id"`
	Preview  string    `json:"preview"`
	Duration int64     `json:"duration"`
	Name     string    `json:"name"`
	URL      *VideoURL `json:"url,omitempty"`
}

// VideoURL represents video URL model.
type VideoURL struct {
	ViewCount int64  `json:"viewCount"`
	URL       string `json:"url"`
	Type      string `json:"type"`
}
