/*
* Ce package permet de transformer un texte contenant une représentation ASCII-ART en texte simple.
* Le texte doit contenir une bonne représentation du code ASCII.
* Le format de la représentation doit etre au format standard.
 */

package reverse

import (
	"fmt"
	"main/utils"
	"strings"
)

/*
* Cette fonction permet de supprimer le signe $.
* @content: le texte sous forme de tableau de string.
* @return: texte aprés transformation sous forme de tableau de string.
 */
func filtre(content []string) []string {
	for i := 0; i < len(content); i++ {
		content[i] = strings.ReplaceAll(content[i], "$", "")
	}
	return content
}

/*
* Cette fonction permet de découper le texte contant des retours à la ligne.
* @content: le texte sous forme de tableau de string.
* @return: la liste contenant chaque expression avec retour à la ligne sous forme de tableau de string.
*
 */
func backLine(content []string) [][]string {
	var expressions [][]string
	var expression []string
	for i := 0; i < len(content); i++ {
		if content[i] != "" {
			for j := 0; j < 8; j++ {
				expression = append(expression, content[i+j])
			}
			i += 7
			expressions = append(expressions, expression)
			expression = []string{}
		} else {
			expressions = append(expressions, []string{""})
		}
	}
	return expressions
}

/*
* Cette fonction permet de recupérer la représentation ASCII-art d'un  caractére dans un mot en ASCII-art.
* @debut, @fin: l'intervale de découpe correspondant à l'indice en entier des colonnes vides qui entoure un caractére.
* @return: c'est la représentation du caractére recupéré sous la forme ASCII-art sous forme de tableau de string.
 */
func getCaractere(debut, fin int, content []string) []string {
	var caractere []string
	for i := 0; i < 8; i++ {
		caractere = append(caractere, content[i][debut:fin])
		caractere[i] = caractere[i] + " "
	}
	return caractere
}

/*
* Cette fonction permet de recupere le numéro de la colonne contenant une espace vide.
* @content: le texte sous forme de tableau de string.
* @return: le numéro de la colone en entier contenant uniquement une espace.
* S'il y'a pas de colonne contenant uniquement d'espace, la fonction retourne -1.
 */
func hasEspace(content []string) int {
	var nbreEspace int
	for i := 0; i < len(content[0]); i++ {
		nbreEspace = 0
		for j := 0; j < 8; j++ {
			if content[j][i] == ' ' {
				nbreEspace++
			}
		}
		if nbreEspace == 8 {
			return i
		}
	}
	return -1
}

/*
* Cette fonction permet de comparer une caractére donnée en représentation ASCII-art
* avec celles du banniére standard et de recupérer son index.
* @caractere: la représentation ASCII-art du caractére à comparer sous forme de tableau de string.
* @return: l'index en entier correspondant du caractere dans le baniére standard.
 */
func isSame(caractere, bannerLines []string) int {
	index := 0
	for i := 0; i < len(bannerLines); i += 9 {
		ligneIdentiques := 0
		for j := 0; j < 8; j++ {
			if caractere[j] == bannerLines[i+j] {
				ligneIdentiques++
			}
		}
		if ligneIdentiques == 8 {
			return index
		}
		index++
	}
	return -1
}

/*
* Cette fonction permet de transormer un contenu en format ASCII-art en format texte simple.
* @texte: le texte ASCII-art à transormer sous forme de tableau de string.
* @bannerLines: le fichier banniére au format standard sous forme de tableau de string.
* @return le texte sous forme de string aprés transformation.
 */
func Transform(texte, bannerLines []string) string {
	texte = filtre(texte)
	exp := ""
	nbreEspace := 0
	listeExpression := backLine(texte)
	for _, expression := range listeExpression {
		if len(expression) != 0 {
			debut := 0
			index := hasEspace(expression)
			for index != -1 {
				mot := getCaractere(debut, index, expression)
				numero := isSame(mot, bannerLines)
				if numero == -1 && nbreEspace%6 == 0 {
					str := " "
					exp += str
					nbreEspace++
				} else if numero != -1 {
					str := string(rune(' ' + numero))
					exp += str
					nbreEspace = 0
				} else {
					nbreEspace++
				}
				for i := 0; i < len(expression); i++ {
					expression[i] = expression[i][index+1:]
				}
				index = hasEspace(expression)
			}
		}
		exp += "\n"
	}
	return exp
}

/**
* Pour executer le programme si le flag est --reverse.
* @flagValue: la valeur du flag.
* @args: la liste des arguments.
* @return true si l'excécution s'est bien passé.
 */

func RunReverse(flagValue string, args []string) bool {
	bannerLines, err := utils.GetBannerLines("standard")
	if err != nil {
		return false
	}
	texte, err := utils.GetBannerLinesForReverse(flagValue)
	if err != nil {
		return false
	}
	transformation := Transform(texte, bannerLines)
	fmt.Print(transformation)
	return true
}
