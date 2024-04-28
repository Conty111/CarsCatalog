package car

// MsgResponse represents default successful response
type MsgResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// CarUpdates is struct for update endpoint
type CarUpdates struct {
	Mark    string `json:"mark"`
	Model   string `json:"model"`
	Year    int    `json:"year"`
	RegNum  string `json:"regNum"`
	OwnerID string `json:"ownerID"`
}
