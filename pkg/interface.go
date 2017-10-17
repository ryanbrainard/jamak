package pkg

import (
	"io"
)

type Frame struct {
	Number    int
	StartTime string
	EndTime   string
	Text      string
	StartWord int
}

type ParseFunc func(reader io.Reader, items chan<- *Frame, options map[string]string) error

type FormatFunc func(items <-chan *Frame, writer io.Writer, options map[string]string) error

const OPT_READLANG_TITLE = "readlang.title"
const OPT_READLANG_YOUTUBEID = "readlang.youTubeID"
const OPT_READLANG_ACCESS_TOKEN = "readlang.access_token"
const OPT_READLANG_BOOK_ID = "readlang.book_id"
