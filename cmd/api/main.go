package main

import (
	"fmt"
	"github.com/8soat-grupo35/fastfood-order/external"

	"github.com/8soat-grupo35/fastfood-order/internal/api/server"
)

func main() {
	fmt.Println("Iniciado o servidor Rest com GO")
	cfg := external.GetConfig()
	server.Start(cfg)
}
