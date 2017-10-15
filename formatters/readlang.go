package formatters

import (
	"io"

	"encoding/json"
	"github.com/ryanbrainard/jamak/pkg"
	"strconv"
	"strings"
)

func FormatReadlang(frames <-chan *pkg.Frame, w io.Writer, options map[string]string) error {
	book := &ReadlangBook{}
	for frame := range frames {
		startTimeStr := strings.Replace(frame.StartTime, ",", ".", 1)
		startTimeStr = strings.Replace(startTimeStr, "00:00:", "", 1) // TODO: support
		startTime, err := strconv.ParseFloat(startTimeStr, 64)
		if err != nil {
			return err
		}

		book.AudioMap = append(book.AudioMap, ReadlangAudiomapFrame{
			StartTime: startTime,
			StartWord: frame.StartWord,
		})
	}

	// TODO: streaming?
	bookJson, err := json.Marshal(book)
	if err != nil {
		return err
	}

	w.Write(bookJson)

	return nil
}

type ReadlangBook struct {
	AudioMap []ReadlangAudiomapFrame `json:"audioMap"`
}

type ReadlangAudiomapFrame struct {
	StartTime float64 `json:"t"`
	StartWord int     `json:"w"`
}
