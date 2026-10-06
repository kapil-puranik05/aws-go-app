package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Greeting struct {
	Message string `json:"greeting"`
}

func Greet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	response := Greeting{
		Message: "Hello",
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/greet", Greet)
	server := &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}
	go func() {
		log.Println("Server started on port 3000")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to shut down the server gracefully: %v", err)
	}
	log.Println("Server shutdown successful")
}
