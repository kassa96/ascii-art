package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

var BannierDirectory = "./banners/"

/*
* Pour écrire dans un fichier.
* @texte: le texte à enregistrer.
* @destination: chemin d'accés du fichier.
* @return: return true si le fichier a été bien crée et  que le contenu a été bien enregistrer.
 */
func WriteTexte(texte, destination string) bool {
	if destination == "" {
		return false
	}
	err := ioutil.WriteFile(destination, []byte(texte), 0777)
	return err == nil
}

/*
* Pour supprimer les caractéres qui ne se trouvent pas dans le tableau d'encodage ASCII.
* @texte: contenu sur lequel on faut supprimer les caractéres non ASCII.
* @return: texte apres netoiyage.
 */
func FilterInput(texte string) string {
	result := ""
	for _, char := range texte {
		if char >= 32 && char <= 126 {
			result += string(char)
		}
	}
	return result
}

/*
* Pour lire le contenu d'un fichier.
* @temp: Le chemin d'acces du fichier.
* @return: le contenu sous forme de tableau de string.
 */
func GetBannerLines(fileName string) ([]string, error) {
	path := BannierDirectory + fileName + ".txt"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("fichier bannier inexistant")
		return []string{}, err
	}
	scanner := bufio.NewScanner(file)
	bannerLines := []string{}
	for scanner.Scan() {
		bannerLines = append(bannerLines, scanner.Text())
	}
	return bannerLines, nil
}

func GetBannerLinesForReverse(fileName string) ([]string, error) {
	fileName = "./examples/" + fileName
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("fichier bannier inexistant")
		return []string{}, err
	}
	scanner := bufio.NewScanner(file)
	bannerLines := []string{}
	for scanner.Scan() {
		bannerLines = append(bannerLines, scanner.Text())
	}
	return bannerLines, nil
}
