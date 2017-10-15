package parsers

import (
	"bufio"
	"io"
	"strings"

	"github.com/ryanbrainard/jamak/pkg"
)

func ParseNaverTable(r io.Reader, items chan<- *pkg.Item, options map[string]string) error {
	i := 0
	rawTerms := []string{}
	scanner := bufio.NewScanner(r)
	scanner.Split(SplitDefs)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		if len(rawTerms) == 0 {
			rawTerms = strings.Split(line, " ")
			continue
		}

		hangulTerm, hanjaTerm := splitHangul(rawTerms[i])

		items <- &pkg.Item{
			Hangul: sanitize(string(hangulTerm)),
			Hanja:  sanitize(string(hanjaTerm)),
			Def: pkg.Translation{
				English: sanitize(string(line)),
			},
		}

		i++
	}

	return nil
}

func SplitDefs(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF {
		if len(data) == 0 {
			return 0, nil, nil
		}
		return len(data), data, nil
	}

	oneByte := []byte("1")[0]
	for i, d := range data {
		if d == oneByte {
			return i + 1, data[0:i], nil
		}
	}

	return 0, nil, nil
}
