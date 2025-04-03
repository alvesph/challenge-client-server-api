package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Cotation struct {
	Bid string `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Printf("Error when creating the request %v\n", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Printf("Request client timed out: %v\n", err)
		} else {
			log.Printf("Error reading response: %v\n", err)
		}
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error: %v\n", resp.Status)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v\n", err)
		return
	}

	var data Cotation
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("Error unmarshalling response body: %v\n", err)
		return
	}

	content := "Dólar: " + data.Bid

	err = os.WriteFile("cotacao.txt", []byte(content), 0644)
	if err != nil {
		log.Printf("Error writing response: %v\n", err)
		return
	}

	log.Println("Quotation saved successfully in the file cotacao.txt!")
}
