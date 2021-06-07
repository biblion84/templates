package main

import "os"

type TemplateData struct {
	Title       string
	Body        string
	Files       []os.DirEntry
	ButtonValue string
	Artist      Artist
	Countries   []Country
	Members     []int
}

type Country struct {
	Name   string
	Cities []string
}

type Artist struct {
	Name       string
	Age        int
	MusicGenre string
}
