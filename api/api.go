package api

import (
	"Groupie_Tracker/structure"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func GetToken() structure.Token {

	// URL de L'API
	urlApi := "https://accounts.spotify.com/api/token"

	// Initialisation du client HTTP qui va émettre/demander les requêtes
	httpClient := http.Client{
		Timeout: time.Second * 2, // Timeout apres 2sec
	}

	//Paramètres de type body à inserer à la req POST
	data := url.Values{}
	data.Set("grant_type", "client_credentials") //Name, Value
	data.Set("client_id", "967f549670ed4aefbfedf2b746202c75")
	data.Set("client_secret", "ad8b9b61687d4c3387784e0a7e34c54e")

	// Création de la requête HTTP vers L'API avec initialisation de la methode HTTP, la route et le corps de la requête
	req, errReq := http.NewRequest(http.MethodPost, urlApi, strings.NewReader(data.Encode())) // Méthode de req, url de l'API, Paramètres de la req (On converti les strings en flux lisible io.reader)
	if errReq != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errReq.Error())
	}

	// Ajout d'une métadonnée dans le header
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Execution de la requête HTTP vars L'API
	res, errResp := httpClient.Do(req)
	if errResp != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errResp.Error())
		return structure.Token{Error: errResp.Error()}
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	// Lecture et récupération du corps de la requête HTTP
	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errBody.Error())
	}

	// Déclaration de la variable qui va contenir les données
	var decodeData structure.Token

	// Decodage des données en format JSON et ajout des donnée à la variable: decodeData
	json.Unmarshal(body, &decodeData)

	// Affichage des données
	if decodeData.Error != "" {
		return decodeData
	} else {
		fmt.Println("Token récupéré avec succès : ", decodeData.AccessToken)
		return decodeData
	}
}

func GetAlbum(Token string, id string) structure.AllAlbums {
	// URL de L'API
	urlApi := fmt.Sprintf("https://api.spotify.com/v1/artists/%s/albums", id)

	// Initialisation du client HTTP qui va émettre/demander les requêtes
	httpClient := http.Client{
		Timeout: time.Second * 2, // Timeout apres 2sec
	}

	// Création de la requête HTTP vers L'API avec initialisation de la methode HTTP, la route et le corps de la requête
	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errReq.Error())
	}

	//Paramètre de type query à inserer à la req GET
	q := req.URL.Query()
	q.Add("include_groups", "album")
	req.URL.RawQuery = q.Encode()

	// Ajout d'une métadonnée dans le header
	req.Header.Add("Authorization", "Bearer "+Token)

	// Execution de la requête HTTP vars L'API
	res, errResp := httpClient.Do(req)
	if errResp != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errResp.Error())
		return structure.AllAlbums{Error: errResp.Error()}
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	// Lecture et récupération du corps de la requête HTTP
	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errBody.Error())
	}

	// Déclaration de la variable qui va contenir les données
	var decodeData structure.AllAlbums

	// Decodage des données en format JSON et ajout des donnée à la variable: decodeData
	json.Unmarshal(body, &decodeData)

	// Affichage des données
	if decodeData.Error != "" {
		return decodeData
	} else {
		fmt.Printf("album de %s récupéré : \n", id)
		println(decodeData.AlbumItems)
		return decodeData
	}
}

func GetTrack(Token string, id string) structure.Track {
	// URL de L'API
	urlApi := "https://api.spotify.com/v1/tracks/" + id

	// Initialisation du client HTTP qui va émettre/demander les requêtes
	httpClient := http.Client{
		Timeout: time.Second * 2, // Timeout apres 2sec
	}

	// Création de la requête HTTP vers L'API avec initialisation de la methode HTTP, la route et le corps de la requête
	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errReq.Error())
	}

	// Ajout d'une métadonnée dans le header
	req.Header.Add("Authorization", "Bearer "+Token)

	// Execution de la requête HTTP vars L'API
	res, errResp := httpClient.Do(req)
	if errResp != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errResp.Error())
		return structure.Track{Error: structure.Error{Message: errResp.Error()}}
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	// Lecture et récupération du corps de la requête HTTP
	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errBody.Error())
	}

	// Déclaration de la variable qui va contenir les données
	var decodeData structure.Track

	// Decodage des données en format JSON et ajout des donnée à la variable: decodeData
	json.Unmarshal(body, &decodeData)

	// Affichage des données
	if decodeData.Error.Message != "" {
		return decodeData
	} else {
		fmt.Printf("album de %s récupéré : \n", id)
		fmt.Println(decodeData)
		return decodeData
	}
}
