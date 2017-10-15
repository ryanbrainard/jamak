package pkg

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	in := &bytes.Buffer{}
	out := &bytes.Buffer{}
	parser := func(r io.Reader, items chan<- *Frame, options map[string]string) error {
		items <- &Frame{Text: "시험"}
		return nil
	}
	formatter := func(items <-chan *Frame, w io.Writer, options map[string]string) error {
		for item := range items {
			w.Write([]byte(item.Text))
		}
		return nil
	}
	options := map[string]string{}

	err := Run(in, out, parser, formatter, options)
	assert.Nil(t, err)
	assert.Equal(t, "시험", out.String())
}

func TestRun_ParserError(t *testing.T) {
	in := &bytes.Buffer{}
	out := &bytes.Buffer{}
	parser := func(r io.Reader, items chan<- *Frame, options map[string]string) error {
		return fmt.Errorf("boom: parser")
	}
	formatter := func(items <-chan *Frame, w io.Writer, options map[string]string) error {
		for item := range items {
			w.Write([]byte(item.Text))
		}
		return nil
	}
	options := map[string]string{}

	err := Run(in, out, parser, formatter, options)
	assert.NotNil(t, err)
	assert.Equal(t, "", out.String())
}

func TestRun_FormatterError(t *testing.T) {
	in := &bytes.Buffer{}
	out := &bytes.Buffer{}
	parser := func(r io.Reader, items chan<- *Frame, options map[string]string) error {
		items <- &Frame{Text: "시험"}
		close(items)
		return nil
	}
	formatter := func(items <-chan *Frame, w io.Writer, options map[string]string) error {
		return fmt.Errorf("boom: formatter")
	}
	options := map[string]string{}

	err := Run(in, out, parser, formatter, options)
	assert.NotNil(t, err)
	assert.Equal(t, "", out.String())
}

func TestRun_NilParser(t *testing.T) {
	in := &bytes.Buffer{}
	out := &bytes.Buffer{}
	formatter := func(items <-chan *Frame, w io.Writer, options map[string]string) error {
		return nil
	}
	options := map[string]string{}

	err := Run(in, out, nil, formatter, options)
	assert.NotNil(t, err)
	assert.Equal(t, "", out.String())
}

func TestRun_NilFormatter(t *testing.T) {
	in := &bytes.Buffer{}
	out := &bytes.Buffer{}
	parser := func(r io.Reader, items chan<- *Frame, options map[string]string) error {
		items <- &Frame{Text: "시험"}
		close(items)
		return nil
	}
	options := map[string]string{}

	err := Run(in, out, parser, nil, options)
	assert.NotNil(t, err)
	assert.Equal(t, "", out.String())
}
