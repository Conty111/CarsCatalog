package car

type MsgResponse struct {
	Status  string `jsonapi:"status"`
	Message string `jsonapi:"message"`
}
