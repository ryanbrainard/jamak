package parsers

import (
	"testing"

	"github.com/ryanbrainard/jamak/pkg"
	"github.com/stretchr/testify/require"
	"strings"
)

func TestParseYoutubeDownsub(t *testing.T) {
	videoUrl := strings.NewReader("https://youtu.be/lwEkc0iJ5Es")

	frames := make(chan *pkg.Frame, 1000)
	opts := map[string]string{pkg.OPT_LANGUAGE: "id"}
	err := ParseYoutubeDownsub(videoUrl, frames, opts)
	require.NoError(t, err)
	require.Equal(t, "Alih bahasa dipersembahkan oleh  The Classmates Team @ALexTV", (<-frames).Text)
	require.Equal(t, opts[pkg.OPT_YOUTUBEID], "lwEkc0iJ5Es")
	require.Equal(t, opts[pkg.OPT_TITLE], "E1 Sekolah 2013 || Korean Drama's School 2013 English Subtitle ||")
}
