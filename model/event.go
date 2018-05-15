package model

import "time"

type Event struct {
	ID             string   `json:"id"`
	ArtistID       string   `json:"artist_id"`
	URL            string   `json:"url"`
	OnSaleDatetime string   `json:"on_sale_datetime"`
	Datetime       DateTime `json:"datetime"`
	Description    string   `json:"description"`
	Venue          Venue    `json:"venue"`
	Offers         []struct {
		Type   string `json:"type"`
		URL    string `json:"url"`
		Status string `json:"status"`
	} `json:"offers"`
	Lineup []string `json:"lineup"`
}

type DateTime struct {
	time.Time
}

func (t *DateTime) UnmarshalJSON(data []byte) error {
	current := string(data[1 : len(data)-1])
	t1, err := time.Parse("2006-01-02T15:04:05", current)

	if err != nil {
		t1, err = time.Parse(time.RFC3339, current)
		if err != nil {
			return err
		}
	}

	*t = DateTime{t1}

	return nil
}

func (t *DateTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(time.RFC3339)+2)
	b = append(b, '"')
	b = t.AppendFormat(b, time.RFC3339)
	b = append(b, '"')
	return b, nil
}
