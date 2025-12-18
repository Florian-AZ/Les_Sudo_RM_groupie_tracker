package controller

import (
	"Groupie_Tracker/api"
	"Groupie_Tracker/structure"
	"Groupie_Tracker/token"
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	data := structure.PageData{
		Title:   "Accueil",
		Message: "Bienvenue 🎉",
	}

	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, data)
}

func Damso(w http.ResponseWriter, r *http.Request) {
	// Préparation des données à envoyer au template HTML
	AHTML := structure.AlbumData{}

	// Récupération du token pour toute la session de l'utilisateur
	token := token.GetValidToken()
	// Récupération des albums de l'artiste via l'API Spotify
	A := api.GetAlbum(token, "2UwqpfQtNuhBwviIC0f2ie") //Dasmso Spotify ID: 2UwqpfQtNuhBwviIC0f2ie
	if A.Error != "" {
		fmt.Printf("controller.Damso - Erreur - récupération de l'album : %s %s\n\n", A.Error, A.ErrorDescription)
	} else {
		fmt.Printf("controller.Damso - Succès - Album récupéré brut : %v\n\n", A.AlbumItems)
		for i, item := range A.AlbumItems {
			fmt.Printf("controller.Damso - Succès - Album %d Nom de l'album: %s\nDate de sortie: %s\nNombre de pistes: %d\nURL Spotify: %s\nImage: %s\n\n",
				i, item.Name, item.ReleaseDate, item.TotalTracks, item.URL.Spotify, item.Image[1].URL)
		}

		for _, i := range A.AlbumItems {
			data := structure.Data{
				Image:       i.Image[1].URL,
				Name:        i.Name,
				ReleaseDate: i.ReleaseDate,
				TotalTracks: i.TotalTracks,
				URL:         i.URL.Spotify,
			}
			AHTML.Data = append(AHTML.Data, data)
		}
	}

	data := structure.PageData{
		Title:     "Damso",
		Message:   "Bienvenue sur la page de Damso 🎤",
		AlbumData: AHTML,
	}
	tmpl := template.Must(template.ParseFiles("template/damso.html"))
	tmpl.Execute(w, data)
}

func Laylow(w http.ResponseWriter, r *http.Request) {
	THTML := structure.TrackData{}
	token := token.GetValidToken()
	Tr := api.GetTrack(token, "67Pf31pl0PfjBfUmvYNDCL") //Laylow Track ID: 67Pf31pl0PfjBfUmvYNDCL
	if Tr.Error.Message != "" {
		fmt.Printf("controller.Laylow - Erreur - récupération du track : %d %s\n\n", Tr.Error.Status, Tr.Error.Message)
	} else {
		fmt.Printf("\nTrack récupéré : %s\nAlbum: %s\n", Tr.Name, Tr.Album.Name)
		fmt.Printf("%d Nom de musique: %s\nNom de l'artiste: %s\nNom de l'album: %s\nDate de sortie: %s\nURL Spotify: %s\nImage: %s\n\n",
			0, Tr.Name, Tr.Artists[0].Name, Tr.Album.Name, Tr.Album.ReleaseDate, Tr.Album.URL.Spotify, Tr.Album.Image[1].URL)

		THTML = structure.TrackData{
			TrackName:    Tr.Name,
			AlbumName:    Tr.Album.Name,
			AlbumRelease: Tr.Album.ReleaseDate,
			AlbumURL:     Tr.Album.URL.Spotify,
			AlbumImage:   Tr.Album.Image[1].URL,
			ArtistName:   Tr.Artists[0].Name,
		}
	}

	data := structure.PageData{
		Title:     "Laylow",
		Message:   "Bienvenue sur la page de Laylow 🎤",
		TrackData: THTML,
	}
	tmpl := template.Must(template.ParseFiles("template/laylow.html"))
	tmpl.Execute(w, data)
}
