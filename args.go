package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type optionStruct struct {
	gameDir string
	modDir  string
	link    bool
	pkg     bool
}

func currentDirectory() string {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(fmt.Sprintf("Error getting current directory: %s", err))
	}
	return currentDir
}

func absolutize(path string) string {
	absolutePath, _ := filepath.Abs(path)
	return absolutePath
}

func validateOptions(options optionStruct) error {
	var error error = nil
	if (options.link || options.pkg) && (options.gameDir != "") {
		if _, err := os.Stat(filepath.Join(options.gameDir, "Mods")); err != nil {
			if os.IsNotExist(err) {
				error = errors.New("Game directory path is incorrect")
			}
		}
		modName, err := GetModName(options.modDir)
		if err != nil {
			error = err
		} else if len(modName) == 0 {
			error = errors.New("Mod About.xml incomplete (missing 'name')")
		}
	} else {
		error = errors.New("Missing required parameters")
	}
	return error
}

func GetOptions() optionStruct {
	gameDirPtr := flag.String("gameDir", "", "Path to the Rimworld game directory")
	modDirPtr := flag.String("modDir", currentDirectory(), "Path to the mod directory")
	linkPtr := flag.Bool("link", false, "Link mod directory to game directory")
	pkgPtr := flag.Bool("pkg", false, "Package mod into archive")
	flag.Parse()
	options := optionStruct{
		gameDir: *gameDirPtr,
		modDir:  *modDirPtr,
		link:    *linkPtr,
		pkg:     *pkgPtr,
	}
	options.gameDir = absolutize(options.gameDir)
	options.modDir = absolutize(options.modDir)

	if err := validateOptions(options); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	return options
}
