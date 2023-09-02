/*
* Ce package permet d'aligner la sortie du console.
 */
package justify

import (
	"fmt"
	"main/output"
	"main/utils"
	"os"
	"regexp"
	"strings"
	"syscall"
	"unsafe"
)

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

var largeurEcran = GetConsoleWidth() // Largeur de l'écran par défaut

/**
* Cette fonction prend une chaîne de caractères texte en entrée et renvoie la même chaîne de caractères sans alignement.
* Elle est utilisée lorsque le type d'alignement spécifié est "gauche" ou non spécifié.
 */
func AlignementGauche(input string, bannerLines []string) string {
	input = strings.TrimSpace(input)
	reg := regexp.MustCompile(`^(\\n)+$`)
	if reg.MatchString(input) {
		if input == "\\n" {
			return "\n"
		}
		input = input[2:]
	}
	if input == "" {
		return ""
	}
	inputAsciiArt := output.AsciiArtSave(input, bannerLines)
	lines := strings.Split(inputAsciiArt, "\n")
	ascii := ""
	for _, line := range lines {
		nbreEspace := largeurEcran - len(line)
		if nbreEspace < 0 {
			fmt.Println("The ascii art representation can't be displayed.")
			os.Exit(0)
		}
		espace := strings.Repeat(" ", nbreEspace)
		ascii += line + espace
		ascii += "\n"
	}
	return ascii
}

func GetCharacterAscii(char rune, lines []string) []string {
	var result []string
	for i := 0; i < 8; i++ {
		result = append(result, lines[int(char-32)*9+i]+"\n")
	}
	return result
}

func GetCharacterWidth(input string, lines []string) int {
	taille := 0
	words := strings.Split(input, " ")
	for _, word := range words {
		for _, char := range word {
			taille += len(GetCharacterAscii(char, lines)[0])
		}
	}
	return taille
}

/**
* Cette fonction prend une chaîne de caractères texte en entrée et renvoie une nouvelle chaîne de caractères
* où le texte est centré horizontalement par rapport à la largeur de l'écran.
 */
func AlignementCenter(input string, bannerLines []string) string {
	input = strings.TrimSpace(input)
	reg := regexp.MustCompile(`^(\\n)+$`)
	if reg.MatchString(input) {
		if input == "\\n" {
			return "\n"
		}
		input = input[2:]
	}
	if input == "" {
		return ""
	}
	inputAsciiArt := output.AsciiArtSave(input, bannerLines)
	lines := strings.Split(inputAsciiArt, "\n")
	ascii := ""
	for _, line := range lines {
		nbreEspace := largeurEcran - len(line)
		if nbreEspace < 0 {
			fmt.Println("The ascii art representation can't be displayed.")
			os.Exit(0)
		}
		espace := strings.Repeat(" ", nbreEspace/2)
		ascii += espace + line + espace
		ascii += "\n"
	}
	return ascii
}

/**
* Cette fonction prend une chaîne de caractères texte en entrée et renvoie une nouvelle chaîne de caractères
* où le texte est aligné sur la droite de l'écran.
 */
func AlignementDroite(input string, bannerLines []string) string {
	input = strings.TrimSpace(input)
	reg := regexp.MustCompile(`^(\\n)+$`)
	if reg.MatchString(input) {
		if input == "\\n" {
			return "\n"
		}
		input = input[2:]
	}
	if input == "" {
		return ""
	}
	inputAsciiArt := output.AsciiArtSave(input, bannerLines)
	lines := strings.Split(inputAsciiArt, "\n")
	ascii := ""
	for _, line := range lines {
		nbreEspace := largeurEcran - len(line)
		if nbreEspace < 0 {
			fmt.Println("The ascii art representation can't be displayed.")
			os.Exit(0)
		}
		espace := strings.Repeat(" ", nbreEspace)
		ascii += espace + line
		ascii += "\n"
	}
	return ascii
}

/*
** Pour recupérer le nombre de caractere imprimable sur la ligne d'une console.
* @return: la taille de l'écran.
 */
func GetConsoleWidth() int {
	ws := &winsize{}
	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		panic(errno)
	}

	return int(ws.Col)
}

func Justify(input string, bannerLines []string) string {
	input = strings.Trim(input, " ")
	reg := regexp.MustCompile(`^(\\n)+$`)
	if reg.MatchString(input) {
		if input == "\\n" {
			return "\n"
		}
		input = input[2:]
	}
	if input == "" {
		return ""
	}

	expressions := strings.Split(input, "\\n")
	var listExpression []string
	for _, expression := range expressions {

		//Count the number of non white space characters
		nbreCaractere := len(strings.ReplaceAll(expression, " ", ""))

		expression = strings.TrimSpace(expression)

		//Remove duplicated spaces
		for strings.Contains(expression, "  ") {
			expression = strings.ReplaceAll(expression, "  ", " ")
		}

		//Split by space to get words
		mots := strings.Split(expression, " ")
		var mots_ascii []string
		for _, mot := range mots {
			mot = strings.TrimSpace(mot)
			mots_ascii = append(mots_ascii, output.AsciiArtSave(mot, bannerLines))
		}

		inputSize := GetCharacterWidth(expression, bannerLines)
		if inputSize > largeurEcran {
			fmt.Println("Can't display the ascii art representation.")
			os.Exit(0)
		}
		nbMots := len(mots)
		nbEspaces := largeurEcran - inputSize + nbreCaractere
		espacesParMot := nbEspaces
		espacesSupplementaires := 0
		if nbMots != 1 {
			espacesParMot = nbEspaces / (nbMots - 1)
			espacesSupplementaires = nbEspaces % (nbMots - 1)
		}
		ascii := ""
		for i := 0; i < 8; i++ {
			espacesSup := espacesSupplementaires
			if len(mots) == 0 {
				continue
			}
			if len(mots) == 1 && mots[0] == "" {
				ascii += "\n"
				break
			}
			for j := 0; j < len(mots); j++ {
				lignes := strings.Split(mots_ascii[j], "\n")
				ascii += lignes[i]
				espace := ""
				if j != len(mots)-1 {
					espace = strings.Repeat(" ", espacesParMot)
				}
				ascii += espace
				if espacesSup > 0 {
					ascii += " "
					espacesSup--
				}
			}
			ascii += "\n"
		}
		listExpression = append(listExpression, ascii)
	}

	return strings.Join(listExpression, "")
}

/**
* Pour executer le programme si le flag est --align.
* @flagValue: la valeur du flag.
* @args: la liste des arguments.
* @return true si l'excécution s'est bien passé.
 */

func RunAlign(flagValue string, args []string) bool {
	bannierName := "standard"

	if len(args) == 2 {
		bannierName = args[1]
	}
	bannerLines, err := utils.GetBannerLines(bannierName)
	if err != nil {
		return false
	}
	if flagValue == "justify" {
		input := utils.FilterInput(args[0])
		ascii := Justify(input, bannerLines)
		fmt.Print(ascii)
	} else if flagValue == "left" {
		input := utils.FilterInput(args[0])
		ascii := AlignementGauche(input, bannerLines)
		fmt.Print(ascii)

	} else if flagValue == "right" {
		input := utils.FilterInput(args[0])
		ascii := AlignementDroite(input, bannerLines)
		fmt.Print(ascii)

	} else if flagValue == "center" {
		input := utils.FilterInput(args[0])
		ascii := AlignementCenter(input, bannerLines)
		fmt.Print(ascii)
	}
	return true
}
