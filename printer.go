package apollo

import "github.com/fatih/color"

// SuccessPrint prints a success message
func SuccessPrint(text string) {
	green := color.New(color.FgGreen, color.Bold)
	green.Println("[+] " + text)
}

// FailPrint prints a failing message
func FailPrint(text string) {
	red := color.New(color.FgRed, color.Bold)
	red.Println("[-] " + text)
}

// WarnPrint prints a warning message
func WarnPrint(text string) {
	yellow := color.New(color.FgYellow, color.Bold)
	yellow.Println("[!] " + text)
}

// InfoPrint prints a informative message
func InfoPrint(text string) {
	blue := color.New(color.FgBlue, color.Bold)
	blue.Println("[$] " + text)
}
