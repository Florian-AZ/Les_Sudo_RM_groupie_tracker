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
	mux.HandleFunc("/recherche", controller.Recherche)
	mux.HandleFunc("/artiste/", controller.Artiste)
	mux.HandleFunc("/album/", controller.Album)
	mux.HandleFunc("/titre/", controller.Titre)
	mux.HandleFunc("/inscription", controller.Inscription)
	mux.HandleFunc("/connexion", controller.Connexion)
	mux.HandleFunc("/deconnexion", controller.Deconnexion)
	mux.HandleFunc("/favoris", controller.Favoris)
	mux.HandleFunc("/apropos", controller.APropos)

	// Ajout des fichiers statiques
	fileServer := http.FileServer(http.Dir("./assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fileServer))

	return mux
}
