package parsers

import (
	"testing"

	"github.com/ryanbrainard/jamak/pkg"
	"github.com/stretchr/testify/require"
	"strings"
)

func TestParseYoutubeDownsub(t *testing.T) {
	videoUrl := strings.NewReader("https://youtu.be/4WSJrpo0EPQ")

	frames := make(chan *pkg.Frame, 100)
	opts := map[string]string{}
	err := ParseYoutubeDownsub(videoUrl, frames, opts)
	require.NoError(t, err)
	require.Equal(t, testFrames[0], <-frames)
	require.Equal(t, opts[pkg.OPT_READLANG_YOUTUBEID], "4WSJrpo0EPQ")
	require.Equal(t, opts[pkg.OPT_READLANG_TITLE], "[바나나 액츄얼리 시즌2 EP1] 나도 너랑 하고 싶어.")
}