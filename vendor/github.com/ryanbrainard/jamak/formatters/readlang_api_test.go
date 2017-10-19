package formatters

import (
	"github.com/ryanbrainard/jamak/pkg"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFormatReadlangApi(t *testing.T) {
	if os.Getenv("READLANG_ACCESS_TOKEN") == "" {
		t.Skip("READLANG_ACCESS_TOKEN not set")
	}

	frames, out := setupTestFormat()
	err := FormatReadlangApi(frames, out, map[string]string{
		pkg.OPT_READLANG_ACCESS_TOKEN: os.Getenv("READLANG_ACCESS_TOKEN"),
		pkg.OPT_READLANG_TITLE:        "Test Title",
	})
	assert.Nil(t, err)
}
