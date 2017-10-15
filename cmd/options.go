package cmd

import (
	"github.com/ryanbrainard/jamak/formatters"
	"github.com/ryanbrainard/jamak/parsers"
	"github.com/ryanbrainard/jamak/pkg"
)

type Capabilities struct {
	Parsers    map[string]string
	Formatters map[string]string
}

var AppCapabilities = Capabilities {
	Parsers: map[string]string{
		"srt":         "SRT",
	},
	Formatters: map[string]string{
		"srt":         "SRT",
		"readlang":         "Readlang audiomap",
	},
}

func Keys(m map[string]string) []string {
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func ParseOptParser(s string) pkg.ParseFunc {
	switch s {
	case "srt":
		return parsers.ParseSRT
	default:
		return nil
	}
}

func ParseOptFormatter(s string) pkg.FormatFunc {
	switch s {
	//case "srt":
	//	return formatters.FormatSRT
	case "readlang":
		return formatters.FormatReadlang
	default:
		return nil
	}
}
