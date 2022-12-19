package main

import (
	"errors"
	"fred/internal"
	"os"
	"path/filepath"
	"strings"
)

var configFileName string = ".fredrc"

func getSourceUrls(out internal.Printer) []string {
	file, err := getConfigFile()
	if err != nil {
		out.Error(err, "unable to find rc file")
		return []string{}
	}

	cnt, err := os.ReadFile(file)
	if err != nil {
		out.Error(err, "unable to read rc file %q", file)
		return []string{}
	}

	raw := strings.Split(string(cnt), "\n")
	urls := make([]string, 0, len(raw))
	for _, url := range raw {
		url = strings.TrimSpace(url)
		if "" == url {
			continue
		}
		urls = append(urls, url)
	}
	return urls
}

func getConfigDir() (string, error) {
	if dir, err := os.UserConfigDir(); err == nil { // If all is well and we have standard config dir
		return dir, nil
	}

	if dir, err := os.UserHomeDir(); err == nil { // If all is well and we have standard config dir
		return dir, nil
	}

	return "", errors.New("unable to determine global config directory")
}

func getConfigFile() (string, error) {
	dir, err := getConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, configFileName), nil
}
