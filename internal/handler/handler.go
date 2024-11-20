package handler

import (
	"context"
	"fmt"
	"log"

	"api-multithreading-go/configs"
	"api-multithreading-go/internal/api"
)

func HandleCEP(cep string) {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()

	addressCh := make(chan string, 2)
	errorCh := make(chan error, 2)

	go func() {
		addr, err := api.FetchBrasilAPI(ctx, cep)
		if err != nil {
			errorCh <- err
		} else {
			addressCh <- fmt.Sprintf("Resposta BrasilAPI: %+v", addr)
		}
	}()

	go func() {
		addr, err := api.FetchViaCEP(ctx, cep)
		if err != nil {
			errorCh <- err
		} else {
			addressCh <- fmt.Sprintf("Resposta ViaCEP: %+v", addr)
		}
	}()

	select {
	case addr := <-addressCh:
		log.Println(addr)
	case <-ctx.Done():
		log.Println("Timeout: Nenhuma API respondeu a tempo.")
	case err := <-errorCh:
		log.Println("Erro:", err)
	}
}
