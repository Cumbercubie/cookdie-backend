package models

type EventType string

type Event struct {
	Type       EventType              `json:"type"`
	EventId    string                 `json:"event_id"`
	UserId     string                 `json:"user_id"`
	BusinessId string                 `json:"business_id"`
	Date       string                 `json:"date"`
	Payload    map[string]interface{} `json:"payload"`
}

type Review struct {
	UserId     string  `json:"user_id"`
	BusinessId string  `json:"business_id"`
	Date       string  `json:"date"`
	ReviewId   string  `json:"review_id,omitempty"`
	Stars      float64 `json:"stars,omitempty"`
	Useful     int64   `json:"useful,omitempty"`
	Funny      int64   `json:"funny,omitempty"`
	Cool       int64   `json:"cool,omitempty"`
	Text       string  `json:"text,omitempty"`
}

type Checkin struct {
	
}
