package data

import (
	"Groupie_Tracker/structure"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// InitSessionData initialise les données de la page web
func InitSessionData() *structure.SessionData {
	fmt.Printf("data.InitSessionData - Initialisation des données de la session\n\n")
	// Initialisation des données de la session
	return &structure.SessionData{
		LogIn: false,
		Utilisateur: structure.Utilisateur{
			Nom:     "",
			Favoris: []string{},
		},
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

func TemplateHTMLArtist(Artist structure.Api_Artist, TopTracks structure.Api_TopTracks, Albums structure.Api_ArtistAlbums) structure.Html_Artist {
	// Remplissage des données pour le template HTML
	var html_A structure.Html_Artist
	// Données de l'artiste
	html_A.Artist.ArtistId = Artist.Id
	html_A.Artist.ArtistName = Artist.Name
	html_A.Artist.NbFollowers = Artist.Followers.Total
	for _, A_genres := range Artist.Genres {
		html_A.Artist.Genres = append(html_A.Artist.Genres, A_genres)
	}
	html_A.Artist.Images = GetImageAtIndex(Artist.Images, 1)
	html_A.Artist.ArtistURL = Artist.URL.Spotify
	// Données des Top Tracks
	for _, tt_items := range TopTracks.Tracks {
		trackData := structure.Html_TrackData{
			AlbumURL:         tt_items.Album.URL.Spotify,
			AlbumId:          tt_items.Album.Id,
			AlbumName:        tt_items.Album.Name,
			ReleaseDate:      tt_items.Album.ReleaseDate,
			TotalTracks:      tt_items.Album.TotalTracks,
			Artists:          FormatArtists(tt_items.Artists),
			TrackName:        tt_items.Name,
			DurationMs:       tt_items.DurationMs,
			DurationFormated: FormatDuration(tt_items.DurationMs),
			TrackURL:         tt_items.URL.Spotify,
			TrackId:          tt_items.Id,
			Images:           GetImageAtIndex(tt_items.Album.Images, 1),
		}
		html_A.TopTracks = append(html_A.TopTracks, trackData)
	}
	// Données des Albums
	for _, alb_items := range Albums.Items {
		albumData := structure.Html_AlbumData{
			TotalTracks: alb_items.TotalTracks,
			AlbumURL:    alb_items.URL.Spotify,
			AlbumId:     alb_items.Id,
			AlbumName:   alb_items.Name,
			ReleaseDate: alb_items.ReleaseDate,
			Artists:     FormatArtists(alb_items.Artists),
			Images:      GetImageAtIndex(alb_items.Images, 1),
		}
		html_A.Albums = append(html_A.Albums, albumData)
	}
	return html_A
}

func TemplateHTMLAlbums(AlbumTracks structure.Api_AlbumsTracks) structure.Html_AlbumTracks {
	// Remplissage des données pour le template HTML
	var html_Al structure.Html_AlbumTracks
	// Données de l'album
	html_Al.AlbumID = AlbumTracks.AlbumID
	html_Al.Images = GetImageAtIndex(AlbumTracks.Images, 1)
	html_Al.AlbumName = AlbumTracks.AlbumName
	html_Al.Release_date = AlbumTracks.Release_date
	html_Al.AlbumArtists = FormatArtists(AlbumTracks.AlbumArtists)
	// Données des Tracks de l'album
	for _, at_items := range AlbumTracks.Tracks.Items {
		trackData := structure.Html_AlbumTracks_Items{
			Artists:          FormatArtists(at_items.Artists),
			DurationMs:       at_items.DurationMs,
			DurationFormated: FormatDuration(at_items.DurationMs),
			TrackURL:         at_items.URL.Spotify,
			TrackId:          at_items.Id,
			TrackName:        at_items.Name,
		}
		html_Al.Items = append(html_Al.Items, trackData)
	}
	return html_Al
}

func TemplateHTMLTrack(Track structure.Api_Track) structure.Html_TrackData {
	// Remplissage des données pour le template HTML
	var html_T structure.Html_TrackData
	// Données du track
	html_T.AlbumURL = Track.Album.URL.Spotify
	html_T.AlbumId = Track.Album.Id
	html_T.AlbumName = Track.Album.Name
	html_T.ReleaseDate = Track.Album.ReleaseDate
	html_T.TotalTracks = Track.Album.TotalTracks
	html_T.Artists = FormatArtists(Track.Artists)
	html_T.TrackName = Track.Name
	html_T.DurationMs = Track.DurationMs
	html_T.DurationFormated = FormatDuration(Track.DurationMs)
	html_T.TrackURL = Track.URL.Spotify
	html_T.TrackId = Track.Id
	html_T.Images = GetImageAtIndex(Track.Album.Images, 1)
	return html_T
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

func GetIdFromUrl(r *http.Request) string {
	// Récupération de l'ID de l'artiste depuis l'URL
	// Exemple d'URL : /artiste/12345
	// On découpe l'URL coupant par les "/"
	// Exemple : ["", "artiste", "12345"]
	parties := strings.Split(r.URL.Path, "/")
	// Si les parties sont inférieures ou égales à 2 ou que la 3ème partie est vide, on retourne une chaîne vide
	if len(parties) < 3 {
		return ""
	}
	return parties[2]
}

func GetPageOffset(pagestr string) (int, int) {
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
	fmt.Printf("data - GetPageOffset - Page : %d | Offset : %d\n\n", page, offset)
	return page, offset
}

// TemplateErreur prépare les données d'erreur pour le template HTML en fonction du type (API ou Application (Serveur Go)) et du statut de l'erreur
func TemplateErreur(status int, errType string) structure.Html_Erreur {
	switch errType {
	// Gestion des erreurs provenant de l'API Spotify
	case "API":
		switch status {
		//400 — Mauvaise requête - Paramètres manquants ou invalides (q, type, offset, etc.))
		case 400:
			return structure.Html_Erreur{
				Status:  400,
				Message: "Requête invalide. Veuillez vérifier les paramètres fournis.",
			}
		// 401 — Non autorisé (token invalide / expiré)
		case 401:
			return structure.Html_Erreur{
				Status:  401,
				Message: "Accès interdit à cette ressource.",
			}
		//403 Accès interdit - Droits insuffisants (le token n'a pas les permissions nécessaires)
		case 403:
			return structure.Html_Erreur{
				Status:  403,
				Message: "Accès refusé. Vous n'avez pas les permissions nécessaires.",
			}
		//404 Ressource non trouvée (Artiste / album / titre inexistant ou supprimé)
		case 404:
			return structure.Html_Erreur{
				Status:  404,
				Message: "Ressource introuvable sur Spotify.",
			}
		//429 Trop de requêtes - Rate limit Spotify atteint
		case 429:
			return structure.Html_Erreur{
				Status:  429,
				Message: "Trop de requêtes envoyées. Veuillez réessayer plus tard.",
			}
		//500 Erreur interne du service Spotify (server Spotify en erreur)
		case 500:
			return structure.Html_Erreur{
				Status:  500,
				Message: "Erreur interne du service Spotify.",
			}
		//503 Service indisponible - Maintenance ou surcharge du service Spotify
		case 503:
			return structure.Html_Erreur{
				Status:  503,
				Message: "Service Spotify momentanément indisponible.",
			}
		// Erreur par défaut pour les autres codes d'erreur
		default:
			return structure.Html_Erreur{
				Status:  status,
				Message: "Une erreur d'API est survenue. Veuillez réessayer plus tard.",
			}
		}
	// Gestion des erreurs provenant de l'application (Serveur Go)
	case "APP":
		switch status {
		// Paramètre invalide (ID vide, page négative, etc.)
		case 400:
			return structure.Html_Erreur{
				Status:  400,
				Message: "Paramètre requis manquant dans l'URL.",
			}
		// Paramètre manquant dans l’URL (/artiste/ sans ID)
		case 404:
			return structure.Html_Erreur{
				Status:  404,
				Message: "Paramètre invalide fourni.",
			}
		// Erreur de parsing des données - JSON / conversion (json.Unmarshal, strconv.Atoi, etc.)
		case 500:
			return structure.Html_Erreur{
				Status:  500,
				Message: "Erreur lors du traitement des données.",
			}
		// Timeout API - Timeout HTTP (http.Client{Timeout: ...})
		case 504:
			return structure.Html_Erreur{
				Status:  504,
				Message: "Délai d'attente dépassé lors de la communication avec l'API Spotify.",
			}
		default:
			return structure.Html_Erreur{
				Status:  status,
				Message: "Une erreur d'application est survenue. Veuillez réessayer plus tard.",
			}
		}
	// Erreur par défaut pour les types inconnus
	default:
		return structure.Html_Erreur{
			Status:  status,
			Message: "Une erreur inconnue est survenue.",
		}
	}
}

func CreationCompte(nomUtilisateur string, mdp string) string {
	//Verification si l'utilisateur existe déjà
	if VerifUtilisateur(nomUtilisateur) == "" {
		return "Nom d'utilisateur déjà existant"
	}

	// Récupération des données existantes //

	// Lecture du fichier des utilisateurs
	data, err := os.ReadFile("compte/compte.json")
	if err != nil { //si erreur lors de la lecture
		return err.Error() //retourne erreur
	}

	//si le fichier a des data on essaye de parser (récupération) en slice d'Utilisateur
	// Déclaration de la slice des utilisateurs
	var utilisateurs []structure.Utilisateur
	if len(data) > 0 {
		err := json.Unmarshal(data, &utilisateurs)
		if err != nil { //decode json
			return err.Error() //erreur si invalide
		}
	}

	// Ajout du nouvel utilisateur //
	// Création du nouvel utilisateur
	nouvelUtilisateur := structure.Utilisateur{
		Nom:        nomUtilisateur,
		MotDePasse: mdp,
		Favoris:    []string{},
	}

	// Ajout du nouvel utilisateur à la liste
	utilisateurs = append(utilisateurs, nouvelUtilisateur)

	// Conversion de la liste des utilisateurs en JSON
	utilisateursJSON, err := json.MarshalIndent(utilisateurs, "", "  ")
	if err != nil {
		return err.Error()
	}

	// Écriture dans le fichier
	// Permissions 0644 : propriétaire en lecture/écriture, groupe et autres en lecture seule
	err = os.WriteFile("compte/compte.json", utilisateursJSON, 0644)
	if err != nil {
		return err.Error()
	}
	return "" //retourne chaîne vide si pas d'erreur
}

func ConnexionCompte(nomUtilisateur string, mdp string, sessionData *structure.SessionData) string {
	//Verification si l'utilisateur existe déjà
	if VerifUtilisateur(nomUtilisateur) != "" {
		return "Utilisateur inéxistant. Veuillez vous inscrire."
	}
	// Récupération des données existantes //

	// Lecture du fichier des utilisateurs
	data, err := os.ReadFile("compte/compte.json")
	if err != nil { //si erreur lors de la lecture
		return err.Error() //retourne erreur
	}

	//si le fichier a des data on essaye de parser (récupération) en slice d'Utilisateur
	// Déclaration de la slice des utilisateurs
	var utilisateurs []structure.Utilisateur
	if len(data) <= 0 {
		return "Aucun utilisateur existant sur le système. Veuillez vous inscrire."
	} else {
		err := json.Unmarshal(data, &utilisateurs)
		if err != nil { //decode json
			return err.Error() //erreur si invalide
		}
	}
	// Vérification des identifiants
	for _, u := range utilisateurs {
		if u.Nom == nomUtilisateur {
			if u.MotDePasse == mdp {
				// Connexion réussie
				sessionData.LogIn = true
				sessionData.Utilisateur = u
				return ""
			} else {
				// Mot de passe incorrect
				return "Mot de passe incorrect"
			}
		}
	}
	return "" //retourne chaîne vide si pas d'erreur
}

func VerifUtilisateur(nomUtilisateur string) string {
	// Récupération des données existantes //

	// Lecture du fichier des utilisateurs
	data, err := os.ReadFile("compte/compte.json")
	if err != nil { //si erreur lors de la lecture
		return err.Error() //retourne erreur
	}

	//si le fichier a des data on essaye de parser (récupération) en slice d'Utilisateur
	// Déclaration de la slice des utilisateurs
	var utilisateurs []structure.Utilisateur
	if len(data) > 0 {
		err := json.Unmarshal(data, &utilisateurs)
		if err != nil { //decode json
			return err.Error() //erreur si invalide
		}
	}

	// Vérification et ajout du nouvel utilisateur //

	// Vérification si l'utilisateur existe déjà
	for _, u := range utilisateurs {
		if u.Nom == nomUtilisateur {
			// Utilisateur trouvé
			return ""
		}
	}
	return "Utilisateur inéxistant"
}
