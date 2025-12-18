package controller

import (
	"Groupie_Tracker/api"
	"Groupie_Tracker/structure"
	"fmt"
	"html/template"
	"net/http"
)

var Token *string

func Home(w http.ResponseWriter, r *http.Request) {
	data := structure.PageData{
		Title:   "Accueil",
		Message: "Bienvenue 🎉",
	}
	// Récupération du token pour toute la session de l'utilisateur
	T := api.GetToken()
	if T.Error != "" {
		fmt.Println("Erreur lors de la récupération du token : ", T.Error, " ", T.ErrorDescription)
	} else {
		Token = &T.AccessToken
		fmt.Println("Token récupéré : ", *Token)
	}

	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, data)
}

func Damso(w http.ResponseWriter, r *http.Request) {
	AHTML := structure.AlbumData{}
	A := api.GetAlbum(*Token, "2UwqpfQtNuhBwviIC0f2ie") //Dasmso Spotify ID: 2UwqpfQtNuhBwviIC0f2ie
	if A.Error != "" {
		fmt.Println("Erreur lors de la récupération de l'album : ", A.Error, " ", A.ErrorDescription)
	} else {
		fmt.Println("\nAlbum récupéré : ", A.AlbumItems)
		for i, item := range A.AlbumItems {
			fmt.Printf("%d Nom de l'album: %s\nDate de sortie: %s\nNombre de pistes: %d\nURL Spotify: %s\nImage: %s\n\n",
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
	Tr := api.GetTrack(*Token, "67Pf31pl0PfjBfUmvYNDCL") //Laylow Track ID: 67Pf31pl0PfjBfUmvYNDCL
	if Tr.Error.Message != "" {
		fmt.Println("Erreur lors de la récupération du track : ", Tr.Error.Status, " ", Tr.Error.Message)
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
