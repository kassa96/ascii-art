package main

import (
	"main/ascii"
	"main/colors"
	"main/flaghandler"
	"main/justify"
	"main/output"
	"main/reverse"
	"os"
)

func main() {
	flaghandler.AddAceptedFlags("output", "color", "align", "reverse")
	args := os.Args[1:]
	// S'il y'a pas d'argument envoie une sortie vide.
	if len(args) == 0 {
		return
	}
	argFlag := args[0]
	// On traite les cas ou on a pas de flag.
	if !flaghandler.HasFlagPrefix(argFlag) {
		if len(args) == 1 {
			ascii.RunASCII(args)
		} else if len(args) > 2 {
			flaghandler.Messagefs()
		} else if flaghandler.IsAcceptedBannier(args) {
			ascii.RunFS(args)
		}
		return
	}
	// On traite le cas ou on a un flag fourni.
	flagName, flagValue := flaghandler.GetFlagWitheValue(argFlag)
	// On verifie si le flag, la valeur et les arguments sont correctes.
	args = args[1:]
	if flaghandler.CheckAllFlagWithValuAndArguments(flagName, flagValue, args, colors.Colors) {
		if flagName == "output" {
			output.RunOutput(flagValue, args)
		} else if flagName == "color" {
			colors.RunColor(flagValue, args)
		} else if flagName == "align" {
			justify.RunAlign(flagValue, args)
		} else if flagName == "reverse" {
			reverse.RunReverse(flagValue, args)
		}
	}
}
