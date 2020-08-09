package apollo

import (
	"encoding/json"
	"io/ioutil"
)

// ParseUsers parses a JSON file containing user definitions
// into an array of User
func ParseUsers(file string) []User {
	data, err := ioutil.ReadFile(file)
	HandleErr(err, "Couldn't read file "+file)

	var users []User
	json.Unmarshal(data, &users)

	return users
}

// AddAuthUser adds an authorized user and suppresses
// its output
func AddAuthUser(user User) {
	ShellCmd("useradd "+user.Name+" | cat &>/dev/null", "Failed to add user "+user.Name)
}

// AddAuthSudoer adds an authorized sudoer and suppresses
// its output
func AddAuthSudoer(user User) {
	if user.Sudo {
		ShellCmd("adduser "+user.Name+" sudo | cat &>/dev/null", "Failed to add "+user.Name+" to sudo group")
	}
}

// RemUnauthSudoer removes an authorized sudoer and suppresses
// its output
func RemUnauthSudoer(user User) {
	if !user.Sudo {
		ShellCmd("deluser "+user.Name+" sudo | cat &>/dev/null", "Failed to remove "+user.Name+" from sudo group")
	}
}

// UserJSONToNames returns a slice of usernames from an array
// of Users
func UserJSONToNames(users []User) []string {
	var usernames []string
	for _, user := range users {
		usernames = append(usernames, user.Name)
	}

	return usernames
}

// GetNonDef returns the non-default users on the system (UBU16 only)
func GetNonDef(img []User, def []User) []string {
	defUsernames := UserJSONToNames(def)
	imgUsernames := UserJSONToNames(img)
	sysUsernames := CleanSplitSlice(ShellCmdCapture("cat /etc/passwd | cut -f1 -d \":\"", "Failed to read /etc/passwd"), "\n")

	var nonDef []string
	for _, el := range sysUsernames {
		if !FindStrInSlice(el, defUsernames) {
			if !FindStrInSlice(el, imgUsernames) {
				nonDef = append(nonDef, el)
			}
		}
	}

	return nonDef
}

// BinFalse sets a user's shell to /bin/false
func BinFalse(user User) {
	ShellCmd("usermod -s /bin/false "+user.Name+" | cat &>/dev/null", "Failed to set "+user.Name+"'s shell to /bin/false")
}

// BinBash sets a user's shell to /bin/bash
func BinBash(user User) {
	ShellCmd("usermod -s /bin/bash "+user.Name+" | cat &>/dev/null", "Failed to set "+user.Name+"'s shell to /bin/bash")
}
