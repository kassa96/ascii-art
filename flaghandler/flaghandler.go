package flaghandler

import (
	"fmt"
	"strings"
)

var Listflag []string

func AddAceptedFlags(flags ...string) {
	Listflag = flags
}

func HasFlagPrefix(arg string) bool {
	return strings.HasPrefix(arg, "--")
}

func GetFlagWitheValue(arg string) (string, string) {
	if !HasFlagPrefix(arg) {
		return "", ""
	}
	arg = arg[2:]
	if strings.HasSuffix(arg, "=") {
		return arg[:len(arg)-1], ""
	}
	tab := strings.Split(arg, "=")
	if len(tab) == 1 {
		return tab[0], ""
	}
	flagName, flagValue := tab[0], strings.Join(tab[1:], "")
	return flagName, flagValue
}

func IsAcceptedFlag(flagName string) bool {
	for _, flag := range Listflag {
		if flag == flagName {
			return true
		}
	}
	return false
}

func Messagefs() {
	fmt.Println("Usage: go run . [STRING] [BANNER]")
	fmt.Println()
	fmt.Println("go run . something standard")
}

func MessageColor() {
	fmt.Println("Usage: go run . [OPTION] [STRING]")
	fmt.Println()
	fmt.Println("go run . --color=<color> <letters to be colored> \"something\"")
}
func MessageOutput() {
	fmt.Println("Usage: go run . --output=<fileName> [BANNER]")
	fmt.Println()
	fmt.Println("EX: go run . --output=<fileName.txt> something standard")
}

func MessageAlign() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
	fmt.Println()
	fmt.Println("Example: go run . --align=right something standard")
}

func MessageReverse() {
	fmt.Println("Usage: go run . [OPTION]")
	fmt.Println()
	fmt.Println("EX: go run . --reverse=<fileName>")
}

func MessageEmptyFlag(flagName string) {
	if flagName == "color" {
		MessageColor()
	} else if flagName == "align" {
		MessageAlign()
	} else if flagName == "output" {
		MessageOutput()
	} else if flagName == "reverse" {
		MessageReverse()
	}
}

func CheckArgFlagColor(args []string) bool {
	if len(args) == 0 || len(args) > 2 {
		MessageColor()
		return false
	}
	return true
}

func CheckArgFlagOutput(args []string) bool {
	if len(args) == 0 || len(args) > 2 {
		MessageOutput()
		return false
	}
	return true
}

func CheckArgFlagAlign(args []string) bool {
	if len(args) == 0 || len(args) > 2 {
		MessageAlign()
		return false
	}
	return true
}

func CheckArgFlagReverse(args []string) bool {
	if len(args) > 0 {
		MessageReverse()

		return false
	}
	return true
}

func CheckValueFlagColor(value string, colorsName map[string]string) bool {
	value = strings.ToLower(value)
	_, isColor := colorsName[value]
	if !isColor {
		fmt.Println("This color is not accepted.")
	}
	return isColor
}

func CheckValueFlagAlign(value string) bool {
	value = strings.ToLower(value)
	if value == "left" || value == "right" || value == "center" || value == "justify" {
		return true
	}
	MessageAlign()
	return false
}

func CheckValueFlagOutput(fileName string) bool {
	if len(fileName) <= 4 || !strings.HasSuffix(fileName, ".txt") {
		fmt.Println("error: the fileName has not a correct format.")
		return false
	}
	return true
}

func CheckValueFlagReverse(fileName string) bool {
	if len(fileName) <= 4 || !strings.HasSuffix(fileName, ".txt") {
		fmt.Println("error: the fileName has not a correct format.")
		return false
	}
	return true
}

func IsAcceptedBannier(args []string) bool {
	bannier := strings.ToLower(args[1])
	if len(args) == 2 && (bannier == "standard" || bannier == "shadow" || bannier == "thinkertoy") {
		return true
	}
	fmt.Println("This banner name is not accepted.")
	return false
}

func CheckAllFlagWithValuAndArguments(flagName, flagValue string, args []string, colorsName map[string]string) bool {
	// Quand la valeur du flag est vide.
	if flagName == "" || !IsAcceptedFlag(flagName) {
		fmt.Println("The flag name is incorrect or doesn't exist.")
		return false
	}
	if flagValue == "" {
		MessageEmptyFlag(flagName)
		return false
	}
	if flagName == "output" && CheckValueFlagOutput(flagValue) && CheckArgFlagOutput(args) {
		return true
	} else if flagName == "color" && CheckValueFlagColor(flagValue, colorsName) && CheckArgFlagColor(args) {
		return true
	} else if flagName == "align" && CheckValueFlagAlign(flagValue) && CheckArgFlagAlign(args) {
		return true
	} else if flagName == "reverse" && CheckValueFlagReverse(flagValue) && CheckArgFlagReverse(args) {
		return true
	}
	return false
}
