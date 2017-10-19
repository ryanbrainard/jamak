package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/ryanbrainard/jamak/cmd"
	"github.com/ryanbrainard/jamak/pkg"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	basePath := "cmd/jamak-web"
	router.LoadHTMLGlob(basePath + "/templates/*.tmpl.html") // TODO: relative path
	router.Static("/static", basePath+"/static")
	indexTemplate := "index.tmpl.html"
	render := func(c *gin.Context, m Model, code int) {
		c.Negotiate(code, gin.Negotiate{
			Offered:  []string{binding.MIMEHTML, binding.MIMEJSON},
			HTMLName: indexTemplate,
			HTMLData: m,
			JSONData: m,
		})
	}

	router.OPTIONS("", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
	})

	router.GET("/", func(c *gin.Context) {
		var m Model
		m.init()

		if err := c.Bind(&m); err != nil {
			log.Printf("at=get.bind error=%q", err)
			m.Error = fmt.Sprint(err)
			render(c, m, http.StatusBadRequest)
			return
		}

		render(c, m, http.StatusOK)
	})

	router.POST("/", func(c *gin.Context) {
		var m Model
		m.init()

		if err := c.Bind(&m); err != nil {
			log.Printf("at=post.bind error=%q", err)
			m.Error = err.Error()
			render(c, m, http.StatusBadRequest)
			return
		}

		output := &bytes.Buffer{}

		err := pkg.Run(
			strings.NewReader(m.Input),
			output,
			cmd.ParseOptParser(m.Parser),
			cmd.ParseOptFormatter(m.Formatter),
			m.Options,
		)

		if err != nil {
			log.Printf("at=post.run error=%q", err)
			m.Error = err.Error()
			render(c, m, http.StatusBadRequest)
			return
		}

		m.Output = output.String()

		render(c, m, http.StatusOK)
	})

	router.Run(":" + port)
}

// TODO: rename; break out into separate models, but KISS for now
type Model struct {
	Input        string `form:"input"     json:"input"`
	Output       string `form:"output"    json:"output"`
	Error        string `form:"error"     json:"error"`
	Parser       string `form:"parser"    json:"parser"`
	Formatter    string `form:"formatter" json:"formatter"`
	Options      map[string]string `form:"options"   json:"options"`
	Capabilities cmd.Capabilities
}

func (m *Model) init() {
	m.Capabilities = cmd.AppCapabilities
}
