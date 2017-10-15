package parsers

import (
	"github.com/ryanbrainard/jamak/pkg"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func ParseDownsub(r io.Reader, frames chan<- *pkg.Frame, options map[string]string) error {
	videoUrlBytes, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	videoUrl, err := url.Parse(string(videoUrlBytes))
	if err != nil {
		return err
	}

	baseUrl, _ := url.Parse("http://downsub.com")
	searchUrl, _ := baseUrl.Parse("/?url=" + url.QueryEscape(videoUrl.String()))

	resp, err := http.Get(searchUrl.String())
	if err != nil {
		log.Printf("download type=downsub.search url=%q err=%q", searchUrl.String(), err)
		return err
	}
	defer resp.Body.Close()

	subtitleUrl, err := baseUrl.Parse(extractSubtitleUrl(resp.Body))
	if err != nil {
		return err
	}

	resp, err = http.Get(subtitleUrl.String())
	if err != nil {
		log.Printf("download type=downsub.srt url=%q err=%q", subtitleUrl.String(), err)
		return err
	}
	defer resp.Body.Close()

	// TODO: extract title!

	return ParseSRT(resp.Body, frames, options)
}

func extractSubtitleUrl(r io.Reader) string {
	z := html.NewTokenizer(r)
	stack := &tokenStack{}
	for {
		switch z.Next() {
		case html.StartTagToken:
			if string(z.Raw()) == `<div class="panel-body col-md-7 col-sm-7" id="show">` {
				stack.Push(z.Token())
			}
			if stack.Depth() > 0 {
				name, _ := z.TagName()
				if string(name) == "a" {
					stack.Push(z.Token())
				}
			}
			if stack.Depth() > 1 {
				key, value, _ := z.TagAttr()
				if string(key) == "href" {
					stack.Push(z.Token())
					return string(value)
				}
			}
		case html.ErrorToken:
			return ""
		}
	}
}

type tokenStack struct {
	stack []html.Token
}

func (ts *tokenStack) Push(v html.Token) {
	ts.stack = append(ts.stack, v)
}

func (ts *tokenStack) Pop() html.Token {
	res := ts.stack[ts.Depth()-1]
	ts.stack = ts.stack[:ts.Depth()-1]
	return res
}

func (ts *tokenStack) Peek() html.Token {
	return ts.stack[ts.Depth()-1]
}

func (ts *tokenStack) Depth() int {
	return len(ts.stack)
}
