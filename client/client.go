package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Cotation struct {
	Bid string `json:"bid"`
}

func main() {
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatalf("Error when making the request %v\n", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: %v\n", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v\n", err)
	}

	var data Cotation
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalf("Error unmarshalling response body: %v\n", err)
	}

	fmt.Println("Dólar: ", data.Bid)
}
