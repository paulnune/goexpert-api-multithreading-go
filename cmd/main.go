package main

import (
	"api-multithreading-go/internal/handler"
)

func main() {
	cep := "01153000"
	handler.HandleCEP(cep)
}
