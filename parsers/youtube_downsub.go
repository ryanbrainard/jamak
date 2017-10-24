package parsers

import (
	"fmt"
	"github.com/ryanbrainard/jamak/pkg"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

var YOUTUBE_REGEX = regexp.MustCompile(`^((?:https?:)?\/\/)?((?:www|m)\.)?((?:youtube\.com|youtu.be))(\/(?:[\w\-]+\?v=|embed\/|v\/)?)([\w\-]+)(\S+)?$`)

// Dirty hack because YouTube API is dumb
// See https://stackoverflow.com/questions/46864428/how-do-some-sites-download-youtube-captions
func ParseYoutubeDownsub(r io.Reader, frames chan<- *pkg.Frame, options map[string]string) error {
	videoUrlBytes, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	videoUrl, err := url.Parse(strings.TrimSpace(string(videoUrlBytes)))
	if err != nil {
		return err
	}

	youtubeUrlMatches := YOUTUBE_REGEX.FindStringSubmatch(videoUrl.String())
	if youtubeUrlMatches == nil || len(youtubeUrlMatches) < 6 {
		return fmt.Errorf("Invalid YouTube URL: %s", videoUrl)
	}
	options[pkg.OPT_YOUTUBEID] = youtubeUrlMatches[5]

	baseUrl, _ := url.Parse("http://downsub.com")
	searchUrl, _ := baseUrl.Parse("/?url=" + url.QueryEscape(videoUrl.String()))

	resp, err := http.Get(searchUrl.String())
	if err != nil {
		log.Printf("download type=downsub.search url=%q err=%q", searchUrl.String(), err)
		return err
	}
	defer resp.Body.Close()

	title, srtRelUrl := extractTitleSrtUrl(resp.Body, options[pkg.OPT_LANGUAGE], options[pkg.OPT_TRACK_KIND])
	if options[pkg.OPT_TITLE] == "" {
		if title == "" {
			return fmt.Errorf("Could not extract title")
		}
		options[pkg.OPT_TITLE] = title
	}
	if srtRelUrl == "" {
		return fmt.Errorf("Could not extract SRT URL")
	}

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
func extractTitleSrtUrl(r io.Reader, languageCode string, trackKind string) (title string, srtUrl string) {
	z := html.NewTokenizer(r)
	titleStack := &tokenStack{}
	srtStack := &tokenStack{}
	potentialSrtUrl := ""
	for {
		switch z.Next() {
		case html.StartTagToken:
			if string(z.Raw()) == `<span class="media-heading" style="font-weight: bold;">` {
				titleStack.Push(z.Token())
			}

			if string(z.Raw()) == `<div class="panel-body col-md-7 col-sm-7" id="show">` {
				srtStack.Push(z.Token())
			}
			if srtStack.Depth() >= 1 {
				name, _ := z.TagName()
				if string(name) == "a" {
					srtStack.Push(z.Token())
				}
			}
			if srtStack.Depth() >= 2 {
				key, value, _ := z.TagAttr()
				if string(key) == "href" {
					srtStack.Push(z.Token())
					potentialSrtUrl = string(value)
				}
			}
		case html.TextToken:
			if titleStack.Depth() > 0 {
				title = string(z.Text())
				titleStack.Pop()
			}
			if srtStack.Depth() > 2 {
				lang := string(z.Text())
				matchesLanguage := strings.Contains(lang, pkg.SupportedLanguages[languageCode])
				matchesTrackKind := trackKind != "ASR" || strings.Contains(lang, "(auto-generated)")
				if potentialSrtUrl != "" && matchesLanguage && matchesTrackKind {
					srtUrl = potentialSrtUrl
					return
				}
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
