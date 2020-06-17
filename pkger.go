package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func packageMod(modDir string) {
	modName, err := GetModName(modDir)
	if err != nil {
		fmt.Printf("Error retrieving mod name: %s\n", err)
		os.Exit(1)
	}

	outputFile, err := os.Create(fmt.Sprintf("%s.zip", modName))
	if err != nil {
		fmt.Printf("Error while creating zip file: %s", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	writer := zip.NewWriter(outputFile)
	defer writer.Close()

	addToZip(writer, modDir, modName)

	writer.Close()

	fmt.Printf("Mod successfully packaged to %s", outputFile.Name())
}

func addToZip(writer *zip.Writer, dir string, innerPath string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("Error while reading directory for packaging process: %s", err)
		os.Exit(1)
	}

	for _, file := range files {
		if file.IsDir() {
			addToZip(writer, filepath.Join(dir, file.Name()), filepath.Join(innerPath, file.Name()))
		} else {
			data, err := ioutil.ReadFile(filepath.Join(dir, file.Name()))
			if err != nil {
				fmt.Printf("Error while reading file for packaging proccess: %s", err)
				os.Exit(1)
			}

			zipFile, err := writer.Create(filepath.Join(innerPath, file.Name()))
			if err != nil {
				fmt.Printf("Error while creating file inside zip for packaging proccess: %s", err)
				os.Exit(1)
			}

			_, err = zipFile.Write(data)
			if err != nil {
				fmt.Printf("Error while writing to file inside zip for packaging proccess: %s", err)
				os.Exit(1)
			}
		}
	}
}
