package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
)

func checkExisting(path string) {
	if _, err := os.Stat(path); err != nil {
		// Exists
	} else {
		fmt.Printf("Cannot link as folder %s already exists.", path)
		inputReader := bufio.NewReader(os.Stdin)
		fmt.Println("Attempt folder removal? [Y/N]")
		inputText, _ := inputReader.ReadString('\n')
		if inputText == "Y" {
			err := os.Remove(path)
			if err != nil {
				fmt.Printf("Removal failed: %s", err)
				os.Exit(1)
			}
		} else {
			fmt.Println("Exiting.")
			os.Exit(1)
		}
	}
}

func link(gameDir string, modDir string) {
	modName, err := GetModName(modDir)
	if err != nil {
		fmt.Printf("Error retrieving mod name: %s\n", err)
	}
	var gameDirModPath string = filepath.Join(gameDir, "Mods", modName)
	checkExisting(gameDirModPath)

	interruptChannel := make(chan os.Signal, 1)
	signal.Notify(interruptChannel, os.Interrupt)

	os.Symlink(modDir, gameDirModPath)
	fmt.Printf("Successfully linked.\nUse ^C to unlink and exit.\n")

	for {
		interruptEvent := <-interruptChannel
		fmt.Printf("\nRecieved %s. Unlinking.\n", interruptEvent)
		if err := os.Remove(gameDirModPath); err != nil {
			fmt.Printf("Error while unlinking: %s\n", err)
		}
		os.Exit(1)
	}

}
