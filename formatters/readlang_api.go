package formatters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ryanbrainard/jamak/pkg"
	"io"
	"io/ioutil"
	"net/http"
)

const readlangUrl = "https://readlang.com"

func FormatReadlangApi(frames <-chan *pkg.Frame, w io.Writer, options map[string]string) error {
	if options[pkg.OPT_READLANG_ACCESS_TOKEN] == "" {
		return fmt.Errorf(pkg.OPT_READLANG_ACCESS_TOKEN + " is required")
	}

	var err error

	bookId := options[pkg.OPT_READLANG_BOOK_ID]
	if bookId == "" {
		bookId, err = readlangCreateTempBook(options)
		if err != nil {
			return err
		}
	}

	// TODO: why isn't this working with a pipe?
	buf := &bytes.Buffer{}

	err = FormatReadlang(frames, buf, options)
	if err != nil {
		return err
	}

	_, err = readlangApiUpsert("PATCH", "/book/"+bookId, buf, options)
	if err != nil {
		return err
	}

	fmt.Fprint(w, readlangUrl+"/library/"+bookId)
	return nil
}

func readlangCreateTempBook(options map[string]string) (string, error) {
	reqBody, err := json.Marshal(CreateBookRequest{
		Title:  "jamak temp title",
		Body:   "jamak temp body",
		Public: false,
		Source: "jamak",
	})
	if err != nil {
		return "", err
	}

	return readlangApiUpsert("POST", "/book", bytes.NewReader(reqBody), options)
}

func readlangApiUpsert(method string, path string, body io.Reader, options map[string]string) (string, error) {
	req, err := http.NewRequest(method, readlangUrl+"/api"+path, body)
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "Bearer "+options[pkg.OPT_READLANG_ACCESS_TOKEN])
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	result := UpsertBookResult{}
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return "", err
	}

	if !result.Success {
		return "", fmt.Errorf("failed to %s readlang book", method)
	}

	return result.BookId, nil
}

type CreateBookRequest struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Public bool   `json:"public"`
	Source string `json:"source"`
}

type UpsertBookResult struct {
	BookId  string `json:"bookID"`
	Success bool   `json:"success"`
}
