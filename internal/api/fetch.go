package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"api-multithreading-go/dto"
)

func fetchAPI(ctx context.Context, url, provider string) (dto.Address, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return dto.Address{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dto.Address{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return dto.Address{}, errors.New("API " + provider + " retornou status " + resp.Status)
	}

	var address dto.Address
	if err := json.NewDecoder(resp.Body).Decode(&address); err != nil {
		return dto.Address{}, err
	}

	address.Provider = provider
	return address, nil
}
