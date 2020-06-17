package main

func main() {
	options := GetOptions()

	if options.link {
		link(options.gameDir, options.modDir)
	}

	if options.pkg {
		packageMod(options.modDir)
	}
}
