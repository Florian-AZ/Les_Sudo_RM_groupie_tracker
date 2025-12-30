package data

import "Groupie_Tracker/structure"

// InitWebData initialise les données de la page web
func InitWebData() *structure.PageData {
	return &structure.PageData{
		LogIn: false,
	}
}

//Restructuration des données obtenues via l'API Spotify en une structure plus lisible pour le template HTML

// TemplateHTMLSearch restructure les données de la recherche structure plus lisible
// mettre la fonction ici
