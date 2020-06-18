package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

func dataFolderName(gameDir string) string {
	const is64bit = uint64(^uintptr(0)) == ^uint64(0)
	if runtime.GOOS == "windows" {
		if is64bit {
			return filepath.Join(gameDir, "RimWorldWin64_Data", "Managed")
		} else {
			return filepath.Join(gameDir, "RimWorldWin_Data", "Managed")
		}
	} else {
		return filepath.Join(gameDir, "RimWorldLinux_Data", "Managed")
	}
}

func copyFile(src string, dest string) {
	srcFile, err := ioutil.ReadFile(src)
	destFile, err := os.Create(dest)
	if err != nil {
		fmt.Printf("Error while opening and creating fetch files: %s", err)
		os.Exit(1)
	}

	_, err = destFile.Write(srcFile)
	if err != nil {
		fmt.Printf("Error while writing to destination fetch file: %s", err)
		os.Exit(1)
	}
	destFile.Close()
}

func fetch(gameDir string, libDir string) {
	_, err := os.Stat(libDir)
	_, err = os.Stat(dataFolderName(gameDir))
	if err != nil {
		fmt.Printf("Error while validating fetch directory: %s", err)
		os.Exit(1)
	}

	managedDir, err := ioutil.ReadDir(dataFolderName(gameDir))
	if err != nil {
		fmt.Printf("Error reading managed directory: %s", err)
		os.Exit(1)
	}

	for _, file := range managedDir {
		if file.Name() == "Assembly-CSharp.dll" || file.Name() == "UnityEngine.CoreModule.dll" {
			copyFile(filepath.Join(dataFolderName(gameDir), file.Name()), filepath.Join(libDir, file.Name()))
		}
	}
}
