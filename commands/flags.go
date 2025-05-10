package commands

type FlagDef struct {
	name      string
	shorthand string
	usage     string
}

var versionFlag = FlagDef{
	name:      "version",
	shorthand: "v",
	usage:     "Shows version info",
}

var nameFlag = FlagDef{
	name:      "name",
	shorthand: "n",
	usage:     "Greets the specified name",
}
