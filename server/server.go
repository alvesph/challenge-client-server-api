package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type Cotation struct {
	Usdbrl struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func main() {
	http.HandleFunc("/cotacao", handler)
	log.Println("Server is starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		log.Printf("Error creating request: %v\n", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			http.Error(w, "Request timed out: ", http.StatusRequestTimeout)
			log.Printf(" Request timed out: %v\n", err)
		} else {
			http.Error(w, "Error reading response", http.StatusInternalServerError)
			log.Printf("Error reading response: %v\n", err)
		}
		return
	}

	defer req.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Error from external API", resp.StatusCode)
		log.Printf("Error from external API: %v\n", resp.Status)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
		log.Printf("Error reading response body: %v\n", err)
		return
	}

	var data Cotation
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "Error unmarshalling response body", http.StatusInternalServerError)
		log.Printf("Error unmarshalling response body: %v\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.Usdbrl)
}
