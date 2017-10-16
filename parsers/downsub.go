package parsers

import (
	"github.com/ryanbrainard/jamak/pkg"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"fmt"
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

	title, srtRelUrl := extractTitleSrtUrl(resp.Body)
	if title == "" {
		return fmt.Errorf("extraction-failed part=title")
	}
	if srtRelUrl == "" {
		return fmt.Errorf("extraction-failed part=srt-url")
	}

	options[pkg.OPT_READLANG_TITLE] = title

	srtUrl, err := baseUrl.Parse(srtRelUrl)
	if err != nil {
		return err
	}

	resp, err = http.Get(srtUrl.String())
	if err != nil {
		log.Printf("download type=downsub.srt url=%q err=%q", srtUrl.String(), err)
		return err
	}
	defer resp.Body.Close()

	return ParseSRT(resp.Body, frames, options)
}

// dirty, dirty screen scrapping
func extractTitleSrtUrl(r io.Reader) (title string, srtUrl string) {
	z := html.NewTokenizer(r)
	titleStack := &tokenStack{}
	srtStack := &tokenStack{}
	for {
		switch z.Next() {
		case html.StartTagToken:
			if string(z.Raw()) == `<span class="media-heading" style="font-weight: bold;">` {
				titleStack.Push(z.Token())
			}

			if string(z.Raw()) == `<div class="panel-body col-md-7 col-sm-7" id="show">` {
				srtStack.Push(z.Token())
			}
			if srtStack.Depth() > 0 {
				name, _ := z.TagName()
				if string(name) == "a" {
					srtStack.Push(z.Token())
				}
			}
			if srtStack.Depth() > 1 {
				key, value, _ := z.TagAttr()
				if string(key) == "href" {
					srtStack.Push(z.Token())
					srtUrl = string(value)
					return
				}
			}
		case html.TextToken:
			if titleStack.Depth() > 0 {
				title = string(z.Text())
				titleStack.Pop()
			}
		case html.EndTagToken:
			for _, stack := range []*tokenStack{titleStack, srtStack} {
				if stack.Depth() > 0 && stack.Peek().Data == z.Token().Data {
					stack.Pop()
				}
			}
		case html.ErrorToken:
			return
		}
	}
	
	return
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
