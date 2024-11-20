package api

import (
	"api-multithreading-go/dto"
	"context"
)

const ViaCEPURL = "http://viacep.com.br/ws/"

func FetchViaCEP(ctx context.Context, cep string) (dto.Address, error) {
	url := ViaCEPURL + cep + "/json/"
	return fetchAPI(ctx, url, "ViaCEP")
}
