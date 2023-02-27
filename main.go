package main

import (
	"embed"
	"net/http"
	"os"
	"path"

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

func main() {
	var file string

	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	file = path.Clean(file)

	// Create an instance of the app structure
	app := NewApp(file)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Photo Viewer",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
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
