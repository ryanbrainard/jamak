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
		// TODO: move to parser
		startTimeSplit := strings.Split(frame.StartTime, ":")

		startTimeHour, err := strconv.ParseFloat(startTimeSplit[0], 64)
		if err != nil {
			return err
		}

		startTimeMin, err := strconv.ParseFloat(startTimeSplit[1], 64)
		if err != nil {
			return err
		}

		startTimeSec, err := strconv.ParseFloat(strings.Replace(startTimeSplit[2], ",", ".", 1), 64)
		if err != nil {
			return err
		}

		startTime := (startTimeHour * 60 * 60) + (startTimeMin * 60) + startTimeSec

		book.Body += frame.Text + "\n\n"
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
	Body string `json:"body"`
}

type ReadlangAudiomapFrame struct {
	StartTime float64 `json:"t"`
	StartWord int     `json:"w"`
}
