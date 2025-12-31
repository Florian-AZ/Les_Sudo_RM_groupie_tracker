package data

import (
	"Groupie_Tracker/structure"
	"fmt"
	"time"
)

// InitSessionData initialise les données de la page web
func InitSessionData() *structure.SessionData {
	fmt.Printf("data.InitSessionData - Initialisation des données de la session\n\n")
	// Initialisation des données de la session
	return &structure.SessionData{
		Utilisateur: "",
		LogIn:       false,
	}
}

//Restructuration des données obtenues via l'API Spotify en une structure plus lisible pour le template HTML

// TemplateHTMLSearch → Transforme la réponse brute de l'API (Api_Recherche) en une structure plus lisible et prête pour le template HTML (Html_Recherche).
func TemplateHTMLSearch(S structure.Api_Recherche) structure.Html_Recherche {
	// Remplissage des données pour le template HTML
	var html_S structure.Html_Recherche
	// Boucle pour les tracks
	/*TrackData[]
	AlbumURL    string
	AlbumId     string
	AlbumName   string
	ReleaseDate string
	TotalTracks int
	Artists[]
		ArtistURL  string
		ArtistId   string
		ArtistName string
	TrackName   string
	DurationMs  int
	DurationFormated string
	TrackURL    string
	TrackId     string
	Images      string
	*/
	for _, tra_items := range S.Tracks.Items {
		// Remplissage des données d'un track
		trackData := structure.Html_TrackData{
			AlbumURL:         tra_items.Album.URL.Spotify,
			AlbumId:          tra_items.Album.Id,
			AlbumName:        tra_items.Album.Name,
			ReleaseDate:      tra_items.Album.ReleaseDate,
			TotalTracks:      tra_items.Album.TotalTracks,
			Artists:          FormatArtists(tra_items.Artists),
			TrackName:        tra_items.Name,
			DurationMs:       tra_items.DurationMs,
			DurationFormated: FormatDuration(tra_items.DurationMs),
			TrackURL:         tra_items.URL.Spotify,
			TrackId:          tra_items.Id,
			Images:           GetImageAtIndex(tra_items.Album.Images, 1),
		}
		html_S.TrackData = append(html_S.TrackData, trackData)

	}
	// Boucle pour les artists
	/*ArtistData[]
	ArtistURL   string
	NbFollowers int
	Genres      []string
	ArtistId    string
	ArtistName  string
	Images      string
	*/
	for _, art_items := range S.Artists.Items {
		artistData := structure.Html_ArtistData{
			ArtistURL:   art_items.URL.Spotify,
			NbFollowers: art_items.Followers.Total,
			Genres:      art_items.Genres,
			ArtistId:    art_items.Id,
			ArtistName:  art_items.Name,
			Images:      GetImageAtIndex(art_items.Images, 1),
		}
		html_S.ArtistData = append(html_S.ArtistData, artistData)
	}
	// Boucle pour les albums
	/*AlbumData[]
	TotalTracks int
	AlbumURL    string
	AlbumId     string
	AlbumName   string
	ReleaseDate string
	Artists[]
		ArtistURL  string
		ArtistId   string
		ArtistName string
	Images 		string
	*/
	for _, alb_items := range S.Albums.Items {
		albumData := structure.Html_AlbumData{
			TotalTracks: alb_items.TotalTracks,
			AlbumURL:    alb_items.URL.Spotify,
			AlbumId:     alb_items.Id,
			AlbumName:   alb_items.Name,
			ReleaseDate: alb_items.ReleaseDate,
			Artists:     FormatArtists(alb_items.Artists),
			Images:      GetImageAtIndex(alb_items.Images, 1),
		}
		html_S.AlbumData = append(html_S.AlbumData, albumData)
	}
	return html_S
}

// FormatArtists formate une slice d'Api_Artist en une slice de Html_Items_ArtistData
func FormatArtists(Artist_Items []structure.Api_Artist) []structure.Html_Items_ArtistData {
	var artists []structure.Html_Items_ArtistData
	for _, art_items := range Artist_Items {
		artistData := structure.Html_Items_ArtistData{
			ArtistURL:  art_items.URL.Spotify,
			ArtistId:   art_items.Id,
			ArtistName: art_items.Name,
		}
		artists = append(artists, artistData)
	}
	return artists
}

// GetImageAtIndex vérifie si l'index de la slice d'images existe et retourne son URL, sinon retourne une chaîne vide
func GetImageAtIndex(Img_Items []structure.Api_Images, index int) string {
	if len(Img_Items) > index {
		return Img_Items[index].URL
	}
	return ""
}

// FormatDuration convertit une durée en millisecondes en une chaîne de caractères au format "mm:ss"
func FormatDuration(ms int) string {
	/*	- time.Duration → type Go qui représente une durée
		- ms est en millisecondes → on multiplie par time.Millisecond pour créer un objet time.Duration
		- d.Minutes() → retourne la durée totale en minutes flottantes (float64). Exemple : 134012 ms → 2.23353 minutes.
		- int(...) → convertit en entier → on garde juste la partie entière. Résultat : 2 minutes.
	*/
	d := time.Duration(ms) * time.Millisecond
	minutes := int(d.Minutes())
	seconds := int(d.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}
