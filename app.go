package main

import (
	"context"
	"encoding/base64"
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
	imgs    [][]string
	size    int
	current int
	first   string
}

// NewApp creates a new App application struct
func NewApp(file string) *App {
	return &App{
		current: 0,
		first:   file,
	}
}

/* type string_vec []string */

// Array of accepted extensions MUST REMAIN SORTED
var exts = []string{
	"bmp",
	"gif",
	"ico",
	"jpeg",
	"jpg",
	"png",
	"svg",
	"webp",
}

/* func (s string_vec) exists(str string) bool {
	for txt := range s {
		if str == s[txt] {
			return true
		}
	}
	return false
} */

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	dir, first := filepath.Split(a.first)
	if dir == "" {
		dir = "."
	}
	files, err := os.ReadDir(dir)
	if err != nil {
		return
	}
	size := 0
	a.imgs = make([][]string, len(files))
	for _, v := range files {
		if !v.IsDir() {
			name := v.Name()
			ext := strings.ToLower(filepath.Ext(name))
			_, exists := sort.Find(len(exts), func(i int) int {
				return strings.Compare(ext[1:], exts[i])
			})
			mime := mime.TypeByExtension(ext)
			if exists {
				content, err := os.ReadFile(filepath.Join(dir, name))
				if err != nil {
					println(err)
				}
				data := base64.StdEncoding.EncodeToString(content)
				array := make([]string, 3)
				array[0] = mime
				array[1] = data
				array[2] = name
				a.imgs[size] = array
				size++
				if name == first {
					a.current = size - 1
				}
			}
		}
	}
	a.imgs = a.imgs[0:size]
	a.size = len(a.imgs)
	if a.size == 0 {
		panic(fmt.Sprintf("Could not find images in this directory.\nDirectory: %s\nFile: %s", dir, first))
	}
}

func (a *App) First() []string {
	a.current = valid(a.current, a.size)
	return a.imgs[a.current]
}

func (a *App) Next() []string {
	a.current = valid(a.current+1, a.size)
	return a.imgs[a.current]
}

func (a *App) Priv() []string {
	a.current = valid(a.current-1, a.size)
	return a.imgs[a.current]
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
