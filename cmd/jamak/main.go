package main

import (
	"flag"
	"os"
	"strings"

	"github.com/ryanbrainard/jamak/cmd"
	"github.com/ryanbrainard/jamak/pkg"
)

var fParser = flag.String("parser", "", "type of parser for input ["+strings.Join(cmd.AppCapabilities.Parsers, "|")+"]")
var fFormatter = flag.String("formatter", "", "type of formatter for output ["+strings.Join(cmd.AppCapabilities.Formatters, "|")+"]")
var fReadlangAccessToken = flag.String(pkg.OPT_READLANG_ACCESS_TOKEN, "", "")
var fReadlangBookId = flag.String(pkg.OPT_READLANG_BOOK_ID, "", "")
var fReadlangLanguage = flag.String(pkg.OPT_LANGUAGE, "", "")

func main() {
	flag.Parse()
	err := pkg.Run(
		os.Stdin,
		os.Stdout,
		cmd.ParseOptParser(*fParser),
		cmd.ParseOptFormatter(*fFormatter),
		map[string]string{
			pkg.OPT_READLANG_ACCESS_TOKEN: *fReadlangAccessToken,
			pkg.OPT_READLANG_BOOK_ID:      *fReadlangBookId,
			pkg.OPT_LANGUAGE:              *fReadlangLanguage,
		},
	)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}
