package formatters

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/ryanbrainard/jamak/pkg"
	"os"
)

func TestFormatReadlangApi(t *testing.T) {
	if os.Getenv("READLANG_ACCESS_TOKEN") == "" {
		t.Skip("READLANG_ACCESS_TOKEN not set")
	}

	frames, out := setupTestFormat()
	err := FormatReadlangApi(frames, out, map[string]string{
		pkg.OPT_READLANG_ACCESS_TOKEN: os.Getenv("READLANG_ACCESS_TOKEN"),
	})
	assert.Nil(t, err)
}
