package parsers

import (
	"bufio"
	"io"

	"github.com/ryanbrainard/jamak/pkg"
	"strconv"
	"strings"
)

func ParseSRT(r io.Reader, frames chan<- *pkg.Frame, options map[string]string) error {
	scanner := bufio.NewScanner(r)

	frame := &pkg.Frame{}
	wordCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			frames <- frame
			frame = &pkg.Frame{}
			continue
		}

		if frame.Number == 0 {
			number, err := strconv.Atoi(line)
			if err != nil {
				return err
			}

			frame.Number = number
			continue
		}

		if frame.StartTime == "" {
			times := strings.Split(line, " --> ")
			frame.StartTime = times[0]
			frame.EndTime = times[1]
			continue
		}

		if frame.StartWord == 0 {
			frame.StartWord = wordCount
		}
		wordCount = wordCount + strings.Count(line, " ") + 1

		if frame.Text != "" {
			frame.Text += "\n"
		}

		frame.Text += line
	}

	if frame.Text != "" {
		frames <- frame
	}

	return nil
}
