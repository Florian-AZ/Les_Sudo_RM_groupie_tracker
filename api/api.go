package api

import (
	"Groupie_Tracker/structure"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func SearchBar(token string, quary string) structure.Api_Recherche {
	// URL de L'API
	urlApi := "https://api.spotify.com/v1/search"

	// Initialisation du client HTTP qui va émettre/demander les requêtes
	httpClient := http.Client{
		Timeout: time.Second * 2, // Timeout apres 2sec
	}

	// Création de la requête HTTP vers L'API avec initialisation de la methode HTTP, la route et le corps de la requête
	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		fmt.Printf("api.SearchBar - Erreur - NewRequest : %s\n\n", errReq.Error())
	}

	// Ajout d'une métadonnée dans le header
	req.Header.Add("Authorization", "Bearer "+token)

	//Paramètre de type query à inserer à la req GET
	q := req.URL.Query()
	q.Add("q", quary)
	q.Add("type", "artist,track,album")
	q.Add("limit", "10")
	req.URL.RawQuery = q.Encode()

	// Execution de la requête HTTP vars L'API
	res, errResp := httpClient.Do(req)
	if errResp != nil {
		fmt.Printf("api.SearchBar - Erreur - Do : %s\n\n", errResp.Error())
		return structure.Api_Recherche{Error: structure.Api_Error{Message: errResp.Error()}}
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	// Lecture et récupération du corps de la requête HTTP
	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		fmt.Printf("api.SearchBar - Erreur - ReadAll : %s\n\n", errBody.Error())
	}

	// Déclaration de la variable qui va contenir les données
	var decodeData structure.Api_Recherche

	// Decodage des données en format JSON et ajout des donnée à la variable: decodeData
	json.Unmarshal(body, &decodeData)

	// Affichage des données
	if decodeData.Error.Message != "" {
		fmt.Printf("api.SearchBar - Erreur - %s\n\n", decodeData.Error.Message)
		return decodeData
	} else {
		fmt.Printf("api.SearchBar - Succès -  brut: %v\n\n", decodeData)
		return decodeData
	}
}

func GetAlbum(Token string, id string) structure.Api_AllAlbums {
	// URL de L'API
	urlApi := fmt.Sprintf("https://api.spotify.com/v1/artists/%s/albums", id)

	// Initialisation du client HTTP qui va émettre/demander les requêtes
	httpClient := http.Client{
		Timeout: time.Second * 2, // Timeout apres 2sec
	}

	// Création de la requête HTTP vers L'API avec initialisation de la methode HTTP, la route et le corps de la requête
	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		fmt.Printf("api.GetAlbum - Erreur - NewRequest : %s\n\n", errReq.Error())
	}

	//Paramètre de type query à inserer à la req GET
	q := req.URL.Query()
	q.Add("include_groups", "album")
	q.Add("limit", "10")
	req.URL.RawQuery = q.Encode()

	// Ajout d'une métadonnée dans le header
	req.Header.Add("Authorization", "Bearer "+Token)

	// Execution de la requête HTTP vars L'API
	res, errResp := httpClient.Do(req)
	if errResp != nil {
		fmt.Printf("api.GetAlbum - Erreur - Do : %s\n\n", errResp.Error())
		return structure.Api_AllAlbums{Error: structure.Api_Error{Message: errResp.Error()}}
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	// Lecture et récupération du corps de la requête HTTP
	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		fmt.Printf("api.GetAlbum - Erreur - ReadAll : %s\n\n", errBody.Error())
	}

	// Déclaration de la variable qui va contenir les données
	var decodeData structure.Api_AllAlbums

	// Decodage des données en format JSON et ajout des donnée à la variable: decodeData
	json.Unmarshal(body, &decodeData)

	// Affichage des données
	if decodeData.Error.Message != "" {
		return decodeData
	} else {
		fmt.Printf("api.GetAlbum - Succès - album de %s récupéré : \n\n", id)
		fmt.Printf("api.GetAlbum - Succès - album brut: %v\n\n", decodeData.Items)
		return decodeData
	}
}

func GetTrack(Token string, id string) structure.Api_Track {
	// URL de L'API
	urlApi := "https://api.spotify.com/v1/tracks/" + id

	// Initialisation du client HTTP qui va émettre/demander les requêtes
	httpClient := http.Client{
		Timeout: time.Second * 2, // Timeout apres 2sec
	}

	// Création de la requête HTTP vers L'API avec initialisation de la methode HTTP, la route et le corps de la requête
	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		fmt.Printf("api.GetTrack - Erreur - NewRequest : %s\n\n", errReq.Error())
	}

	// Ajout d'une métadonnée dans le header
	req.Header.Add("Authorization", "Bearer "+Token)

	// Execution de la requête HTTP vars L'API
	res, errResp := httpClient.Do(req)
	if errResp != nil {
		fmt.Printf("api.GetTrack - Erreur - Do : %s\n\n", errResp.Error())
		return structure.Api_Track{Error: structure.Api_Error{Message: errResp.Error()}}
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	// Lecture et récupération du corps de la requête HTTP
	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		fmt.Printf("api.GetTrack - Erreur - ReadAll : %s\n\n", errBody.Error())
	}

	// Déclaration de la variable qui va contenir les données
	var decodeData structure.Api_Track

	// Decodage des données en format JSON et ajout des donnée à la variable: decodeData
	json.Unmarshal(body, &decodeData)

	// Affichage des données
	if decodeData.Error.Message != "" {
		fmt.Printf("api.GetTrack - Erreur - %s\n\n", decodeData.Error.Message)
		return decodeData
	} else {
		fmt.Printf("api.GetTrack - Succès - track de %s récupéré : \n\n", id)
		fmt.Printf("api.GetTrack - Succès - track brut: %v\n\n", decodeData)
		return decodeData
	}
}

/*
func ExempleApi(token string) structure.Exemple {
	// URL de L'API
	urlApi := "https://api.spotify.com/v1/"

	// Initialisation du client HTTP qui va émettre/demander les requêtes
	httpClient := http.Client{
		Timeout: time.Second * 2, // Timeout apres 2sec
	}

	// Création de la requête HTTP vers L'API avec initialisation de la methode HTTP, la route et le corps de la requête
	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		fmt.Printf("api. - Erreur - NewRequest : %s\n\n", errReq.Error())
	}

	// Ajout d'une métadonnée dans le header
	req.Header.Add("Authorization", "Bearer "+token)

	//Paramètre de type query à inserer à la req GET
	q := req.URL.Query()
	q.Add("")
	req.URL.RawQuery = q.Encode()

	// Execution de la requête HTTP vars L'API
	res, errResp := httpClient.Do(req)
	if errResp != nil {
		fmt.Printf("api. - Erreur - Do : %s\n\n", errResp.Error())
		return structure.Exemple{Error: structure.Error{Message: errResp.Error()}}
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	// Lecture et récupération du corps de la requête HTTP
	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		fmt.Printf("api. - Erreur - ReadAll : %s\n\n", errBody.Error())
	}

	// Déclaration de la variable qui va contenir les données
	var decodeData structure.Exemple

	// Decodage des données en format JSON et ajout des donnée à la variable: decodeData
	json.Unmarshal(body, &decodeData)

	// Affichage des données
	if decodeData.Error.Message != "" {
		fmt.Printf("api. - Erreur - %s\n\n", decodeData.Error.Message)
		return decodeData
	} else {
		fmt.Printf("api. - Succès -  brut: %v\n\n", decodeData)
		return decodeData
	}
}
*/
