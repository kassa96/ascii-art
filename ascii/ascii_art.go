/*
* Ce package permet d'imprimer les caracteres du code ASCII sous la forme graphique ASCII-ART.
 */
package ascii

import (
	"fmt"
	"main/utils"
	"regexp"
	"strings"
)

/*
* Cette fonction permet d'imprimer une expresion sous la forme ASCII-art.
* @input: l'entré à imprimer.
* @bannerLines: la représentation graphique de chaque caractére du code asciii sous forme de tableau de string.
 */
func AsciiArt(input string, bannerLines []string) {
	reg := regexp.MustCompile(`^(\\n)+$`)
	if reg.MatchString(input) {
		if input == "\\n" {
			fmt.Println()
			return
		}
		input = input[2:]
	}
	if input == "" {
		return
	}
	inputSlice := strings.Split(input, "\\n")
	for _, inp := range inputSlice {
		if inp == "" {
			fmt.Print("\n")
			continue
		}
		for i := 0; i < 8; i++ {
			for _, c := range inp {
				fmt.Print(bannerLines[int(c-32)*9+i])
			}
			fmt.Print("\n")
		}
	}
}

/**
* Pour executer le programme simple.
* @args: la liste des arguments.
* @return true si l'excécution s'est bien passé.
 */
func RunASCII(args []string) bool {
	input := utils.FilterInput(args[0])
	banner, err := utils.GetBannerLines("standard")
	if err != nil {
		return false
	}
	AsciiArt(input, banner)
	return true
}

/**
* Pour executer le programme avec 2 argument (argument 2: bannier).
* @args: la liste des arguments.
* @return true si l'excécution s'est bien passé.
 */
func RunFS(args []string) bool {
	input := utils.FilterInput(args[0])
	banierType := strings.ToLower(args[1])
	banner, err := utils.GetBannerLines(banierType)
	if err != nil {
		return false
	}
	AsciiArt(input, banner)
	return true
}
