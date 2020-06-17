package steam

import (
	"os"
	"os/user"
	"path/filepath"
	"runtime"

	"golang.org/x/sys/windows/registry"
)

func findSteam() string {
	if runtime.GOOS == "windows" {
		key, err := registry.OpenKey(registry.LOCAL_MACHINE, "SOFTWARE\\WOW6432Node\\Valve\\Steam", registry.QUERY_VALUE)
		installPath, _, err := key.GetStringValue("InstallPath")
		if err != nil {
			return ""
		}

		return installPath
	} else if runtime.GOOS == "linux" {
		user, err := user.Current()
		steamPath := filepath.Join(user.HomeDir, ".local", "share", "Steam")
		_, err = os.Stat(filepath.Join(steamPath, "steamapps", "libraryfolders.vdf"))
		if err != nil {
			return ""
		}

		return steamPath
	}
	return ""
}

func FindGameDir() string {
	steamPath := findSteam()
	if steamPath == "" {
		return ""
	}

	libraryList := findSteamLibraries(steamPath)
	if libraryList == nil {
		return ""
	}

	gameDir := findRimWorldDirectory(libraryList)
	if gameDir == "" {
		return ""
	}

	return gameDir
}
