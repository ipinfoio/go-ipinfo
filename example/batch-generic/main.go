package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/ipinfo/go/v2/ipinfo/cache"
)

func main() {
	client := ipinfo.NewClient(
		nil,
		ipinfo.NewCache(cache.NewInMemory().WithExpiration(5*time.Minute)),
		os.Getenv("TOKEN"),
	)
	for i := 0; i < 3; i++ {
		fmt.Printf("doing lookup #%v\n", i)
		batchResult, err := client.GetBatch(
			[]string{
				"1.1.1.1",
				"1.1.1.1/country",
				"8.8.8.8",
				"8.8.8.8/country",
				"AS321",
			},
			ipinfo.BatchReqOpts{
				BatchSize:       2,
				TimeoutPerBatch: 0,
				TimeoutTotal:    5,
			},
		)
		if err != nil {
			log.Fatal(err)
		}
		for k, v := range batchResult {
			fmt.Printf("k=%v v=%v\n", k, v)
		}
		fmt.Println()
	}
}
