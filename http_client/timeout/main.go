package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	c := http.Client{Timeout: 10 * time.Second}

	resp, err := c.Get("http://google.com")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
