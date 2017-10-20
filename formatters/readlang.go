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

		book.Title = options[pkg.OPT_TITLE]
		book.Body += frame.Text + "\n\n"
		book.YouTubeID = options[pkg.OPT_YOUTUBEID]
		book.AudioMap = append(book.AudioMap, ReadlangAudiomapFrame{
			StartTime: startTime,
			StartWord: frame.StartWord,
		})
	}

	bookJson, err := json.Marshal(book)
	if err != nil {
		return err
	}

	_, err = w.Write(bookJson)
	if err != nil {
		return err
	}

	return nil
}

type ReadlangBook struct {
	Title            string                  `json:"title,omitempty"`
	Body             string                  `json:"plainText,omitempty"`
	HtmlMarkup       bool                    `json:"htmlMarkup"`       // always false
	GeneratedVersion int                     `json:"generatedVersion"` // always 0
	YouTubeID        string                  `json:"youTubeID,omitempty"`
	AudioMap         []ReadlangAudiomapFrame `json:"audioMap,omitempty"`
}

type ReadlangAudiomapFrame struct {
	StartTime float64 `json:"t"`
	StartWord int     `json:"w"`
}
