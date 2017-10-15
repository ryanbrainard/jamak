package parsers

import (
	"testing"

	"github.com/ryanbrainard/jamak/pkg"
	"github.com/stretchr/testify/assert"
	"strings"
)

func TestParseDownsub(t *testing.T) {
	videoUrl := strings.NewReader("https://youtu.be/4WSJrpo0EPQ")

	frames := make(chan *pkg.Frame, 100)
	ParseDownsub(videoUrl, frames, map[string]string{})
	assert.Equal(t, testFrames[0], <-frames)
}
