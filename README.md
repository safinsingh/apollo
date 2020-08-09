<p align="center">
  <h1>apollo</h1>
</p>

<h3 align="center">
  <i>
    🤖 An automated Linux system administration suite
  </i>
</h3>

<p align="center">
  <img alt="Version" src="https://img.shields.io/badge/Version-1.0-red.svg" />
  <img alt="License: MIT" src="https://img.shields.io/badge/License-GPLv3-orange.svg" />
  <img alt="Made with Go" src="https://img.shields.io/badge/Made%20with-Go-yellow.svg" />
  <img alt="A+" src="https://goreportcard.com/badge/github.com/safinsingh/apollo">
  <img alt="PRs Welcome" src="https://img.shields.io/badge/PRs-Welcome-blue.svg">
  <img alt="awesome" src="https://img.shields.io/badge/Awesome-Yes-blueviolet">
  <br>
  <img alt="gh-actions" src="https://github.com/safinsingh/apollo/workflows/Go/badge.svg" />
</p>

<hr>

```sh
go get github.com/safinsingh/apollo
```

## 🔮 Usage

```go
import "github.com/safinsingh/apollo"
```

#### func AddAuthSudoer

```go
func AddAuthSudoer(user User)
```

AddAuthSudoer adds an authorized sudoer and supresses its output

#### func AddAuthUser

```go
func AddAuthUser(user User)
```

AddAuthUser adds an authorized user and supresses its output

#### func AddPkgs

```go
func AddPkgs(pkgs []string)
```

AddPkgs loops through a slice of packages and installs them

#### func BinBash

```go
func BinBash(user User)
```

BinBash sets a user's shell to /bin/bash

#### func BinFalse

```go
func BinFalse(user User)
```

BinFalse sets a user's shell to /bin/false

#### func CleanSplitSlice

```go
func CleanSplitSlice(toSlice string, toSplitAt string) []string
```

CleanSplitSlice returns a cleaned up (no empty element) slice from a string and
a delimiter

#### func Config

```go
func Config(key string, val []string, file string, delim string, comment string, force bool)
```

Config configures a certain policy in a given file

#### func CopyConfig

```go
func CopyConfig(file1 string, file2 string)
```

CopyConfig copies the content of a file and writes it to another

#### func ExpandSliceAtDelim

```go
func ExpandSliceAtDelim(slice []string, delim string) string
```

ExpandSliceAtDelim is a utility function used to expands a slice to a string
with a certain delimiter

#### func FailPrint

```go
func FailPrint(text string)
```

FailPrint prints a failing message

#### func FileExists

```go
func FileExists(file string) bool
```

FileExists returns a boolean containing whether the file exists or not

#### func FindStrInSlice

```go
func FindStrInSlice(str string, slice []string) bool
```

FindStrInSlice returns whether any element in a slice is equal to a string

#### func GetFileName

```go
func GetFileName(fileName string) string
```

GetFileName returns the name of a file from its absolute path

#### func GetGroup

```go
func GetGroup(baselineDir string)
```

GetGroup captures the /etc/group file

#### func GetIptables

```go
func GetIptables(baselineDir string)
```

GetIptables captures the iptables config

#### func GetLsof

```go
func GetLsof(baselineDir string)
```

GetLsof captures the result of lsof -i

#### func GetManual

```go
func GetManual() []string
```

GetManual returns a slice with the manually installed packages on the system
(assuming /var/log/installer/initial-status.gz is not compromised)

#### func GetNetstat

```go
func GetNetstat(baselineDir string)
```

GetNetstat captures the output of netstat

#### func GetNonDef

```go
func GetNonDef(img []User, def []User) []string
```

GetNonDef returns the non-default users on the system (UBU16 only)

#### func GetPasswd

```go
func GetPasswd(baselineDir string)
```

GetPasswd captures the /etc/passwd file

#### func GetShadow

```go
func GetShadow(baselineDir string)
```

GetShadow captures the /etc/shadow file

#### func HandleErr

```go
func HandleErr(err error, customErrCmt string)
```

HandleErr is a global error handler for apollo

#### func InfoPrint

```go
func InfoPrint(text string)
```

InfoPrint prints a informative message

#### func ReadFileNoCmt

```go
func ReadFileNoCmt(file string, cmt string) []string
```

ReadFileNoCmt reads a file and eliminates empty or commented out lines and
returns a slice with the lines

#### func RemUnauthSudoer

```go
func RemUnauthSudoer(user User)
```

RemUnauthSudoer removes an authorized sudoer and supresses its output

#### func RemovePkgs

```go
func RemovePkgs(pkgs []string)
```

RemovePkgs loops through a slice of packages and purges them

#### func RootCheck

```go
func RootCheck()
```

RootCheck checks if an executable is being run as root

#### func SafeCopyConfig

```go
func SafeCopyConfig(file1 string, file2 string, f1delim string, f1comment string, f2delim string, f2comment string)
```

SafeCopyConfig safely merges two configuration files with Config

#### func ShellCmd

```go
func ShellCmd(command string, customErr string)
```

ShellCmd executes a bash command

#### func ShellCmdCapture

```go
func ShellCmdCapture(command string, customErr string) string
```

ShellCmdCapture executes a bash command and captures its output

#### func SuccessPrint

```go
func SuccessPrint(text string)
```

SuccessPrint prints a success message

#### func UserJSONToNames

```go
func UserJSONToNames(users []User) []string
```

UserJSONToNames returns a slice of usernames from an array of Users

#### func WarnPrint

```go
func WarnPrint(text string)
```

WarnPrint prints a warning message

#### func WriteToBaseDir

```go
func WriteToBaseDir(dir string, fileName string, content []byte)
```

WriteToBaseDir is a wrapper around ioutil.WriteFile that writes files to a
baseline directory

#### type Service

```go
type Service struct {
	Name     string `json:"name"`
	Critical bool   `json:"critical"`
}
```

Service definition for the JSON unmarshaler

#### type User

```go
type User struct {
	Name string `json:"name"`
	Sudo bool   `json:"sudo"`
}
```

User definition for the JSON unmarshaler

#### func ParseUsers

```go
func ParseUsers(file string) []User
```

ParseUsers parses a JSON file containing user definitions into an array of User

## 👨‍💻 Author

Linkedin: [Safin Singh](https://www.linkedin.com/in/safin-singh-b2630918a/) <br>
GitHub: [safinsingh](https://github.com/safinsingh) <br>
Dribbble: [Safin Singh](https://dribbble.com/safinsingh/) <br>
YouTube: [Safin Singh](https://www.youtube.com/channel/UCvb01sUdAgcPAG1j0SLxAtA)

## 🤝 Contributing

Contributions, PRs, issues and feature requests are welcome! Feel free to check out my [issues page](https://github.com/safinsingh/midnight/issues).

## ❤️ Show your support

Give a ⭐️ if this project helped you!
Hope you enjoy it!
