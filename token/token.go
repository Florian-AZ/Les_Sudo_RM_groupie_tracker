package token

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

var (
	accesToken     string
	dateExpiration time.Time
)

func GetValidToken() string {
	//on vérifie si le token est encore valide, c'est à dire si la date actuelle est avant la date d'expiration
	if time.Now().Before(dateExpiration) && accesToken != "" {
		fmt.Printf("token.GetValidToken - Actuel - Token : %s\n\n", accesToken)
		fmt.Printf("token.GetValidToken - Actuel - Date d'expiration : %s\n\n", dateExpiration)
		return accesToken
	} else {
		T := GetToken()
		accesToken = T.AccessToken
		fmt.Printf("token.GetValidToken - Nouveau - Token : %s\n\n", accesToken)
		dateExpiration = time.Now().Add(time.Hour)
		fmt.Printf("token.GetValidToken - Nouveau - Date d'expiration : %s\n\n", dateExpiration)
		return accesToken
	}
}

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
		fmt.Printf("token.GetToken - Erreur - NewRequest : %s\n\n", errReq.Error())
	}

	// Ajout d'une métadonnée dans le header
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Execution de la requête HTTP vars L'API
	res, errResp := httpClient.Do(req)
	if errResp != nil {
		fmt.Printf("token.GetToken - Erreur - Do : %s\n\n", errResp.Error())
		return structure.Token{Error: errResp.Error()}
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	// Lecture et récupération du corps de la requête HTTP
	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		fmt.Printf("token.GetToken - Erreur - ReadAll : %s\n\n", errBody.Error())
	}

	// Déclaration de la variable qui va contenir les données
	var decodeData structure.Token

	// Decodage des données en format JSON et ajout des donnée à la variable: decodeData
	json.Unmarshal(body, &decodeData)

	// Affichage des données
	if decodeData.Error != "" {
		fmt.Printf("token.GetToken - Erreur - Token non récupéré : %s\n\n", decodeData.ErrorDescription)
		return decodeData
	} else {
		fmt.Printf("token.GetToken - Token récupéré avec succès : %s\n\n", decodeData.AccessToken)
		return decodeData
	}
}
