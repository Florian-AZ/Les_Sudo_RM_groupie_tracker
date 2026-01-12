package api

import (
	"Groupie_Tracker/structure"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

func SearchBar(token string, quary string, offset int) structure.Api_Recherche {
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
	q.Add("offset", strconv.Itoa(offset))
	req.URL.RawQuery = q.Encode()

	// Execution de la requête HTTP vars L'API
	res, errResp := httpClient.Do(req)
	if errResp != nil {
		// Si une erreur est survenue lors de l'appel à l'API
		fmt.Printf("api.SearchBar - Erreur - Do : %s\n\n", errResp.Error())
		return structure.Api_Recherche{
			Error: structure.Api_Error{
				Status:  504,
				Message: errResp.Error()}}
	}

	// Assurer la fermeture du corps de la réponse HTTP une fois la fonction terminée
	defer res.Body.Close()

	// Si le statut de la réponse HTTP n'est pas 200 OK (c'est à dire une erreur)
	if res.StatusCode != http.StatusOK {
		return structure.Api_Recherche{
			Error: structure.Api_Error{
				Status:  res.StatusCode,
				Message: http.StatusText(res.StatusCode),
			},
		}
	}

	// Lecture et récupération du corps de la requête HTTP
	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		fmt.Printf("api.SearchBar - Erreur - ReadAll : %s\n\n", errBody.Error())
		return structure.Api_Recherche{
			Error: structure.Api_Error{
				Status:  500,
				Message: errBody.Error(),
			}}
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

func GetArtistData(Token string, id string, offset int) structure.Api_Artist {
	// URL de L'API
	urlApi := "https://api.spotify.com/v1/artists/" + id

	// Initialisation du client HTTP qui va émettre/demander les requêtes
	httpClient := http.Client{
		Timeout: time.Second * 2, // Timeout apres 2sec
	}

	// Création de la requête HTTP vers L'API avec initialisation de la methode HTTP, la route et le corps de la requête
	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		fmt.Printf("api.GetArtistData - Erreur - NewRequest : %s\n\n", errReq.Error())
	}

	//Paramètre de type query à inserer à la req GET
	q := req.URL.Query()
	q.Add("limit", "10")
	q.Add("offset", strconv.Itoa(offset))
	req.URL.RawQuery = q.Encode()

	// Ajout d'une métadonnée dans le header
	req.Header.Add("Authorization", "Bearer "+Token)

	// Execution de la requête HTTP vars L'API
	res, errResp := httpClient.Do(req)
	if errResp != nil {
		fmt.Printf("api.GetArtistData - Erreur - Do : %s\n\n", errResp.Error())
		return structure.Api_Artist{
			Error: structure.Api_Error{
				Status:  504,
				Message: errResp.Error()}}
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return structure.Api_Artist{
			Error: structure.Api_Error{
				Status:  res.StatusCode,
				Message: http.StatusText(res.StatusCode),
			},
		}
	}

	// Lecture et récupération du corps de la requête HTTP
	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		fmt.Printf("api.GetArtistData - Erreur - ReadAll : %s\n\n", errBody.Error())
	}

	// Déclaration de la variable qui va contenir les données
	var decodeData structure.Api_Artist

	// Decodage des données en format JSON et ajout des donnée à la variable: decodeData
	json.Unmarshal(body, &decodeData)

	// Affichage des données
	if decodeData.Error.Message != "" {
		return decodeData
	} else {
		fmt.Printf("api.GetArtistData - Succès - artiste %s récupéré : \n\n", id)
		fmt.Printf("api.GetArtistData - Succès - artiste brut: %v\n\n", decodeData)
		return decodeData
	}
}

func GetArtistTopTracks(Token string, id string) structure.Api_TopTracks {
	// URL de L'API
	urlApi := fmt.Sprintf("https://api.spotify.com/v1/artists/%s/top-tracks", id)

	// Initialisation du client HTTP qui va émettre/demander les requêtes
	httpClient := http.Client{
		Timeout: time.Second * 2, // Timeout apres 2sec
	}

	// Création de la requête HTTP vers L'API avec initialisation de la methode HTTP, la route et le corps de la requête
	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		fmt.Printf("api.GetArtistTopTracks - Erreur - NewRequest : %s\n\n", errReq.Error())
	}

	// Ajout d'une métadonnée dans le header
	req.Header.Add("Authorization", "Bearer "+Token)

	// Execution de la requête HTTP vars L'API
	res, errResp := httpClient.Do(req)
	if errResp != nil {
		fmt.Printf("api.GetArtistTopTracks - Erreur - Do : %s\n\n", errResp.Error())
		return structure.Api_TopTracks{
			Error: structure.Api_Error{
				Status:  504,
				Message: errResp.Error()}}
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return structure.Api_TopTracks{
			Error: structure.Api_Error{
				Status:  res.StatusCode,
				Message: http.StatusText(res.StatusCode),
			},
		}
	}

	// Lecture et récupération du corps de la requête HTTP
	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		fmt.Printf("api.GetArtistTopTracks - Erreur - ReadAll : %s\n\n", errBody.Error())
	}

	// Déclaration de la variable qui va contenir les données
	var decodeData structure.Api_TopTracks

	// Decodage des données en format JSON et ajout des donnée à la variable: decodeData
	json.Unmarshal(body, &decodeData)

	// Affichage des données
	if decodeData.Error.Message != "" {
		return decodeData
	} else {
		fmt.Printf("api.GetArtistTopTracks - Succès - album de %s récupéré : \n\n", id)
		fmt.Printf("api.GetArtistTopTracks - Succès - album brut: %v\n\n", decodeData)
		return decodeData
	}
}

func GetArtistAlbums(Token string, id string, offset int) structure.Api_ArtistAlbums {
	// URL de L'API
	urlApi := fmt.Sprintf("https://api.spotify.com/v1/artists/%s/albums", id)

	// Initialisation du client HTTP qui va émettre/demander les requêtes
	httpClient := http.Client{
		Timeout: time.Second * 2, // Timeout apres 2sec
	}

	// Création de la requête HTTP vers L'API avec initialisation de la methode HTTP, la route et le corps de la requête
	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		fmt.Printf("api.GetArtistAlbums - Erreur - NewRequest : %s\n\n", errReq.Error())
	}

	//Paramètre de type query à inserer à la req GET
	q := req.URL.Query()
	q.Add("limit", "10")
	q.Add("offset", strconv.Itoa(offset))
	req.URL.RawQuery = q.Encode()

	// Ajout d'une métadonnée dans le header
	req.Header.Add("Authorization", "Bearer "+Token)

	// Execution de la requête HTTP vars L'API
	res, errResp := httpClient.Do(req)
	if errResp != nil {
		fmt.Printf("api.GetArtistAlbums - Erreur - Do : %s\n\n", errResp.Error())
		return structure.Api_ArtistAlbums{
			Error: structure.Api_Error{
				Status:  504,
				Message: errResp.Error()}}
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return structure.Api_ArtistAlbums{
			Error: structure.Api_Error{
				Status:  res.StatusCode,
				Message: http.StatusText(res.StatusCode),
			},
		}
	}

	// Lecture et récupération du corps de la requête HTTP
	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		fmt.Printf("api.GetArtistAlbums - Erreur - ReadAll : %s\n\n", errBody.Error())
	}

	// Déclaration de la variable qui va contenir les données
	var decodeData structure.Api_ArtistAlbums

	// Decodage des données en format JSON et ajout des donnée à la variable: decodeData
	json.Unmarshal(body, &decodeData)

	// Affichage des données
	if decodeData.Error.Message != "" {
		return decodeData
	} else {
		fmt.Printf("api.GetArtistAlbums - Succès - album de %s récupéré : \n\n", id)
		fmt.Printf("api.GetArtistAlbums - Succès - album brut: %v\n\n", decodeData)
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
		// Si une erreur est survenue lors de l'appel à l'API
		fmt.Printf("api. - Erreur - Do : %s\n\n", errResp.Error())
		return structure.Exemple{
			Error: structure.Exemple_Error{
				Status:  504,
				Message: errResp.Error()}}
	}

	// Assurer la fermeture du corps de la réponse HTTP une fois la fonction terminée
	defer res.Body.Close()

	// Si le statut de la réponse HTTP n'est pas 200 OK (c'est à dire une erreur)
	if res.StatusCode != http.StatusOK {
		return structure.Exemple{
			Error: structure.Exemple_Error{
				Status:  res.StatusCode,
				Message: http.StatusText(res.StatusCode),
			},
		}
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
