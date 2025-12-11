package router

import (
	"Groupie_Tracker/controller"
	"net/http"
)

// New crée et retourne un nouvel objet ServeMux configuré avec les routes de l'application
func New() *http.ServeMux {
	mux := http.NewServeMux()

	// Routes de ton app
	mux.HandleFunc("/", controller.Home)
	mux.HandleFunc("/album/Damso", controller.Damso)
	mux.HandleFunc("/track/Laylow", controller.Laylow)

	// Ajout des fichiers statiques
	fileServer := http.FileServer(http.Dir("./assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fileServer))

	return mux
}
