package structure

type SessionData struct {
	LogIn       bool
	Utilisateur Utilisateur
}

// Ensemlble des structures pour l'API Spotify //

// Structure avec le token d'accès à l'API Spotify
type Api_Token struct {
	AccessToken      string `json:"access_token"`
	ErrorStatus      int    `json:"status"`
	ErrorDescription string `json:"error_description"`
}

// Structure avec les données de la recherche
type Api_Recherche struct {
	Tracks Api_AllTracks `json:"tracks"`
	/*JSON Structure
	tracks
	items[]
		album
			external_urls
			id
			images[1]
			name
			release_date
			total_tracks

		artists[]
			external_urls
			id
			name

		duration_ms
		external_urls
		id
		name
	*/
	Artists struct {
		Items []Api_Artist `json:"items"`
		/*JSON Structure
		artists
			items[]
				external_urls
				folowers
					total

				genres[]
				id
				images[1]
				name
		*/
	} `json:"artists"`
	Albums Api_AllAlbums `json:"albums"`
	/*JSON Structure
	albums
		items[]
			total_tracks
			external_urls
			id
			images[1]
			name
			release_date
			artists[]
				external_urls
				id
				name
	*/
	Error Api_Error `json:"error"`
}

// structure avec les données de l'album que l'on veut récupérer
type Api_AllAlbums struct {
	Items []Api_Albums `json:"items"`
	Error Api_Error    `json:"error"`
	/*JSON Structure
	albums
		items[]
	*/
}

type Api_Albums struct {
	TotalTracks int              `json:"total_tracks"`
	URL         Api_ExternalUrls `json:"external_urls"`
	Id          string           `json:"id"`
	Images      []Api_Images     `json:"images"`
	Name        string           `json:"name"`
	ReleaseDate string           `json:"release_date"`
	Artists     []Api_Artist     `json:"artists"`
	/*JSON Structure
	total_tracks
	external_urls
	id
	images[1]
	name
	release_date
	artists[]
		external_urls
		id
		name
	*/
}

type Api_AllTracks struct {
	Items []Api_Track `json:"items"`
	Error Api_Error   `json:"error"`
}

type Api_Track struct {
	Album      Api_Track_Album  `json:"album"`
	Artists    []Api_Artist     `json:"artists"`
	DurationMs int              `json:"duration_ms"`
	URL        Api_ExternalUrls `json:"external_urls"`
	Id         string           `json:"id"`
	Name       string           `json:"name"`
	Error      Api_Error        `json:"error"`
	/*JSON Structure
	Tracks
		album
		artists[]
		duration_ms
		external_urls
		id
		name
	*/
}

type Api_Track_Album struct {
	URL         Api_ExternalUrls `json:"external_urls"`
	Id          string           `json:"id"`
	Images      []Api_Images     `json:"images"`
	Name        string           `json:"name"`
	ReleaseDate string           `json:"release_date"`
	TotalTracks int              `json:"total_tracks"`
	/*JSON Structure
	album
		external_urls
		id
		images[1]
		name
		release_date
		total_tracks
	*/
}

type Api_TopTracks struct {
	Tracks []Api_Track `json:"tracks"`
	Error  Api_Error   `json:"error"`
}

type Api_Artist struct {
	URL       Api_ExternalUrls `json:"external_urls"`
	Followers Api_Followers    `json:"followers"`
	Genres    []string         `json:"genres"`
	Id        string           `json:"id"`
	Images    []Api_Images     `json:"images"`
	Name      string           `json:"name"`
	Error     Api_Error        `json:"error"`
	/*JSON Structure
	Arists
		external_urls
		folowers
		genres[]
		id
		images[1]
		name
	*/
}

type Api_ArtistAlbums struct {
	Items []Api_Albums `json:"items"`
	Error Api_Error    `json:"error"`
}

type Api_AlbumsTracks struct {
	AlbumID      string       `json:"id"`
	Images       []Api_Images `json:"images"`
	AlbumName    string       `json:"name"`
	Release_date string       `json:"release_date"`
	AlbumArtists []Api_Artist `json:"artists"`
	Tracks       struct {
		Items []Api_AlbumsTracks_Items `json:"items"`
	} `json:"tracks"`
	Error Api_Error `json:"error"`
}

type Api_AlbumsTracks_Items struct {
	Artists    []Api_Artist     `json:"artists"`
	DurationMs int              `json:"duration_ms"`
	URL        Api_ExternalUrls `json:"external_urls"`
	Id         string           `json:"id"`
	Name       string           `json:"name"`
}

type Api_Followers struct {
	Total int `json:"total"`
	/*JSON Structure
	folowers
		total
	*/
}

type Api_ExternalUrls struct {
	Spotify string `json:"spotify"`
	/*JSON Structure
	external_urls
	    spotify
	*/
}

type Api_Images struct {
	URL string `json:"url"`
	/*JSON Structure
	images[]
		url
	*/
}

type Api_Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Ensemble des structures formatées pour le template HTML //

// Structure avec les données formatées pour le template HTML
type Html_Recherche struct {
	/* Donnée de la recherche que l'on envoie au html
	TrackData[]
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
	ArtistData[]
		ArtistURL   string
		NbFollowers int
		Genres      []string
		ArtistId    string
		ArtistName  string
		Images      string
	AlbumData[]
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
	TrackData  []Html_TrackData
	ArtistData []Html_ArtistData
	AlbumData  []Html_AlbumData
}

type Html_Artist struct {
	Artist    Html_ArtistData
	TopTracks []Html_TrackData
	Albums    []Html_AlbumData
}

// Sous-Structure de Html_Recherche pour les tracks
type Html_TrackData struct {
	AlbumURL         string
	AlbumId          string
	AlbumName        string
	ReleaseDate      string
	TotalTracks      int
	Artists          []Html_Items_ArtistData
	TrackName        string
	DurationMs       int
	DurationFormated string
	TrackURL         string
	TrackId          string
	Images           string
}

// Sous-Structure de Html_Recherche pour les artists
type Html_ArtistData struct {
	ArtistURL   string
	NbFollowers int
	Genres      []string
	ArtistId    string
	ArtistName  string
	Images      string
}

// Sous-Structure de Html_Recherche pour les albums
type Html_AlbumData struct {
	TotalTracks int
	AlbumURL    string
	AlbumId     string
	AlbumName   string
	ReleaseDate string
	Artists     []Html_Items_ArtistData
	Images      string
}

type Html_AlbumTracks struct {
	AlbumID      string
	Images       string
	AlbumName    string
	Release_date string
	AlbumArtists []Html_Items_ArtistData
	Items        []Html_AlbumTracks_Items
}

type Html_AlbumTracks_Items struct {
	Artists          []Html_Items_ArtistData
	DurationMs       int
	DurationFormated string
	TrackURL         string
	TrackId          string
	TrackName        string
}

// Sous-Structure des artistes dans les tracks et albums
type Html_Items_ArtistData struct {
	ArtistURL  string
	ArtistId   string
	ArtistName string
}

type Html_Erreur struct {
	Status  int
	Message string
}

type Html_Favoris struct {
	Titres   []Html_Favoris_Titre
	Artistes []Html_Favoris_Artiste
	Albums   []Html_Favoris_Album
}

type Html_Favoris_Titre struct {
	Id       string
	Nom      string
	Artistes []Html_Items_ArtistData
	URL      string
	Image    string
}

type Html_Favoris_Artiste struct {
	Id    string
	Nom   string
	URL   string
	Image string
}

type Html_Favoris_Album struct {
	Id         string
	Nom        string
	DateSortie string
	Artistes   []Html_Items_ArtistData
	Image      string
}

// Ensemble des structures regroupant les données nécéssaires pour chaque page web //

// Structure des données pour la page d'accueil
type PageData_Accueil struct {
	LogIn bool
}

// Structure des données pour la page recherche
type PageData_Recherche struct {
	LogIn       bool
	SearchData  Html_Recherche
	SearchQuery string
	Pagination  Pagination
}

type PageData_Artiste struct {
	LogIn                 bool
	ArtistData            Html_Artist
	PaginationTotalTracks Pagination
	PaginationAlbums      Pagination
	ErrTotalTracks        Html_Erreur
	ErrAlbums             Html_Erreur
}

type PageData_Album struct {
	LogIn       bool
	AlbumTracks Html_AlbumTracks
	Pagination  Pagination
}

type PageData_Titre struct {
	LogIn     bool
	TrackData Html_TrackData
}

type PageData_Connexion_Inscription struct {
	Erreur string
}

type PageData_Favoris struct {
	NomUtilisateur string
	Favoris        Html_Favoris
	Pagination     Pagination
}

type PageData_Erreur struct {
	LogIn         bool
	ErreurStatus  int
	ErreurMessage string
}

// Structure de la pagination
type Pagination struct {
	Page       int
	ASuivant   bool
	APrecedent bool
	PageSuiv   int
	PagePrec   int
}

// Structure des données pour les utilisateurs et leurs favoris //

type Utilisateur struct {
	Nom        string              `json:"Nom"`
	MotDePasse string              `json:"MotDePasse"`
	Favoris    Utilisateur_Favoris `json:"favoris"`
}

type Utilisateur_Favoris struct {
	IdTitres   []Favoris_Id `json:"id_titres"`
	IdArtistes []Favoris_Id `json:"id_artistes"`
	IdAlbums   []Favoris_Id `json:"id_albums"`
}

type Favoris_Id struct {
	Id string `json:"id"`
}
