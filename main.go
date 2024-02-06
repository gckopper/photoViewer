package main

import (
	"embed"
	"fmt"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed frontend/dist
var assets embed.FS

type FileLoader struct {
	http.Handler
}

func NewFileLoader() *FileLoader {
	return &FileLoader{}
}

func (h *FileLoader) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var err error
	requestedFilename := strings.TrimPrefix(req.URL.Path, "/")
	ext := filepath.Ext(requestedFilename)
	mime := mime.TypeByExtension(ext)
	//println("Requesting file:", requestedFilename)
	fileData, err := os.ReadFile(requestedFilename)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(fmt.Sprintf("Could not load file %s", requestedFilename)))
	}
	res.Header().Set("Content-Type", mime)
	res.Write(fileData)
}

func changeDir(dir string) {
	err := os.Chdir(dir)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	var file string
    var err error

	if len(os.Args) >= 2 {
		file = os.Args[1]
	}

	file = path.Clean(file)
	dir, first := filepath.Split(file)
    if dir != "" {
        changeDir(dir)
    }
    fileInfo, err := os.Stat(first)
	if err != nil {
		fmt.Println(err)
		return
	}
    if fileInfo.IsDir() {
        changeDir(first)
        first = ""
    }

	// Create an instance of the app structure
	app := NewApp(first)

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "Photo Viewer",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets:  assets,
			Handler: NewFileLoader(),
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		WindowStartState: options.Maximised,
		Bind: []interface{}{
			app,
		},
		EnableFraudulentWebsiteDetection: false,
		LogLevelProduction:               logger.WARNING,
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
