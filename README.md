# Pokédex en Go

Un petit site web fait en Go qui affiche tous les Pokémon en français. On peut voir la liste, ouvrir la fiche de chaque Pokémon et chercher un Pokémon par son nom.


## Description

Le site récupère ses données sur Tyradex, une API gratuite qui donne des informations sur les Pokémons. Mon programme Go envoie des requêtes à cette API, lit la réponse, puis affiche le résultat dans la page en HTML.

Ce que le site permet de faire :

- voir la liste de tous les Pokémon avec leur image
- chercher un Pokémon par son nom
- voir une page d'erreur claire quand il y a un problème

## Lancement

1. Ouvrir un terminal dans le dossier du projet
2. Lancer le serveur avec la commande : go run .
3. Ouvrir le navigateur à l'adresse : http://localhost:8080
4. Si il met trop de temps a repondre mettre sur un navigateur l'adresse suivante : 
http://localhost:8080/pokemon 

Pour arrêter la simulation faire Ctrl + c  sur le terminal
