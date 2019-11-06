package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// 145.100.108.82
// ulimit -n 2000
func main() {

	if len(os.Args) < 2 {
		log.Fatal("you need to pass an IP")
	}

	var (
		start       time.Time
		serviceDown bool
	)

	for {

		fmt.Print("CHECKING... ")

		http.DefaultClient = &http.Client{
			Timeout: 100 * time.Millisecond,
		}

		resp, err := http.Get("http://" + os.Args[1])
		if err != nil || resp.StatusCode != http.StatusOK {
			fmt.Println(err)
			if !serviceDown {
				start = time.Now()
				serviceDown = true
				fmt.Println("SERVICE DOWN:", start)
				time.Sleep(50 * time.Millisecond)
			}
			continue
		}

		fmt.Println("SERVICE UP:", resp.Status)

		// check if the service became reachable again
		if serviceDown && resp.StatusCode == http.StatusOK {
			serviceDown = false
			fmt.Println("DOWNTIME:", time.Since(start))
		}

		time.Sleep(1000 * time.Millisecond)
	}

}
