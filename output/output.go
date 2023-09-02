/*
* Ce package permet d'enregistrer dans un fichier  les caracteres du code ASCII sous la forme graphique ASCII-ART.
 */
package output

import (
	"main/flaghandler"
	"main/utils"
	"regexp"
	"strings"
)

/*
 * Cette fonction permet de recuperer sous forme de string une expresion qui est sous la forme ASCII-art.
 * @input: l'entré à imprimer.
 * @bannerLines: la représentation graphique de chaque caractére du code asciii sous forme de tableau de string.
 */
 func AsciiArtSave(input string, bannerLines []string) string {
	texte := ""
	 reg := regexp.MustCompile(`^(\\n)+$`)
	 if reg.MatchString(input) {
		 if input == "\\n" {
			 texte = "\n"
			 return texte
		 }
		 input = input[2:]
	 }
	 if input == "" {
		 return ""
	 }
	 inputSlice := strings.Split(input, "\\n")
	 for _, inp := range inputSlice {
		 if inp == "" {
			texte += "\n"
			 continue
		 }
		 for i := 0; i < 8; i++ {
			 for _, c := range inp {
				 texte += bannerLines[int(c-32)*9+i]
			 }
			 texte += "\n"
			}
	 }
	 return texte
 }

 /**
* Pour executer le programme si le flag est --output.
* @flagValue: la valeur du flag.
* @args: la liste des arguments.
* @return true si l'excécution s'est bien passé.
 */
 func RunOutput(flagValue string, args []string) {
	input := utils.FilterInput(args[0])
	bannierName := "standard"
	if len(args) == 2 {
		if !flaghandler.IsAcceptedBannier(args) {
			return
		}
		bannierName = args[1]
	}
	banner, err := utils.GetBannerLines(bannierName)
	if err != nil {
		return
	}
	ascii := AsciiArtSave(input, banner)
	ascii += "\n"
	utils.WriteTexte(ascii, flagValue)
}

 