package dao

type Profile struct {
	Email     string  `json:"email"`
	PushToken string  `json:"push_token"`
	Events    []Event `json:"events"`
}

type Group struct {
	GroupName  string  `json:"group_name"`
	GroupOwner string  `json:"group_owner"`
	Events     []Event `json:"events"`
}

type Event struct {
	EventTitle string `json:"event_title"`
	EventDesc  string `json:"event_desc"`
	Location   string `json:"location"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
}
