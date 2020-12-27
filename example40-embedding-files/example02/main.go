package main

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//go:embed assets/* templates/*
	var f embed.FS

	router := gin.Default()

	// example: /public/assets/images/example.png
	router.StaticFS("/public", http.FS(f))

	templ := template.Must(template.New("").ParseFS(f, "templates/*.tmpl", "templates/foo/*.tmpl"))
	router.SetHTMLTemplate(templ)
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	router.GET("/foo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "bar.tmpl", gin.H{
			"title": "Foo website",
		})
	})

	router.GET("favicon.ico", func(c *gin.Context) {
		file, _ := f.ReadFile("assets/favicon.ico")
		c.Data(
			http.StatusOK,
			"image/x-icon",
			file,
		)
	})

	router.Run(":8080")
}
