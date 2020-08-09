package apollo

// GetManual returns a slice with the manually installed packages
// on the system (assuming /var/log/installer/initial-status.gz
// is not compromised)
func GetManual() []string {
	manual := ShellCmdCapture("comm -23 <(apt-mark showmanual | sort -u) <(gzip -dc /var/log/installer/initial-status.gz | sed -n 's/^Package: //p' | sort -u)", "Failed to parse initial-status.gz")
	return CleanSplitSlice(manual, "\n")
}

// AddPkgs loops through a slice of packages and installs them
func AddPkgs(pkgs []string) {
	for _, el := range pkgs {
		ShellCmd("apt install -y "+el+" | cat", "Installing package "+el+" failed.") // ERR HANDLING HERE == JANKY
	}
}

// RemovePkgs loops through a slice of packages and purges them
func RemovePkgs(pkgs []string) {
	for _, el := range pkgs {
		ShellCmd("apt purge -y "+el+" | cat", "Removing package "+el+" failed.") // JANKY
	}
}
