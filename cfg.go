package apollo

import (
	"io/ioutil"
)

// Config configures a certain policy in a given file
func Config(key string, val []string, file string, delim string, comment string, force bool) {
	noCmt := ReadFileNoCmt(file, comment)
	keyPair := key + GetFileName(file) + ExpandSliceAtDelim(val, delim)

	for i, el := range noCmt {
		if CleanSplitSlice(el, comment)[0] == key {
			if el != keyPair {
				if force {
					WarnPrint("File already contains key: " + key + ". Overwriting...")
					noCmt[i] = keyPair
				} else {
					WarnPrint("Force not specified. Skipping config " + key)
				}
			}
		}
	}

	if !FindStrInSlice(keyPair, noCmt) {
		WarnPrint("File does not contain key pair: " + key + ". Adding...")
		noCmt = append(noCmt, keyPair)
	}

	updatedFile := ""
	for _, el := range noCmt {
		updatedFile += el + "\n"
	}

	ioutil.WriteFile(file, []byte(updatedFile), 0o644) // FIX THAT PERM!
}

// CopyConfig copies the content of a file and writes it to another
func CopyConfig(file1 string, file2 string) {
	data, err := ioutil.ReadFile(file1)
	HandleErr(err, "Failed to read file "+file1)

	err2 := ioutil.WriteFile(file2, data, 0o644) // FIX THIS PERM!!
	HandleErr(err2, "Failed to write file "+file2)
}

// SafeCopyConfig safely merges two configuration files with Config
func SafeCopyConfig(file1 string, file2 string, f1delim string, f1comment string, f2delim string, f2comment string) {
	if FileExists(file1) {
		if !FileExists(file2) {
			WarnPrint("File " + file2 + " doesn't exist. Writing secure file...")

			file1Data, err := ioutil.ReadFile(file1)
			HandleErr(err, "Couldn't read file "+file1)

			err2 := ioutil.WriteFile(file2, file1Data, 0o644) // FIX THIS PERM
			HandleErr(err2, "Couldn't write file"+file2)
		} else {
			noCmtFile1 := ReadFileNoCmt(file1, f1comment)

			for _, el := range noCmtFile1 {
				key := CleanSplitSlice(el, f1delim)[0]
				val := CleanSplitSlice(el, f1delim)[1:]

				Config(key, val, file2, f2delim, f2comment, true)
			}
		}
	}
}

// ExpandSliceAtDelim is a utility function used to
// expands a slice to a string with a certain delimiter
func ExpandSliceAtDelim(slice []string, delim string) string {
	var expanded string
	for i, el := range slice {
		if i != len(slice)-1 {
			expanded += el + delim
		} else {
			expanded += el
		}
	}
	return expanded
}
