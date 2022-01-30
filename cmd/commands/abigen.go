package commands

import (
	"flag"

	"github.com/mitchellh/cli"
	"github.com/umbracle/go-web3/cmd/abigen"
)

// VersionCommand is the command to show the version of the agent
type AbigenCommand struct {
	UI cli.Ui
}

// Help implements the cli.Command interface
func (c *AbigenCommand) Help() string {
	return `Usage: greenhouse version

  Display the Greenhouse version`
}

// Synopsis implements the cli.Command interface
func (c *AbigenCommand) Synopsis() string {
	return "Display the Greenhouse version"
}

// Run implements the cli.Command interface
func (c *AbigenCommand) Run(args []string) int {
	var sources string
	var pckg string
	var output string
	var name string

	flag.StringVar(&sources, "source", "", "List of abi files")
	flag.StringVar(&pckg, "package", "main", "Name of the package")
	flag.StringVar(&output, "output", "", "Output directory")
	flag.StringVar(&name, "name", "", "name of the contract")

	flag.Parse()

	abigen.Parse(sources, pckg, output, name)

	return 0
}
