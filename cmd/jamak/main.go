package main

import (
	"flag"
	"os"
	"strings"

	"github.com/ryanbrainard/jamak/cmd"
	"github.com/ryanbrainard/jamak/pkg"
)

var fParser = flag.String("parser", "list", "type of parser for input [" + strings.Join(cmd.Keys(cmd.AppCapabilities.Parsers), "|") + "]")
var fFormatter = flag.String("formatter", "tsv", "type of formatter for output [" + strings.Join(cmd.Keys(cmd.AppCapabilities.Formatters), "|") + "]")

func main() {
	flag.Parse()
	err := pkg.Run(
		os.Stdin,
		os.Stdout,
		cmd.ParseOptParser(*fParser),
		cmd.ParseOptFormatter(*fFormatter),
		map[string]string{},
	)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}
