package cmd

import (
	"github.com/ryanbrainard/jamak/formatters"
	"github.com/ryanbrainard/jamak/parsers"
	"github.com/ryanbrainard/jamak/pkg"
)

var registeredParsers = map[string]pkg.ParseFunc{
	"youtube": parsers.ParseYoutubeDownsub,
	"srt":     parsers.ParseSRT,
}

var registeredFormatters = map[string]pkg.FormatFunc{
	"readlang":     formatters.FormatReadlang,
	"readlang-api": formatters.FormatReadlangApi,
}

type Capabilities struct {
	Parsers    []string
	Formatters []string
}

var AppCapabilities = Capabilities{
	Parsers:    mapKeysParsers(registeredParsers),
	Formatters: mapKeysFormatters(registeredFormatters),
}

func ParseOptParser(s string) pkg.ParseFunc {
	return registeredParsers[s]
}

func ParseOptFormatter(s string) pkg.FormatFunc {
	return registeredFormatters[s]
}

func mapKeysParsers(m map[string]pkg.ParseFunc) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func mapKeysFormatters(m map[string]pkg.FormatFunc) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}
