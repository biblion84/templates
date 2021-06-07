package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(Artists)
	filenames := []string{"meme1.jpg", "memeDrole.jpg", "meme3.jpg"}

	strFile := ""
	for _, file := range filenames {
		fileSplitted := strings.Split(file, ".")
		strFile += fmt.Sprintf("<a href=\"/static/images/%s\">%s</a>\n", file, fileSplitted[0])
	}

	body := fmt.Sprintf(`
		<head>
			<link rel="stylesheet" href="/static/monCss.css">
		</head>
		<body>
			<div id="mesMemes">
				%s
			</div>
		
		</body>`, strFile)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(body))
}

func indexTemplate(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir("./static/")
	if failOnError(err, w) {
		return
	}

	err = templates["index.page.tmpl"].Execute(w, TemplateData{
		Title:       "MyTitle",
		Body:        "Hola",
		Files:       files,
		ButtonValue: "Mon super bouton",
		Artist: Artist{
			Name:       "Eminem",
			Age:        48,
			MusicGenre: "Rap",
		},
		Countries: []Country{
			{Name: "France", Cities: []string{"Paris", "Marseille", "Paris", "Paris", "Paris"}},
			{Name: "Allemagne", Cities: []string{"Berlin", "Marseille", "Paris", "Paris", "Paris"}},
			{Name: "UK", Cities: []string{"Londre", "Marseille", "Paris", "Paris", "Paris"}},
		},
		Members: []int{
			5,
			15,
			53,
			45,
			55,
		},
	})
	if err != nil {
		fmt.Println(err)
	}
}

func artist(w http.ResponseWriter, r *http.Request) {

	err := templates["artist.page.tmpl"].Execute(w, TemplateData{
		Body:        "Hola",
		ButtonValue: "Mon super bouton",
		Artist: Artist{
			Name:       "Eminem",
			Age:        48,
			MusicGenre: "Rap",
		},
		Countries: []Country{
			{Name: "France", Cities: []string{"Paris", "Marseille", "Paris", "Paris", "Paris"}},
			{Name: "Allemagne", Cities: []string{"Berlin", "Marseille", "Paris", "Paris", "Paris"}},
			{Name: "UK", Cities: []string{"Londre", "Marseille", "Paris", "Paris", "Paris"}},
		},
	})
	if err != nil {
		fmt.Println(err)
	}
}
