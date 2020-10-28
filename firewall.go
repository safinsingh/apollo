package apollo

import (
	"apollo/command"
	"runtime"
)

func Firewall() error {
	switch runtime.GOOS {
	case "linux":
		return command.Run("ufw enable")
	case "windows":
		return command.Run("netsh advfirewall set allprofiles state on")
	}
	return nil
}
