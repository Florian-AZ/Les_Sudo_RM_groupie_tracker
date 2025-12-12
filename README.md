# TP GROUPIE-TRACKER

Ce projet est une application web qui interagit avec l'API Spotify pour rechercher et afficherafficher des informations sur les artistes et les morceaux. Il est construit en utilisant Go pour le backend et HTML/CSS pour le frontend.

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

Un site internet sobre avec comme page:

Accueil : Description du projet plus bouton de redirection vers la page de recherche.

Page de recherche: Barre de recherche en haut avec en dessous 6 bulles de catégorie de musique (rock, metal, pop, etc.), si recherche alors les bulles disparaissent laissant place aux résultat.

Page Artiste: donnée de l'artiste, album, titre populaire, related artist.

Page Album: titre dans l'album

Page Titre: info titre, piste de lecture

Page Favoris: regroupe tout les titres favoris de l'utilisateur

Page Connexion : login ou inscription

Menu burger : regroupe les pages inscription + recherche + accueil + favoris



# Donnée à récupérer:
    
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
