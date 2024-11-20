package api

import (
	"context"

	"api-multithreading-go/dto"
)

const BrasilAPIURL = "https://brasilapi.com.br/api/cep/v1/"

func FetchBrasilAPI(ctx context.Context, cep string) (dto.Address, error) {
	url := BrasilAPIURL + cep
	return fetchAPI(ctx, url, "BrasilAPI")
}
