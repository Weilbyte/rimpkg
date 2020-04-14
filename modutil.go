package main

import (
	"errors"
	"github.com/antchfx/xmlquery"
	"os"
	"path"
	"regexp"
)

func removeIllegal(modName string) string {
	illegalRegex := regexp.MustCompile("[\\\\\\\\/:*?\\\"<>|]")
	modName = illegalRegex.ReplaceAllString(modName, "")
	return modName
}

func getModXML(modDir string) (*xmlquery.Node, error) {
	file, err := os.Open(path.Join(modDir, "About", "About.xml"))
	if err != nil {
		if os.IsNotExist(err) {
			err = errors.New("About/About.xml doesn't exist")
		}
	} else {
		fileXML, err := xmlquery.Parse(file)
		return fileXML, err
	}
	return nil, err
}

func GetModName(modDir string) (string, error) {
	modXML, err := getModXML(modDir)
	if err != nil {
		return "", err
	}
	modName := xmlquery.FindOne(modXML, "//ModMetaData/name")
	if modName != nil {
		return removeIllegal(modName.InnerText()), err
	}
	return "", err
}
