package main

func main() {
	options := GetOptions()

	if options.link {
		link(options.gameDir, options.modDir)
	}

	if options.fetchLib != "" {
		fetch(options.gameDir, options.fetchLib)
	}

	if options.pkg {
		packageMod(options.modDir)
	}
}
