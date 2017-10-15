package parsers

import (
	"strings"
	"testing"

	"github.com/ryanbrainard/jamak/pkg"
	"github.com/stretchr/testify/assert"
)

func TestParseListWithBullet(t *testing.T) {
	in := strings.NewReader(" • 안녕 라이언 Hello, Ryan")
	items := make(chan *pkg.Item, 100)
	ParseList(in, items, map[string]string{})
	assert.Equal(t, &pkg.Item{Id: "안녕 라이언", Hangul: "안녕 라이언", Def: pkg.Translation{English: "Hello, Ryan"}}, <-items)
}

func TestParseListWithNumberAndColon(t *testing.T) {
	in := strings.NewReader("1. 안녕 라이언: Hello, Ryan")
	items := make(chan *pkg.Item, 100)
	ParseList(in, items, map[string]string{})
	assert.Equal(t, &pkg.Item{Id: "안녕 라이언", Hangul: "안녕 라이언", Def: pkg.Translation{English: "Hello, Ryan"}}, <-items)
}
