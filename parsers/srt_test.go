package parsers

import (
	"testing"

	"github.com/ryanbrainard/jamak/pkg"
	"github.com/stretchr/testify/assert"
	"os"
)

var testFrames = []*pkg.Frame{
	{
		Number:    1,
		StartTime: "00:00:02,710",
		EndTime:   "00:00:04,410",
		Text:      "예전에 영화에서 봤는데",
		StartWord: 0,
	}, {
		Number:    2,
		StartTime: "00:00:06,923",
		EndTime:   "00:00:08,453",
		Text:      "이렇게 앉아있으니까",
		StartWord: 3,
	},
}

func TestParseSRT(t *testing.T) {
	in, err := os.Open("fixtures/banana-2-1-simple.srt")
	assert.Nil(t, err)

	frames := make(chan *pkg.Frame, 100)
	ParseSRT(in, frames, map[string]string{})
	assert.Equal(t, testFrames[0], <-frames)
	assert.Equal(t, testFrames[1], <-frames)
}
