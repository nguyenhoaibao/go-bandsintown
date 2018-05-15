package model

type Artist struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	URL                string `json:"url"`
	ImageURL           string `json:"image_url"`
	ThumbURL           string `json:"thumb_url"`
	FacebookPageURL    string `json:"facebook_page_url"`
	MBID               string `json:"mbid"`
	TrackerCount       uint   `json:"tracker_count"`
	UpcomingEventCount uint   `json:"upcoming_event_count"`
}
