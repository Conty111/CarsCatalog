package external_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Conty111/CarsCatalog/internal/configs"
	"github.com/rs/zerolog/log"
	"net/http"
)

type Client struct {
	defaultHeader map[string]string
	client        *http.Client
	url           string
}

func NewClient(cfg *configs.Configuration) *Client {
	var apiClient Client

	apiClient.client = &http.Client{Timeout: cfg.APIClient.TimeoutResponse}
	apiClient.defaultHeader = cfg.APIClient.DefaultHeader
	apiClient.url = fmt.Sprintf("%s://%s:%s",
		cfg.APIClient.Scheme, cfg.APIClient.ServerAddress, cfg.APIClient.ServerPort)

	return &apiClient
}

func (c *Client) GetCarInfo(regNum string) (*CarData, error) {
	req, err := http.NewRequest(http.MethodGet, c.url+"/info?regNum="+regNum, nil)
	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to create request")
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	for key, value := range c.defaultHeader {
		req.Header.Set(key, value)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to send request")
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error().
			Int("statusCode", resp.StatusCode).
			Str("status", http.StatusText(resp.StatusCode)).
			Msg("failed to get car info")
		return nil, errors.New("failed to get car info")
	}

	var car CarData
	err = json.NewDecoder(resp.Body).Decode(&car)
	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to get car info")
		return nil, fmt.Errorf("failed to decode body: %w", err)
	}

	return &car, nil
}
