package cmd

import (
	"fmt"
	"os"
	"rare/cmd/helpers"
	"rare/docs"
	"rare/pkg/markdowncli"
	"strings"

	"github.com/urfave/cli"
)

func docsFunction(c *cli.Context) error {
	docname := strings.ToLower(c.Args().First())

	if docname == "" || docname == "list" {
		fmt.Println("Available Docs:")
		entries, _ := docs.DocFS.ReadDir(".")
		for _, entry := range entries {
			fmt.Printf("  %s\n", strings.Title(strings.TrimSuffix(entry.Name(), ".md")))
		}

	} else if file, err := docs.DocFS.Open(docname + ".md"); err == nil {
		if err != nil {
			return cli.NewExitError(err, helpers.ExitCodeInvalidUsage)
		} else {
			markdowncli.WriteMarkdownToTerm(os.Stdout, file)
		}
	} else {
		return cli.NewExitError(fmt.Sprintf("No such doc '%s'", docname), helpers.ExitCodeInvalidUsage)
	}

	return nil
}

func docsCommand() *cli.Command {
	return &cli.Command{
		Name:      "docs",
		Usage:     "Access help documentation",
		ArgsUsage: "[doc]",
		Action:    docsFunction,
	}
}
