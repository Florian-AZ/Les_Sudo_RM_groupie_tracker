package controller

import (
	"Groupie_Tracker/api"
	"Groupie_Tracker/data"
	"Groupie_Tracker/structure"
	"Groupie_Tracker/token"
	"fmt"
	"html/template"
	"net/http"
)

// Initialisation des données de session
var SessionData = data.InitSessionData()

// Fonction pour afficher la page d'accueil
func Home(w http.ResponseWriter, r *http.Request) {
	// Déclaration de la variable qui va contenir les données de la page
	pagedata := structure.PageData_Accueil{
		LogIn: SessionData.LogIn,
	}

	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, pagedata)
}

// Fonction pour afficher la page de recherche
func Recherche(w http.ResponseWriter, r *http.Request) {
	/* Critère                       | GET                      | POST                              |
	   | ----------------------------| -------------------------| ----------------------------------|
	   | Où sont envoyées les données| Dans l’URL (`?key=value`)| Dans le corps (body) de la requête|
	   | Visible dans l’URL          | ✅ Oui                  | ❌ Non                            |
	   | Taille des données          | Limitée                  | Plus grande                       |
	   | Sécurité                    | ❌ Moins sécurisé        | ✅ Plus adapté                   |
	   | Cache navigateur            | ✅ Oui                   | ❌ Non                           |
	   | Usage principal             | Lecture / recherche      | Création / envoi de données       |
	*/
	var pagedata structure.PageData_Recherche
	// Récupération du paramètre GET "search" de l'URL
	query := r.URL.Query().Get("search")
	// Si une recherche est fournie dans l'URL (méthode GET)
	if query != "" {
		// Récupération du paramètre GET "page" de l'URL
		pagestr := r.URL.Query().Get("page")
		page, offset := data.GetPageOffset(pagestr)
		fmt.Printf("controller - Recherche - Recherche fournie : %s | Page : %d | Offset : %d\n\n", query, page, offset)

		// Récupération d'un token valide
		token := token.GetValidToken()
		if token == "" {
			fmt.Printf("controller - Recherche - Erreur : Token invalide\n\n")
			pagedata := data.TemplateErreur(500, "APP")
			Erreur(w, r, pagedata.Status, pagedata.Message)
			return
		}
		Recherche := api.SearchBar(token, query, offset)
		// Si une erreur est survenue lors de la recherche, on affiche une page d'erreur
		if Recherche.Error.Message != "" {
			fmt.Printf("controller - Recherche - Erreur : %d %s\n\n", Recherche.Error.Status, Recherche.Error.Message)
			pagedata := data.TemplateErreur(Recherche.Error.Status, "API")
			Erreur(w, r, pagedata.Status, pagedata.Message)
			return
			// Si la recherche est un succès
		} else {
			// Restructuration des données brutes en données prêtes pour le template HTML
			htmlData := data.TemplateHTMLSearch(Recherche, SessionData)
			fmt.Printf("controller - Recherche - Succès struct HTML brut : %v\n\n", htmlData)
			SearchFormatLog(htmlData, query)
			// Remplissage des données de la page de recherche
			pagedata = structure.PageData_Recherche{
				LogIn:       SessionData.LogIn,
				SearchData:  htmlData,
				SearchQuery: query,
				Pagination: structure.Pagination{
					Page: page,
					//On envoie le bouton "Suivant" si on a plus de 10 résultats (limite max par requête) et que la page actuelle est inférieure à 100 (limite max de l'API Spotify)
					ASuivant:   (len(htmlData.AlbumData) > 10 || len(htmlData.TrackData) >= 10 || len(htmlData.ArtistData) >= 10) && page < 100,
					APrecedent: page > 1,
					PageSuiv:   page + 1,
					PagePrec:   page - 1,
				},
			}
		}
		// Si aucune recherche n'est fournie on affiche une page de recherche vide
	} else {
		pagedata = structure.PageData_Recherche{
			LogIn: SessionData.LogIn,
		}
	}
	// Rendu du template HTML avec les données de la page (avec ou sans recherche)
	tmpl := template.Must(template.ParseFiles("template/recherche.html"))
	tmpl.Execute(w, pagedata)
}

// Fonction pour afficher la page d'un artiste
func Artiste(w http.ResponseWriter, r *http.Request) {
	// Récupération de l'ID de l'artiste depuis l'URL.
	artistID := data.GetIdFromUrl(r)
	// Si aucun ID n'est fourni
	if artistID == "" {
		// Rdiriger vers la page not found
		fmt.Printf("controller - Artiste - Aucun ID d'artiste fourni dans l'URL\n\n")
		// Préparation des données d'erreur
		pagedata := data.TemplateErreur(400, "APP")
		Erreur(w, r, pagedata.Status, pagedata.Message)
		return
	}
	if r.Method == http.MethodPost {
		// Gestion des ajouts/suppressions de favoris
		ajoutFavorisID := r.FormValue("ajout_favoris")
		retirerFavorisID := r.FormValue("retirer_favoris")
		if ajoutFavorisID != "" {
			if !data.IsFavoris(ajoutFavorisID, "artiste", SessionData) {
				data.AjoutFavoris(ajoutFavorisID, "artiste", SessionData)
				fmt.Printf("controller - Artiste - Ajout aux favoris : ID %s\n\n", ajoutFavorisID)
			} else {
				fmt.Printf("controller - Artiste - L'artiste ID %s est déjà dans les favoris\n\n", ajoutFavorisID)
			}
		}
		if retirerFavorisID != "" {
			if data.IsFavoris(retirerFavorisID, "artiste", SessionData) {
				data.RetirerFavoris(retirerFavorisID, "artiste", SessionData)
				fmt.Printf("controller - Artiste - Retrait des favoris : ID %s\n\n", retirerFavorisID)
			} else {
				fmt.Printf("controller - Artiste - L'artiste ID %s n'est pas dans les favoris\n\n", retirerFavorisID)
			}
		}
	}
	// Récupération d'un token valide
	token := token.GetValidToken()
	if token == "" {
		fmt.Printf("controller - Artiste - Erreur : Token invalide\n\n")
		pagedata := data.TemplateErreur(500, "APP")
		Erreur(w, r, pagedata.Status, pagedata.Message)
		return
	}
	ArtisteData := api.GetArtistData(token, artistID, 0)
	// Si une erreur est survenue lors de la récupération des données de l'artiste
	if ArtisteData.Error.Message != "" {
		fmt.Printf("controller - Artiste - Erreur : %d %s\n\n", ArtisteData.Error.Status, ArtisteData.Error.Message)
		pagedata := data.TemplateErreur(ArtisteData.Error.Status, "API")
		Erreur(w, r, pagedata.Status, pagedata.Message)
		return
	}
	// Si la récupération des données de l'artiste est un succès
	fmt.Printf("controller - Artiste - Succès struct brut : %v\n\n", ArtisteData)

	//Artist Top Tracks

	ArtistTopTracks := api.GetArtistTopTracks(token, artistID)
	// Si une erreur est survenue lors de la récupération des Top Tracks
	if ArtistTopTracks.Error.Message != "" {
		fmt.Printf("controller - Artiste - Top Tracks - Erreur : %d %s\n\n", ArtistTopTracks.Error.Status, ArtistTopTracks.Error.Message)
	}

	//Artist Albums
	// Récupération du paramètre GET "pageAlbums" de l'URL pour la pagination
	pagestr := r.URL.Query().Get("pageAlbums")
	pageAlbums, offset := data.GetPageOffset(pagestr)

	ArtistAlbums := api.GetArtistAlbums(token, artistID, offset)
	// Si une erreur est survenue lors de la récupération des Albums
	if ArtistAlbums.Error.Message != "" {
		fmt.Printf("controller - Artiste - Albums - Erreur : %d %s\n\n", ArtistAlbums.Error.Status, ArtistAlbums.Error.Message)
	}

	// Remplissage des données de la page de l'artiste
	htmldata := data.TemplateHTMLArtist(ArtisteData, ArtistTopTracks, ArtistAlbums, SessionData)
	pagedata := structure.PageData_Artiste{
		LogIn:      SessionData.LogIn,
		ArtistData: htmldata,
		PaginationAlbums: structure.Pagination{
			Page: pageAlbums,
			//On envoie le bouton "Suivant" si on a plus de 10 résultats (limite max par requête) et que la page actuelle est inférieure à 100 (limite max de l'API Spotify)
			ASuivant:   (len(ArtistAlbums.Items) > 10) && pageAlbums < 100,
			APrecedent: pageAlbums > 1,
			PageSuiv:   pageAlbums + 1,
			PagePrec:   pageAlbums - 1,
		},
		ErrTotalTracks: data.TemplateErreur(ArtistTopTracks.Error.Status, "API"),
		ErrAlbums:      data.TemplateErreur(ArtistAlbums.Error.Status, "API"),
	}
	tmpl := template.Must(template.ParseFiles("template/artiste.html"))
	tmpl.Execute(w, pagedata)
}

// Fonction pour afficher la page d'un album
func Album(w http.ResponseWriter, r *http.Request) {
	albumID := data.GetIdFromUrl(r)
	if albumID == "" {
		fmt.Printf("controller - Album - Aucun ID d'album fourni dans l'URL\n\n")
		// Préparation des données d'erreur
		pagedata := data.TemplateErreur(400, "APP")
		Erreur(w, r, pagedata.Status, pagedata.Message)
		return
	}
	if r.Method == http.MethodPost {
		// Gestion des ajouts/suppressions de favoris
		ajoutFavorisID := r.FormValue("ajout_favoris")
		retirerFavorisID := r.FormValue("retirer_favoris")
		if ajoutFavorisID != "" {
			if !data.IsFavoris(albumID, "album", SessionData) {
				data.AjoutFavoris(ajoutFavorisID, "album", SessionData)
				fmt.Printf("controller - Album - Ajout aux favoris : ID %s\n\n", ajoutFavorisID)
			} else {
				fmt.Printf("controller - Album - L'album ID %s est déjà dans les favoris\n\n", albumID)
			}
		}
		if retirerFavorisID != "" {
			if data.IsFavoris(albumID, "album", SessionData) {
				data.RetirerFavoris(retirerFavorisID, "album", SessionData)
				fmt.Printf("controller - Album - Retrait des favoris : ID %s\n\n", retirerFavorisID)
			} else {
				fmt.Printf("controller - Album - L'album ID %s n'est pas dans les favoris\n\n", albumID)
			}
		}
	}
	token := token.GetValidToken()
	if token == "" {
		fmt.Printf("controller - Album - Erreur : Token invalide\n\n")
		pagedata := data.TemplateErreur(500, "APP")
		Erreur(w, r, pagedata.Status, pagedata.Message)
		return
	}
	pagestr := r.URL.Query().Get("page")
	page, offset := data.GetPageOffset(pagestr)
	fmt.Printf("controller - Album - Page : %d | Offset : %d\n\n", page, offset)

	AlbumTracks := api.GetAlbum(token, albumID, offset)
	if AlbumTracks.Error.Message != "" {
		fmt.Printf("controller - Album - Erreur : %d %s\n\n", AlbumTracks.Error.Status, AlbumTracks.Error.Message)
		pagedata := data.TemplateErreur(AlbumTracks.Error.Status, "API")
		Erreur(w, r, pagedata.Status, pagedata.Message)
		return
	}
	fmt.Printf("controller - Album - Succès struct brut : %v\n\n", AlbumTracks)
	htmldata := data.TemplateHTMLAlbums(AlbumTracks, SessionData)
	pagedata := structure.PageData_Album{
		LogIn:       SessionData.LogIn,
		AlbumTracks: htmldata,
		Pagination: structure.Pagination{
			Page: page,
			//On envoie le bouton "Suivant" si on a plus de 10 résultats (limite max par requête) et que la page actuelle est inférieure à 100 (limite max de l'API Spotify)
			ASuivant:   (len(AlbumTracks.Tracks.Items) > 10) && page < 100,
			APrecedent: page > 1,
			PageSuiv:   page + 1,
			PagePrec:   page - 1,
		},
	}
	tmpl := template.Must(template.ParseFiles("template/album.html"))
	tmpl.Execute(w, pagedata)
}

// Fonction pour afficher la page d'un titre
func Titre(w http.ResponseWriter, r *http.Request) {
	trackID := data.GetIdFromUrl(r)
	if trackID == "" {
		fmt.Printf("controller - Titre - Aucun ID de titre fourni dans l'URL\n\n")
		// Préparation des données d'erreur
		pagedata := data.TemplateErreur(400, "APP")
		Erreur(w, r, pagedata.Status, pagedata.Message)
		return
	}

	if r.Method == http.MethodPost {
		// Gestion des ajouts/suppressions de favoris
		ajoutFavorisID := r.FormValue("ajout_favoris")
		retirerFavorisID := r.FormValue("retirer_favoris")
		if ajoutFavorisID != "" {
			if !data.IsFavoris(trackID, "titre", SessionData) {
				data.AjoutFavoris(ajoutFavorisID, "titre", SessionData)
				fmt.Printf("controller - Titre - Ajout aux favoris : ID %s\n\n", ajoutFavorisID)
			} else {
				fmt.Printf("controller - Titre - Le titre ID %s est déjà dans les favoris\n\n", trackID)
			}
		}
		if retirerFavorisID != "" {
			if data.IsFavoris(trackID, "titre", SessionData) {
				data.RetirerFavoris(retirerFavorisID, "titre", SessionData)
				fmt.Printf("controller - Titre - Retrait des favoris : ID %s\n\n", retirerFavorisID)
			} else {
				fmt.Printf("controller - Titre - Le titre ID %s n'est pas dans les favoris\n\n", trackID)
			}
		}
	}

	token := token.GetValidToken()
	if token == "" {
		fmt.Printf("controller - Titre - Erreur : Token invalide\n\n")
		pagedata := data.TemplateErreur(500, "APP")
		Erreur(w, r, pagedata.Status, pagedata.Message)
		return
	}

	track := api.GetTrack(token, trackID)
	if track.Error.Message != "" {
		fmt.Printf("controller - Titre - Erreur : %d %s\n\n", track.Error.Status, track.Error.Message)
		pagedata := data.TemplateErreur(track.Error.Status, "API")
		Erreur(w, r, pagedata.Status, pagedata.Message)
		return
	}
	fmt.Printf("controller - Titre - Succès struct brut : %v\n\n", track)
	htmldata := data.TemplateHTMLTrack(track, SessionData)
	pagedata := structure.PageData_Titre{
		LogIn:     SessionData.LogIn,
		TrackData: htmldata,
	}
	tmpl := template.Must(template.ParseFiles("template/titre.html"))
	tmpl.Execute(w, pagedata)
}

// Fonction pour afficher la page d'inscription
func Inscription(w http.ResponseWriter, r *http.Request) {
	if SessionData.LogIn {
		// Redirige vers la page d'accueil si l'utilisateur est déjà connecté
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	var pagedata structure.PageData_Connexion_Inscription
	if r.Method != http.MethodPost {
		pagedata = structure.PageData_Connexion_Inscription{
			Erreur: "",
		}
	} else {
		// Récupération des données du formulaire
		user := r.FormValue("utilisateur")
		mdp := r.FormValue("MdP")
		mdpConf := r.FormValue("MdPConf")
		if mdp != mdpConf {
			pagedata = structure.PageData_Connexion_Inscription{
				Erreur: "Les mots de passe ne correspondent pas.",
			}
		} else {
			err := data.CreationCompte(user, mdp)
			if err != "" {
				pagedata = structure.PageData_Connexion_Inscription{
					Erreur: err,
				}
			} else {
				fmt.Printf("controller - Inscription - Utilisateur %s inscrit avec succès\n\n", user)
				// Redirige vers la page de connexion après une inscription réussie (code 303(requête GET))
				http.Redirect(w, r, "/connexion", http.StatusSeeOther)
				return
			}
		}
	}

	tmpl := template.Must(template.ParseFiles("template/inscription.html"))
	tmpl.Execute(w, pagedata)
}

// Fonction pour afficher la page d'connexion
func Connexion(w http.ResponseWriter, r *http.Request) {
	if SessionData.LogIn {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	var pagedata structure.PageData_Connexion_Inscription
	if r.Method != http.MethodPost {
		pagedata = structure.PageData_Connexion_Inscription{
			Erreur: "",
		}
	} else {
		// Récupération des données du formulaire
		user := r.FormValue("utilisateur")
		mdp := r.FormValue("MdP")
		err := data.ConnexionCompte(user, mdp, SessionData)
		if err != "" {
			pagedata = structure.PageData_Connexion_Inscription{
				Erreur: "Erreur lors de la connexion : " + err,
			}
		} else {
			fmt.Printf("controller - Connexion - Utilisateur %s connecté avec succès\n\n", user)
			// Redirige vers la page d'accueil après une connexion réussie (code 303(requête GET))
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

	}

	tmpl := template.Must(template.ParseFiles("template/connexion.html"))
	tmpl.Execute(w, pagedata)
}

// Fonction pour déconnecter l'utilisateur
func Deconnexion(w http.ResponseWriter, r *http.Request) {
	if !SessionData.LogIn {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodPost {
		if r.FormValue("deconnexion") == "true" {
			// Réinitialisation des données de session
			SessionData.LogIn = false
			SessionData.Utilisateur = structure.Utilisateur{
				Nom: "",
				Favoris: structure.Utilisateur_Favoris{
					IdTitres:   []structure.Favoris_Id{},
					IdArtistes: []structure.Favoris_Id{},
					IdAlbums:   []structure.Favoris_Id{},
				},
			}
			fmt.Printf("controller - Deconnexion - Utilisateur %s déconnecté avec succès\n\n", SessionData.Utilisateur.Nom)
			// Redirige vers la page d'accueil après la déconnexion (code 303(requête GET))
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
	tmpl := template.Must(template.ParseFiles("template/deconnexion.html"))
	tmpl.Execute(w, nil)
}

func Favoris(w http.ResponseWriter, r *http.Request) {
	if !SessionData.LogIn {
		// Redirige vers la page d'accueil si l'utilisateur n'est pas connecté
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	token := token.GetValidToken()
	if token == "" {
		fmt.Printf("controller - Album - Erreur : Token invalide\n\n")
		pagedata := data.TemplateErreur(500, "APP")
		Erreur(w, r, pagedata.Status, pagedata.Message)
		return
	}

	pagestr := r.URL.Query().Get("page")
	page, offset := data.GetPageOffset(pagestr)
	fmt.Printf("controller - Album - Page : %d | Offset : %d\n\n", page, offset)

	htmlData := data.TemplateHTMLFavoris(SessionData.Utilisateur, token, offset)
	pagedata := structure.PageData_Favoris{
		NomUtilisateur: SessionData.Utilisateur.Nom,
		Favoris:        htmlData,
		Pagination: structure.Pagination{
			Page:       page,
			ASuivant:   (len(htmlData.Titres) > 10 || len(htmlData.Artistes) >= 10 || len(htmlData.Albums) >= 10),
			APrecedent: page > 1,
			PageSuiv:   page + 1,
			PagePrec:   page - 1,
		},
	}
	tmpl := template.Must(template.ParseFiles("template/favoris.html"))
	tmpl.Execute(w, pagedata)
}

// Fonction pour afficher une page d'erreur
func Erreur(w http.ResponseWriter, r *http.Request, status int, message string) {
	pagedata := structure.PageData_Erreur{
		LogIn:         SessionData.LogIn,
		ErreurStatus:  status,
		ErreurMessage: message,
	}
	tmpl := template.Must(template.ParseFiles("template/erreur.html"))
	tmpl.Execute(w, pagedata)
}

// Fonction de formatage et d'affichage des résultats de recherche dans la console
func SearchFormatLog(S structure.Html_Recherche, query string) {
	fmt.Printf("///////////////////////////////////////////\n")
	fmt.Printf("Recherche : %s\n\n", query)
	fmt.Printf("Tracks :\n")
	for i, item := range S.TrackData {
		fmt.Printf("Tracks %d:\n", i+1)
		fmt.Printf("Main - SearchBar - \nAlbum URL Spotify: %s\nAlbum id: %s\nAlbum Name: %s\nRelease Date: %s\nTotal Tracks : %d\n\n", item.AlbumURL, item.AlbumId, item.AlbumName, item.ReleaseDate, item.TotalTracks)
		for j, Art := range item.Artists {
			fmt.Printf("Artist %d URL Spotify: %s\nArtist Name: %s\nArtist ID: %s\n", j+1, Art.ArtistURL, Art.ArtistName, Art.ArtistId)
		}
		fmt.Printf("\nTrack Name: %s\nDuration (ms): %d\nDuration (mm:ss): %s\nURL Spotify: %s\nID: %s\nImage (300*300px) URL: %s\n\n", item.TrackName, item.DurationMs, item.DurationFormated, item.TrackURL, item.TrackId, item.Images)
		fmt.Printf("---------------------------------------\n")
	}

	fmt.Printf("///////////////////////////////////////////\n")
	fmt.Printf("\n\nArtists :\n")
	for i, items := range S.ArtistData {
		fmt.Printf("Main - SearchBar - Artist %d\nURL Spotify: %s\nNb Followers : %d\n\n", i+1, items.ArtistURL, items.NbFollowers)
		for j, genres := range items.Genres {
			fmt.Printf("Genre %d: %s\n", j+1, genres)
		}

		fmt.Printf("\nArtist ID: %s\nArtist Name: %s\nImage (300*300px) URL: %s\n\n", items.ArtistId, items.ArtistName, items.Images)
		fmt.Printf("---------------------------------------\n")
	}

	fmt.Printf("///////////////////////////////////////////\n")
	fmt.Printf("\n\nAlbums :\n")
	for i, itema := range S.AlbumData {
		fmt.Printf("Albums %d:\n", i+1)
		fmt.Printf("Main - SearchBar -\nTotal Tracks: %d\nAlbum URL Spotify: %s\nAlbum id: %s\nAlbum Name: %s\nRelease Date: %s\n\n", itema.TotalTracks, itema.AlbumURL, itema.AlbumId, itema.AlbumName, itema.ReleaseDate)
		for k, Art := range itema.Artists {
			fmt.Printf("Artist %d Artist ID: %s\nURL Spotify: %s\nArtist Name: %s\n",
				k+1, Art.ArtistURL, Art.ArtistId, Art.ArtistName)
		}
		fmt.Printf("Image (300*300px) URL: %s\n", itema.Images)
		fmt.Printf("---------------------------------------\n")
	}
	fmt.Printf("///////////////////////////////////////////\n")
}
