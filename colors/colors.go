package colors

/*
* Ce package permet de colorier les lettres d'une représentation ASCII-art.
* Seuls les couleurs contenus dans  cette collection existe.
 */
import (
	"fmt"
	"main/utils"
	"regexp"
	"strings"
)

/*
* Liste des couleurs acceptés et leur équivalent en ANSI.
 */
var Colors = map[string]string{
	"black":        "\033[30m",
	"red":          "\033[31m",
	"green":        "\033[32m",
	"yellow":       "\033[33m",
	"blue":         "\033[34m",
	"magenta":      "\033[35m",
	"cyan":         "\033[36m",
	"white":        "\033[37m",
	"default":      "\033[39m",
	"gray":         "\033[90m",
	"lightred":     "\033[91m",
	"lightgreen":   "\033[92m",
	"lightyellow":  "\033[93m",
	"lightblue":    "\033[94m",
	"lightmagenta": "\033[95m",
	"lightcyan":    "\033[96m",
	"lightwhite":   "\033[97m",
	"blackbold":    "\033[30;1m",
	"redbold":      "\033[31;1m",
	"greenbold":    "\033[32;1m",
	"yellowbold":   "\033[33;1m",
	"bluebold":     "\033[34;1m",
	"magentabold":  "\033[35;1m",
	"cyanbold":     "\033[36;1m",
	"whitebold":    "\033[37;1m",
	"defaultbold":  "\033[39;1m",
	"darkred":      "\033[31;1m",
	"darkgreen":    "\033[32;1m",
	"darkyellow":   "\033[33;1m",
	"darkblue":     "\033[34;1m",
	"darkmagenta":  "\033[35;1m",
	"darkcyan":     "\033[36;1m",
	"darkwhite":    "\033[37;1m",
	"violet":       "\033[38;5;128m",
	"turquoise":    "\033[38;5;45m",
	"pink":         "\033[38;5;206m",
	"orange":       "\033[38;5;202m",
	"lightgray":    "\033[38;5;252m",
	"darkgray":     "\033[38;5;241m",
}

/*
* Cette fonction permet de faire le coloriage.
* @color: La couleur en string.
* @input: la chaine de caractere où il faut chercher les caractere en strings a colorier.
* @souschaine: les caracteres en string à colorier.
* @bannerLines: la representation du code ASCII-art du code ASCII.
 */
func AsciiArt(color string, input string, souschaine string, bannerLines []string) {
	// S'il y'a uniquement des retours à la lignes, affiche les retours à la ligne.
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
	// S'il y'a du texte avec des retours à la ligne, découpe cela en une tableau de chaine de caractere.
	inputSlice := strings.Split(input, "\\n")
	for _, inp := range inputSlice {
		if inp == "" {
			fmt.Print("\n")
			continue
		}
		// Si la souschaine est non vide, colorie toute les caracteres de souschaine.
		if len(souschaine) != 0 {
			for i := 0; i < 8; i++ {
				for _, c := range inp {
					if strings.Contains(strings.ToLower(souschaine), strings.ToLower(string(c))) {
						fmt.Printf("%s%s%s", Colors[color], bannerLines[int(c-32)*9+i], Colors["defaut"])
					} else {
						fmt.Printf("%s%s%s", Colors["default"], bannerLines[int(c-32)*9+i], Colors["defaut"])
					}
				}
				fmt.Print("\n")
			}
		} else {
			// Si la souschaine est vide, colorie toute l'entrée fournie input.
			for i := 0; i < 8; i++ {
				for _, c := range inp {
					fmt.Printf("%s%s%s", Colors[color], bannerLines[int(c-32)*9+i], Colors["defaut"])
				}
				fmt.Print("\n")
			}
		}
	}
}

/**
* Pour executer le programme si le flag est --color.
* @flagValue: la valeur du flag.
* @args: la liste des arguments.
* @return true si l'excécution s'est bien passé.
*/
func RunColor(flagValue string, args []string) bool {
	banner, err := utils.GetBannerLines("standard")
	if err != nil {
		return false
	}
	if len(args) == 2 {
		input := utils.FilterInput(args[1])
		AsciiArt(flagValue, input, args[0], banner)
	} else if len(args) == 1 {
		input := utils.FilterInput(args[0])
		AsciiArt(flagValue, input, input, banner)
	}
	return true
}
