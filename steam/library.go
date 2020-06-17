package steam

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/andygrunwald/vdf"
)

func findSteamLibraries(steamPath string) []string {
	libraryListPath := filepath.Join(steamPath, "steamapps", "libraryfolders.vdf")
	libraryListFile, err := os.Open(libraryListPath)

	libraryListParser := vdf.NewParser(libraryListFile)
	libraryListParsed, err := libraryListParser.Parse()

	if err != nil {
		return nil
	}

	var libraryPaths []string

	libraryPaths = append(libraryPaths, filepath.Join(steamPath, "steamapps", "common"))

	for _, root := range libraryListParsed {
		for key, value := range root.(map[string]interface{}) {
			if key == "TimeNextStatsReport" || key == "ContentStatsID" {
				continue
			}
			libraryPaths = append(libraryPaths, filepath.Join(value.(string), "steamapps", "common"))
		}
	}

	return libraryPaths
}

func findRimWorldDirectory(libraries []string) string {
	for library := range libraries {
		files, _ := ioutil.ReadDir(libraries[library])
		for _, file := range files {
			if strings.ToLower(file.Name()) == "rimworld" {
				return filepath.Join(libraries[library], file.Name())
			}
		}
	}
	return ""
}
