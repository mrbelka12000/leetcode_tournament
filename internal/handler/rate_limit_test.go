package handler

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestRateLimit(t *testing.T) {
	var wg sync.WaitGroup

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for {
				select {
				case <-ctx.Done():
					wg.Done()
					return
				default:
					resp, err := http.Get("http://localhost:3000/tournament/3")
					if err != nil {
						t.Fatal(err)
					}

					if resp.StatusCode == http.StatusTooManyRequests {
						fmt.Println("RATE limit used")
						time.Sleep(2 * time.Second)
					}
					fmt.Println("ok")
					time.Sleep(10000)
				}
			}
		}()
	}

	wg.Wait()
}
