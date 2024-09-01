package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	c := http.Client{Timeout: 10 * time.Second}

	jsonVar := bytes.NewBuffer([]byte(`{"name": "John Doe"}`))

	resp, err := c.Post("http://google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
