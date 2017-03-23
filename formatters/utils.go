package formatters

import (
	"github.com/ryanbrainard/jjogaegi/pkg"
	"io"
)

func writeHeader(out io.Writer, options map[string]string) {
	if s, ok := options[pkg.OPT_HEADER]; ok && s != "" {
		out.Write([]byte(s + "\n"))
	}
}

func formatHangulHanja(item *pkg.Item, options map[string]string) string {
	switch options[pkg.OPT_HANJA] {
	case pkg.OPT_HANJA_PARENTHESIS:
		s := item.Hangul
		if len(item.Hanja) > 0 {
			s += " (" + item.Hanja + ")"
		}
		return s
	default:
		return item.Hangul
	}
}
