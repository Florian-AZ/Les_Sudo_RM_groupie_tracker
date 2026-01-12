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

var SessionData = data.InitSessionData()

func Home(w http.ResponseWriter, r *http.Request) {
	// Déclaration de la variable qui va contenir les données de la page
	pagedata := structure.PageData_Accueil{
		LogIn: SessionData.LogIn,
	}

	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, pagedata)
}

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
			htmlData := data.TemplateHTMLSearch(Recherche)
			fmt.Printf("controller - Recherche - Succès struct HTML brut : %v\n\n", htmlData)
			SearchFormatLog(htmlData, query)
			// Remplissage des données de la page de recherche
			pagedata = structure.PageData_Recherche{
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

func Artiste(w http.ResponseWriter, r *http.Request) {
	// Récupération de l'ID de l'artiste depuis l'URL
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
	// Récupération d'un token valide
	token := token.GetValidToken()
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

	var pagestr string
	//Artist Top Tracks

	ArtistTopTracks := api.GetArtistTopTracks(token, artistID)
	// Si une erreur est survenue lors de la récupération des Top Tracks
	if ArtistTopTracks.Error.Message != "" {
		fmt.Printf("controller - Artiste - Top Tracks - Erreur : %d %s\n\n", ArtistTopTracks.Error.Status, ArtistTopTracks.Error.Message)
	}

	//Artist Albums
	// Récupération du paramètre GET "pageAlbums" de l'URL pour la pagination
	pagestr = r.URL.Query().Get("pageAlbums")
	pageAlbums, offset := data.GetPageOffset(pagestr)

	ArtistAlbums := api.GetArtistAlbums(token, artistID, offset)
	// Si une erreur est survenue lors de la récupération des Albums
	if ArtistAlbums.Error.Message != "" {
		fmt.Printf("controller - Artiste - Albums - Erreur : %d %s\n\n", ArtistAlbums.Error.Status, ArtistAlbums.Error.Message)
	}

	// Remplissage des données de la page de l'artiste
	htmldata := data.TemplateHTMLArtist(ArtisteData, ArtistTopTracks, ArtistAlbums)
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

func Erreur(w http.ResponseWriter, r *http.Request, status int, message string) {
	pagedata := structure.PageData_Erreur{
		LogIn:         SessionData.LogIn,
		ErreurStatus:  status,
		ErreurMessage: message,
	}
	tmpl := template.Must(template.ParseFiles("template/erreur.html"))
	tmpl.Execute(w, pagedata)
}

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
