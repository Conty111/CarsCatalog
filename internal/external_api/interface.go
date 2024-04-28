package external_api

// CarData структура для хранения данных о машине
type CarData struct {
	RegNum string      `json:"regNum"`
	Mark   string      `json:"mark"`
	Model  string      `json:"model"`
	Year   int         `json:"year"`
	Owner  *PeopleData `json:"owner"`
}

// PeopleData структура для хранения данных о владельце
type PeopleData struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

//go:generate go run github.com/vektra/mockery/v3 --name ExternalAPIClient --output ../../test/mocks
type ExternalAPIClient interface {
	GetCarInfo(regNum string) (*CarData, error)
}
