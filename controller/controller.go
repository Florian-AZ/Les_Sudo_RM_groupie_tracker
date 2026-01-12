package controller

import (
	"Groupie_Tracker/api"
	"Groupie_Tracker/data"
	"Groupie_Tracker/structure"
	"Groupie_Tracker/token"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var SessionData = data.InitSessionData()

func Home(w http.ResponseWriter, r *http.Request) {
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
	// Déclaration de la variable qui va contenir les données de la page
	var pagedata structure.PageData_Recherche
	// Récupération du paramètre GET "search" de l'URL
	query := r.URL.Query().Get("search")
	// Si une recherche est fournie dans l'URL (méthode GET)
	if query != "" {
		// Récupération du paramètre GET "page" de l'URL
		pagestr := r.URL.Query().Get("page")
		// Conversion en int avec gestion d'erreur et valeur par défaut (1)
		page := 1
		// Si le paramètre page est fourni et valide
		if pagestr != "" {
			// Conversion en int (Atoi) avec gestion d'erreur
			//Si pas d'erreur et page > 0 on utilise la valeur fournie
			if p, err := strconv.Atoi(pagestr); err == nil && p > 0 {
				page = p
			}
		}
		// Calcul de l'offset pour l'API Spotify (10 résultats par page). Exemple: page 1 -> offset 0, page 2 -> offset 10, page 3 -> offset 20
		offset := (page - 1) * 10

		// Récupération d'un token valide
		token := token.GetValidToken()
		Recherche := api.SearchBar(token, query, offset)
		// Si une erreur est survenue lors de la recherche
		if Recherche.Error.Message != "" {
			fmt.Printf("controller - Recherche - Erreur : %d %s\n\n", Recherche.Error.Status, Recherche.Error.Message)
			pagedata = structure.PageData_Recherche{
				SearchQuery:   query,
				ErreurStatus:  Recherche.Error.Status,
				ErreurMessage: Recherche.Error.Message,
			}
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
					//On envoie le bouton "Suivant" si on a 10 résultats (limite max par requête) et que la page actuelle est inférieure à 100 (limite max de l'API Spotify)
					ASuivant:   (len(htmlData.AlbumData) == 10 || len(htmlData.TrackData) == 10 || len(htmlData.ArtistData) == 10) && page < 100,
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
	// Rendu du template HTML avec les données de la page (avec ou sans recherche ou erreur)
	tmpl := template.Must(template.ParseFiles("template/recherche.html"))
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
