package apollo

import (
	"io/ioutil"
)

// WriteToBaseDir is a wrapper around ioutil.WriteFile
// that writes files to a baseline directory
func WriteToBaseDir(dir string, fileName string, content []byte) {
	ioutil.WriteFile(dir+fileName, content, 0o644) // FIX THIS PERM!!
}

// GetLsof captures the result of lsof -i
func GetLsof(baselineDir string) {
	lsof := ShellCmdCapture("lsof -i", "Failed to run lsof")
	WriteToBaseDir(baselineDir, "lsof", []byte(lsof))
}

// GetNetstat captures the output of netstat
func GetNetstat(baselineDir string) {
	netstat := ShellCmdCapture("netstat -plunt", "Failed to run netstat")
	WriteToBaseDir(baselineDir, "netstat", []byte(netstat))
}

// GetPasswd captures the /etc/passwd file
func GetPasswd(baselineDir string) {
	passwd, err := ioutil.ReadFile("/etc/passwd")
	HandleErr(err, "Failed to read file: /etc/passwd")

	WriteToBaseDir(baselineDir, "passwd", passwd)
}

// GetGroup captures the /etc/group file
func GetGroup(baselineDir string) {
	group, err := ioutil.ReadFile("/etc/group")
	HandleErr(err, "Failed to read file: /etc/group")

	WriteToBaseDir(baselineDir, "group", group)
}

// GetShadow captures the /etc/shadow file
func GetShadow(baselineDir string) {
	shadow, err := ioutil.ReadFile("/etc/shadow")
	HandleErr(err, "Failed to read file: /etc/shadow")

	WriteToBaseDir(baselineDir, "shadow", shadow)
}

// GetIptables captures the iptables config
func GetIptables(baselineDir string) {
	iptables := ShellCmdCapture("iptables -L", "Failed to run iptables")
	WriteToBaseDir(baselineDir, "iptables", []byte(iptables))
}
