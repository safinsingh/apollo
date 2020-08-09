package apollo

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

// HandleErr is a global error handler for apollo
func HandleErr(err error, customErrCmt string) {
	if err != nil {
		log.Fatalln(errors.Wrap(err, customErrCmt))
	}
}

// FileExists returns a boolean containing whether
// the file exists or not
func FileExists(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}

// ReadFileNoCmt reads a file and eliminates empty or
// commented out lines and returns a slice with the lines
func ReadFileNoCmt(file string, cmt string) []string {
	if FileExists(file) {
		data, err := ioutil.ReadFile(file)
		HandleErr(err, "Couldn't read file "+file)
		lines := CleanSplitSlice(string(data), "\n")
		noCmt := []string{}

		for _, el := range lines {
			if !strings.HasPrefix(el, cmt) && el != "" {
				noCmt = append(noCmt, el)
			}
		}

		return noCmt
	}

	HandleErr(errors.New("exit status 1"), "file doesn't exist: "+file)
	return nil
}

// ShellCmd executes a bash command
func ShellCmd(command string, customErr string) {
	cmd := exec.Command("bash", "-c", command)
	if err := cmd.Run(); err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			HandleErr(err, customErr)
		}
	}
}

// ShellCmdCapture executes a bash command and captures its output
func ShellCmdCapture(command string, customErr string) string {
	out, err := exec.Command("bash", "-c", command).Output()
	HandleErr(err, customErr)

	return string(out)
}

// FindStrInSlice returns whether any element in a slice
// is equal to a string
func FindStrInSlice(str string, slice []string) bool {
	for _, el := range slice {
		if el == str {
			return true
		}
	}
	return false
}

// RootCheck checks if an executable is being run as root
func RootCheck() {
	if os.Geteuid() != 0 {
		HandleErr(errors.New("exit status 1"), "Please run as root")
	}
}

// CleanSplitSlice returns a cleaned up (no empty element) slice
// from a string and a delimiter
func CleanSplitSlice(toSlice string, toSplitAt string) []string {
	dirtySlice := strings.Split(toSlice, toSplitAt)
	var cleanSlice []string
	for _, el := range dirtySlice {
		if el != "" {
			cleanSlice = append(cleanSlice, el)
		}
	}

	return cleanSlice
}

// GetFileName returns the name of a file from its absolute path
func GetFileName(fileName string) string {
	fileNameSlice := CleanSplitSlice(fileName, "/")
	return fileNameSlice[len(fileNameSlice)-1]
}
