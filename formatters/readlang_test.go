package formatters

import (
	"bytes"
	"testing"

	"github.com/ryanbrainard/jamak/pkg"
	"github.com/stretchr/testify/assert"
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

func TestFormatReadlang(t *testing.T) {
	frames, out := setupTestFormat()
	err := FormatReadlang(frames, out, map[string]string{})
	assert.Nil(t, err)
	assert.Equal(t, `{"audioMap":[{"t":2.71,"w":0},{"t":6.923,"w":3}],"body":"예전에 영화에서 봤는데\n\n이렇게 앉아있으니까\n\n"}`, out.String())
}

func setupTestFormat() (<-chan *pkg.Frame, *bytes.Buffer) {
	frames := make(chan *pkg.Frame, 2)
	frames <- testFrames[0]
	frames <- testFrames[1]
	close(frames)
	return frames, new(bytes.Buffer)
}
