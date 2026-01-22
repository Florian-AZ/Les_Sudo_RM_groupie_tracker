package main

import (
	"Groupie_Tracker/router"
	"fmt"
	"net/http"
)

func main() {
	mux := router.New()
	fmt.Printf("main - succés - 🚀 Serveur démarré sur http://localhost:8080\n\n")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Printf("main - Erreur - serveur : %s\n\n", err)
	}
}
