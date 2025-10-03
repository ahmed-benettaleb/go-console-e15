# go-console-e15 — Application console météo (documentation pour débutants)

Bienvenue dans la documentation du projet `go-console-e15`. Ce dépôt contient une petite application console écrite en Go qui :

- récupère la météo d'une ville via l'API OpenWeatherMap et l'affiche en console ;
- contient un module d'exemple pour se connecter à une base MongoDB et effectuer des opérations simples (ajout d'utilisateur, login).

Cette documentation explique la structure du projet, comment démarrer l'environnement (PowerShell sous Windows), bonnes pratiques, et exercices pour débutants.

## Arborescence principale

- `main.go` : point d'entrée de l'application. Gère la lecture utilisateur, l'appel à l'API météo et quelques appels de test vers MongoDB.
- `conexion.go` : fonctions pour se connecter à MongoDB et manipuler une collection `users` (connexion, ajout, login).
- `docker-compose.yml` : configuration Docker pour lancer MongoDB en local (utile en développement).
- `go.mod` / `go.sum` : modules et dépendances Go.

## Prérequis

- Go (version 1.25.x recommandée) installé sur votre machine.
- Docker & Docker Compose (pour lancer la base MongoDB en local). Si vous ne souhaitez pas utiliser Docker, un serveur MongoDB accessible fera l'affaire.
- Une clé API OpenWeatherMap (vous pouvez créer un compte gratuit sur https://openweathermap.org/).

## Comment ça marche (en bref)

1. L'utilisateur saisit un nom de ville dans la console.
2. L'application construit une requête HTTP vers OpenWeatherMap (avec la clé API) et récupère la réponse JSON.
3. Le JSON est parsé en structures Go (`WeatherData`) puis affiché (température, humidité, description en français).
4. En complément, le projet contient des fonctions pour se connecter à MongoDB et tester des opérations CRUD basiques.

## Variables d'environnement

- `OPENWEATHER_API_KEY` : (recommandé) clé API OpenWeatherMap. Si non fournie, le code contient une clé de secours (présente uniquement pour faciliter l'apprentissage — à retirer en production).

## Démarrage (PowerShell)

1. Démarrer MongoDB (optionnel, si vous voulez tester la partie Mongo) :

```powershell
docker-compose up -d
```

2. Définir la clé OpenWeather (exemple) :

```powershell
$env:OPENWEATHER_API_KEY = "VOTRE_CLE_OPENWEATHER"
```

3. Lancer l'application (dans le dossier qui contient `main.go`) :

```powershell
go run .
```

Ou compiler puis exécuter :

```powershell
go build -o weatherapp.exe
.\weatherapp.exe
```

4. Utilisation : tapez le nom d'une ville (ex. `Paris`) et appuyez sur Entrée. Tapez `q`, `quit` ou `exit` pour quitter.

## Détails techniques (pour apprendre)

- Requêtes HTTP : package standard `net/http` est utilisé pour faire un `GET` vers l'API.
- JSON : `encoding/json` pour désérialiser la réponse dans la struct `WeatherData`.
- MongoDB : driver officiel `go.mongodb.org/mongo-driver` pour la connexion et opérations CRUD.
- Contexte (`context`) : utilisé pour imposer des timeouts lors de la connexion Mongo.

### Points à lire dans le code

- `getWeather` dans `main.go` : construction de l'URL, requête HTTP, gestion des codes d'erreur, parsing JSON, formatage de la sortie.
- `ConnectMongoDB`, `AddUser`, `Login` dans `conexion.go` : création du client, `Ping`, `InsertOne`, `FindOne`.

## Bonnes pratiques et sécurité (à appliquer)

- Ne jamais commiter de clés API ou mots de passe dans le dépôt.
- Stocker les mots de passe hachés (utiliser `golang.org/x/crypto/bcrypt`) — le code actuel stocke les mots de passe en clair (à corriger).
- Eviter `log.Fatalf` pour les erreurs non-fatales : préférez retourner l'erreur et laisser le caller décider.
- Gérer graceful shutdown : appeler `client.Disconnect(ctx)` pour fermer proprement la connexion Mongo.

## Exercices guidés pour débutants (progression)

1. Exécuter le projet localement et tester plusieurs villes.
2. Retirer la clé codée en dur et forcer l'utilisation de `OPENWEATHER_API_KEY`.
3. Implémenter le hachage des mots de passe dans `AddUser` et la comparaison dans `Login` avec `bcrypt`.
4. Ajouter une commande `history` qui enregistre chaque recherche météo en MongoDB et affiche l'historique.
5. Écrire un test unitaire pour `getWeather` en simulant la réponse HTTP (abstraire `http.Get` ou utiliser un serveur test).

## Dépannage rapide

- Erreur réseau / API : vérifiez votre clé API et la connectivité réseau.
- Erreur Mongo : vérifiez que Docker est lancé et que le conteneur expose le port 27017. `docker ps` et `docker-compose logs mongodb` peuvent aider.
- Build Go : `go build` affichera les erreurs de compilation.

## Commandes utiles (récapitulatif)

```powershell
docker-compose up -d                 # Démarrer MongoDB
$env:OPENWEATHER_API_KEY = "..."   # Définir la clé API en PowerShell
go run .                             # Lancer l'application
go build -o weatherapp.exe           # Compiler
docker-compose logs mongodb          # Voir les logs du conteneur Mongo
```

## Prochaines améliorations recommandées

- Remplacer stockage en clair par bcrypt pour les mots de passe.
- Factoriser la configuration (clé API, Mongo URI) dans un fichier `.env` pour le dev et un système de configuration pour la prod.
- Ajouter des tests unitaires et d'intégration minimalistes.
- Ajouter des messages d'erreur plus parlants et une gestion d'erreurs non-fatale.

---

Si vous voulez, je peux appliquer automatiquement l'une de ces améliorations :

- implémenter `bcrypt` pour `AddUser` et `Login` (+ tests) ;
- corriger le flux d'exécution dans `main.go` pour connecter Mongo avant la boucle et fermer la connexion proprement ;
- ajouter la fonctionnalité `history` en Mongo.

Indiquez quelle amélioration vous voulez que j'implémente ensuite et je m'en occupe.