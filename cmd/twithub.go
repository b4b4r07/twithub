package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

func main() {
	user := ""
	if len(os.Args) > 1 {
		user = os.Args[1]
	} else {
		os.Exit(1)
	}

	url := [2]string{"https://github.com/" + user, "https://twitter.com/" + user}
	var result map[string]int = make(map[string]int)

	var wg sync.WaitGroup
	for i := 0; i < len(url); i++ {
		wg.Add(1)
		go func(i int) {
			resp, err := http.Get(url[i])
			if err != nil {
				log.Fatal(err)
			}
			result[url[i]] = resp.StatusCode
			defer resp.Body.Close()
			wg.Done()
		}(i)
	}
	wg.Wait()

	var available bool = true
	for _, v := range result {
		if v == 200 {
			available = false
		}
	}

	fmt.Println(result)
	if available {
		fmt.Println("ok")
	} else {
		fmt.Println("ng")
	}
}
