package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
)

// AuthMiddleware valida se a requisição está autenticada
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Exemplo de header Authorization: "Basic YWRtaW46YWRtaW4="
		const expectedAuth = "Basic " + base64.StdEncoding.EncodeToString([]byte("admin:admin"))
		if authHeader != expectedAuth {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// handleHello é o handler para a rota principal
func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, authenticated user!")
}

func main() {
	http.Handle("/", AuthMiddleware(http.HandlerFunc(handleHello)))

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
