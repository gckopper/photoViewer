package main

import (
	"embed"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path"
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
	fileData, err := os.ReadFile(requestedFilename)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(fmt.Sprintf("Could not load file %s", requestedFilename)))
	}
	res.Write(fileData)
}

func main() {
	var file string
	flag.StringVar(&file, "file", "", "Opens a specific file")
	flag.Parse()

	file = path.Clean(file)

	// Create an instance of the app structure
	app := NewApp(file)

	// Create application with options
	err := wails.Run(&options.App{
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
		LogLevelProduction: logger.INFO,
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
