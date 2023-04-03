package main

import (
	"context"
	"fmt"
	"mime"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// App struct
type App struct {
	ctx     context.Context
	imgs    []Image
	size    int
	current int
	first   string
}

// App struct
type Image struct {
	Mime string `json:"mime"`
	Url  string `json:"content"`
	Name string `json:"name"`
}

// NewApp creates a new App application struct
func NewApp(file string) *App {
	return &App{
		current: 0,
		first:   file,
	}
}

// Array of accepted extensions MUST REMAIN SORTED
// Being sorted allows for binary search
func exts() []string {
	return []string{
		".bmp",
		".gif",
		".ico",
		".jpeg",
		".jpg",
		".png",
		".svg",
		".webp",
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	exts := exts()
	a.ctx = ctx

	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		return
	}
	size := 0
	a.imgs = make([]Image, len(files))
	for _, v := range files {
		if !v.IsDir() {
			name := v.Name()
			ext := strings.ToLower(filepath.Ext(name))
			_, exists := sort.Find(len(exts), func(i int) int {
				return strings.Compare(ext, exts[i])
			})
			if exists {
				mime := mime.TypeByExtension(ext)
				content := name
				image := Image{
					Mime: mime,
					Url:  content,
					Name: name,
				}
				a.imgs[size] = image
				size++
				if name == a.first {
					a.current = size - 1
				}
			}
		}
	}
	a.imgs = a.imgs[0:size]
	a.size = len(a.imgs)
	if a.size == 0 {
		a.imgs = []Image{
			{
				Mime: "image/png",
				Url:  "..",
				Name: "name",
			},
		}
		a.size = 1
	}
}

func valid(current int, size int) int {
	if current >= size {
		return 0
	}
	if current < 0 {
		return size - 1
	}
	return current
}

func (a *App) First() Image {
	fmt.Println(a.imgs)
	a.current = valid(a.current, a.size)
	return a.imgs[a.current]
}

func (a *App) Next() Image {
	a.current = valid(a.current+1, a.size)
	return a.imgs[a.current]
}

func (a *App) Priv() Image {
	a.current = valid(a.current-1, a.size)
	return a.imgs[a.current]
}
