Présentation
ASCII ART est un programme qui permet d'afficher les caractères ASCII sous forme de représentation graphique. Nous avons trois types de bannières standard, shadow, thinkertoy.

Utilisation
=> go run . [STRING] [Banner]
Les bannières acceptées sont : standard, shadow, thinkertoy.

=> go run . n'affiche rien.

=> go run --output=<fileName> something [Banner]
Le flag --output permet d'enregistrer la représentation graphique de l'expression dans un fichier fourni comme valeur du flag output.

=> go run . --align=right something [Banner]
Le flag --align permet d'aligner la représentation graphique soit à gauche (left), ou à droite (right), ou au centre (center), ou bien de manière justifiée (justify).

=> go run . --color=<color> <letters to be colored> [something]
Le flag --color permet de colorer les lettres désignées en premier argument lors de l'affichage du texte en deuxième argument. S'il y a un seul argument, ce sera tout le texte qui sera colorié. Les couleurs de base prises en compte sont : red, green, yellow, magenta, cyan, blue. On peut les combiner avec light comme lightgreen ou avec bold comme greenbold.

=> go run . --reverse=<fileName>
Le flag --reverse permet de prendre une représentation graphique pour la décoder en texte simple. fileName représente le nom du fichier avec l'extension .txt.
Pour ajouter un fichier, il faut le metre dans le dossier examples.
Il existe 8 examples par défaut dans le dossier examples.
ex: go run . --reverse=example00.txt
