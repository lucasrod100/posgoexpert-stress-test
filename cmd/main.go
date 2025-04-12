package main

import (
	"flag"
	"fmt"

	"github.com/lucasrod100/posgoexpert-stress-test/internal"
)

func main() {
	url := flag.String("url", "", "URL do serviço a ser testado")
	requests := flag.Int("requests", 0, "Número total de requests")
	concurrency := flag.Int("concurrency", 0, "Número de chamadas simultâneas")
	flag.Parse()

	if *url == "" || *requests == 0 || *concurrency == 0 {
		fmt.Println("Erro: URL, número de requests e concorrência devem ser informados.")
		flag.Usage()
		return
	}

	stresstest := internal.NewStressTest(*url, *requests, *concurrency)
	stresstest.Run()
}
