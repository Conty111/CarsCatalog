package helpers

type CarUpdates struct {
	Mark    string `jsonapi:"mark"`
	Model   string `jsonapi:"model"`
	Year    int    `jsonapi:"year"`
	RegNum  string `jsonapi:"regNum"`
	OwnerID string `jsonapi:"ownerID"`
}
