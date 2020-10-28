package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func Touch(file string) error {
	if err := os.MkdirAll(filepath.Dir(file), 0o755); err != nil {
		return err
	}
	_, err := os.Create(file)
	return err
}

func Delete(file string) error {
	return os.RemoveAll(file)
}

func Config(key, val, file, delimeter, comment string) error {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	clean := cleanInternal(string(content), delimeter)
	for idx, conf := range clean {
		pair := strings.Split(conf, delimeter)
		if pair[0] == key {
			clean[idx] = strings.Join([]string{pair[0], delimeter, val}, " ")
		}
	}
	err = ioutil.WriteFile(file, []byte(strings.Join(clean, getEOL())), 0o644)
	if err != nil {
		return err
	}
	return nil
}

func cleanInternal(content, delimeter string) []string {
	ret := []string{}
	for _, line := range strings.Split(content, getEOL()) {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, delimeter) {
			ret = append(ret, trimmed)
		}
	}
	return ret
}
