package cmd

import "github.com/urfave/cli"

var commands []cli.Command = []cli.Command{
	*filterCommand(),
	*histogramCommand(),
	*analyzeCommand(),
	*tabulateCommand(),
	*docsCommand(),
}

func GetSupportedCommands() []cli.Command {
	return commands
}
