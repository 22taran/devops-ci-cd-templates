package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func greet(name string) string {
	if strings.TrimSpace(name) == "" {
		return "Hello, World!"
	}
	return fmt.Sprintf("Hello, %s!", name)
}

func add(a, b int) int {
	return a + b
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": greet("World")})
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
