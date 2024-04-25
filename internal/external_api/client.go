package external_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// CarData структура для хранения данных о машине
type CarData struct {
	RegNum string     `json:"regNum"`
	Mark   string     `json:"mark"`
	Model  string     `json:"model"`
	Year   int        `json:"year"`
	Owner  PeopleData `json:"owner"`
}

// PeopleData структура для хранения данных о владельце
type PeopleData struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type APIClient struct {
	Address string
}

// GetCarInfo возвращает информацию о машине по её регистрационному номеру
func (a *APIClient) GetCarInfo(regNum string) (*CarData, error) {
	url := fmt.Sprintf("%s/info?regNum=%s", a.Address, regNum)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	var carData CarData
	if err := json.Unmarshal(body, &carData); err != nil {
		return nil, err
	}

	return &carData, nil
}
