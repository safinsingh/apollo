package command

import "os/exec"

func Run(command string) error {
	cmd := exec.Command("bash", "-c", command)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func Out(command string) (string, error) {
	out, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
