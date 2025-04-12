package internal

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/fatih/color"
)

type StressTest struct {
	url           string
	totalRequests int
	concurrency   int
}

func NewStressTest(url string, totalRequests int, concurrency int) *StressTest {
	return &StressTest{
		url:           url,
		totalRequests: totalRequests,
		concurrency:   concurrency,
	}
}

func (st *StressTest) Run() {
	responses := make(chan int, st.totalRequests)
	var wg sync.WaitGroup
	startTime := time.Now()

	tasksPerWorker := st.totalRequests / st.concurrency
	remaining := st.totalRequests % st.concurrency

	for i := 0; i < st.concurrency; i++ {
		wg.Add(1)
		requests := tasksPerWorker
		if i < remaining {
			requests++
		}
		go worker(st.url, requests, responses, &wg)
	}

	wg.Wait()
	close(responses)
	duration := time.Since(startTime)
	totalSuccess := 0
	statusCounts := make(map[int]int)

	for result := range responses {
		if result == 200 {
			totalSuccess++
		}
		statusCounts[result]++
	}

	color.Magenta("\nRelatório Final:")
	color.Green("Tempo total: %v", duration)
	color.Blue("Total de requests: %d", st.totalRequests)
	color.Green("Requests com status 200: %d", totalSuccess)
	color.Yellow("Distribuição dos códigos de status:")
	for status, count := range statusCounts {
		if status == 200 {
			color.Green("Status %d: %d", status, count)
		} else {
			color.Red("Status %d: %d", status, count)
		}
	}
}

func worker(url string, requests int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < requests; i++ {
		resp, err := http.Get(url)

		if err != nil {
			fmt.Printf("Erro na request: %v\n", err)
			results <- 0
			continue
		}
		results <- resp.StatusCode
		resp.Body.Close()
	}
}
