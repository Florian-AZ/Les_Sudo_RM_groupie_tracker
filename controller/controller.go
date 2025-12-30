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

var WebData *structure.PageData = data.InitWebData()

func Home(w http.ResponseWriter, r *http.Request) {
	data := structure.PageData{
		LogIn: WebData.LogIn,
	}

	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, data)
}

func Recherche(w http.ResponseWriter, r *http.Request) {
	var pagedata structure.PageData
	if r.Method == http.MethodPost {
		token := token.GetValidToken()
		Recherche := api.SearchBar(token, r.FormValue("recherche"))
		if Recherche.Error.Message != "" {
			fmt.Printf("controller - Recherche - Erreur : %d %s\n\n", Recherche.Error.Status, Recherche.Error.Message)
			pagedata = structure.PageData{
				ErreurStatus:  Recherche.Error.Status,
				ErreurMessage: Recherche.Error.Message,
			}
		} else {
			fmt.Printf("controller - Recherche - Succès brut : %v\n\n", Recherche)
			htmlData := data.TemplateHTMLSearch(Recherche)
			fmt.Printf("controller - Recherche - Succès struct HTML brut : %v\n\n", htmlData)

			pagedata = structure.PageData{
				SearchData: htmlData,
			}
		}
		tmpl := template.Must(template.ParseFiles("template/recherche.html"))
		tmpl.Execute(w, pagedata)
		return
	}

	pagedata = structure.PageData{
		LogIn: WebData.LogIn,
	}
	tmpl := template.Must(template.ParseFiles("template/recherche.html"))
	tmpl.Execute(w, pagedata)
}
func Damso(w http.ResponseWriter, r *http.Request) {
	// Préparation des données à envoyer au template HTML
	html_a := structure.Html_Album{}

	// Récupération du token pour toute la session de l'utilisateur
	token := token.GetValidToken()
	// Récupération des albums de l'artiste via l'API Spotify
	A := api.GetAlbum(token, "2UwqpfQtNuhBwviIC0f2ie") //Dasmso Spotify ID: 2UwqpfQtNuhBwviIC0f2ie
	if A.Error.Message != "" {
		fmt.Printf("controller.Damso - Erreur - récupération de l'album : %d %s\n\n", A.Error.Status, A.Error.Message)
	} else {
		fmt.Printf("controller.Damso - Succès - Album récupéré brut : %v\n\n", A.Items)
		for i, item := range A.Items {
			fmt.Printf("controller.Damso - Succès - Album %d Nom de l'album: %s\nDate de sortie: %s\nNombre de pistes: %d\nURL Spotify: %s\nImage: %s\n\n",
				i, item.Name, item.ReleaseDate, item.TotalTracks, item.URL.Spotify, item.Images[1].URL)
		}

		for _, i := range A.Items {
			data := structure.AlbumData{
				Image:       i.Images[1].URL,
				Name:        i.Name,
				ReleaseDate: i.ReleaseDate,
				TotalTracks: i.TotalTracks,
				URL:         i.URL.Spotify,
			}
			html_a.Data = append(html_a.Data, data)
		}
	}

	data := structure.PageData{
		AlbumData: html_a,
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
			0, Tr.Name, Tr.Artists[0].Name, Tr.Album.Name, Tr.Album.ReleaseDate, Tr.Album.URL.Spotify, Tr.Album.Images[1].URL)

		THTML = structure.TrackData{
			TrackName:    Tr.Name,
			AlbumName:    Tr.Album.Name,
			AlbumRelease: Tr.Album.ReleaseDate,
			AlbumURL:     Tr.Album.URL.Spotify,
			AlbumImage:   Tr.Album.Images[1].URL,
			ArtistName:   Tr.Artists[0].Name,
		}
	}

	data := structure.PageData{
		TrackData: THTML,
	}
	tmpl := template.Must(template.ParseFiles("template/laylow.html"))
	tmpl.Execute(w, data)
}
