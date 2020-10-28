package file

import "runtime"

const (
	UnixEOL = "\n"
	DosEOL  = "\r\n"
)

func getEOL() string {
	switch runtime.GOOS {
	case "linux":
		return UnixEOL
	case "windows":
		return DosEOL
	}
	return "\n"
}
