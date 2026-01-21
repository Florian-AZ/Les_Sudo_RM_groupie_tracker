# TP GROUPIE-TRACKER

Sudo FM est un projet de création d'application web qui interagit avec l'API Spotify pour rechercher et afficher des informations sur les artistes, albums et les titres. Il est construit en utilisant Go pour le backend et HTML/CSS pour le frontend. 

L'application permet aux utilisateurs de créer un compte local et de s'y connecter, de rechercher des artistes, de visualiser leurs albums et titres, et de sauvegarder leurs morceaux préférés dans une liste de favoris.

# Lancement de l'application

Cloner le dépôt :

```bash
git clone https://github.com/votre-utilisateur/TP-API-Spotify.git
cd TP-API-Spotify/site\ web
```
Dans le répertoire `site web`, exécutez l'application Go dans le terminal :

```bash
go run main.go
```
>L'application sera accessible à l'adresse suivante : `http://localhost:8080`

# Maquettes
>Vous pouvez retrouver les maquettes dans le dossier "maquette" du projet.
        
# Tâches par ordre chronologique
- Realisation des maquettes
- Recherche des endpoint de l'API Spotify à utiliser
- Test des endpoints via un logiciel de test d'API
- Création de la fonction backend pour les token API Spotify
- Création de la page d'accueil et de son backend Go
- Ajout de la fonctionnalité de recherche et de sa page
- Ajout de la fonctionnalité d'affichage des artistes et de sa page
- Ajout de la page de gestion des erreurs API/Go avec son backend
- Ajout de la fonctionnalité d'affichage des albums et de sa page
- Ajout de la fonctionnalité d'affichage des titres et de sa page
- Création de la fonctionnalité de connexion et d'inscription
- Ajout de la fonctionnalité de gestion des favoris et de sa page
- Ajout de la page a propos

# Difficultés rencontrées

Lors de la création de l'application, nous avons rencontré plusieurs défis techniques, notamment la compréhension complète de l'APi Spotify et les nuances des balises Golang dans les templates HTML. De plus nous avons dû comprendre comment fonctionne exactement les requêtes HTTP en Go pour pouvoir interagir efficacement avec l'API Spotify. La gestion de l'authentification des utilisateurs locaux et la sauvegarde de favoris avec le JSON ont également été des aspects complexes à mettre en place. 

# Donnée à récupérer via l'API Spotify:
    
- https://api.spotify.com/v1/search
    tracks
        items[]
            album
                external_urls
                    spotify
                id
                images[1]
                    url
                name
                release_date
                total_tracks
            artists[]
                external_urls
                    spotify
                id
                name
            duration_ms
            external_urls
                spotify
            id
            name
    artists
        items[]
            external_urls
                spotify
            folowers
                total
            genres[]
            id
            images[1]
                url
            name
    albums
        items[]
            total_tracks
            external_urls
                spotify
            id
            images[1]
                url
            name
            release_date
            artists[]
                external_urls
                    spotify
                id
                name
- https://api.spotify.com/v1/artists/{id}
    external_urls
        spotify
    followers
        total
    genres[]
    id
    images[1]
        url
    name
- https://api.spotify.com/v1/artists/{id}/top-tracks
    tracks[]
        album
            external_urls
                spotify
            id
            images[1]
                url
            name
            release_date
            total_tracks
        artists[]
            external_urls
                spotify
            id
            name
        duration_ms
        external_urls
            spotify
        id
        name
- https://api.spotify.com/v1/artists/{id}/albums
    items[]
        total_tracks
        external_urls
            spotify
        id
        images[1]
            url
        name
        release_date
        artists[]
            external_urls
                spotify
            id
            name
- https://api.spotify.com/v1/albums/{album id}/tracks
    items[]
        artists[]
            external_urls
                spotify
            id
            name
        duration_ms
        external_urls
            spotify
        id
        name
- https://api.spotify.com/v1/tracks/{track id}
    album
        external_urls
            spotify
        id
        images[1]
            url
        name
        release_date
        total_tracks
    artists[]
        external_urls
            spotify
        id
        name
    duration_ms
    external_urls
        spotify
    id
    name
- https://accounts.spotify.com/api/token
    access_token

# Auteur:

* Florian AZRIA
>Developpeur CSS/HTML

* Emrick RIVET
>Developpeur Golang/HTML


