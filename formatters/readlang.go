package formatters

import (
	"io"

	"github.com/ryanbrainard/jamak/pkg"
)

func FormatReadlang(items <-chan *pkg.Item, w io.Writer, options map[string]string) error {
	return formatXSV(items, w, options, ',')
}
