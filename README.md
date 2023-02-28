# README

## About

This is a simple program to view photos, it uses WebView (through the Wails library) to render the images meaning it can render anything a browser can!

## Requirements to develop and build

- [Wails](https://github.com/wailsapp/wails) v2.3.1
- [Go](https://go.dev/) 1.18

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on <http://localhost:34115>. Connect
to this in your browser, and you can call your Go code from devtools.

## Test files

The folder `test-images` conteins images to test format support and etc

## Building

To build a redistributable, production mode package, use `wails build`.

## Features

- Cross platform
- IS COOL!
- You can change/modify the frontend and make it look however you prefer!
- PLANNED FEATURES
  - Context menu option

## Formats

- Currently supports:
  - JPEG
  - PNG (with HDR support)
  - SVG
  - GIF
  - WEBP (Animeted or not)
  - BMP
  - ICO
  - APNG
- Planned support for:
  - JPEG XL (jxl)
  - RAW
  - AVIF
- May receive support:
  - KRA (Krita file format)
  - TIFF
  - HEIC
- Unsupported and no plan to support:
  - ORA (not used by anyone?)
  - PSD (Adobe...)

## Comparing PhotoViewer with default Windows Photos

| Image Format        | PhotoViewer (Win) | Win Photos | Ubuntu Image Viewer |
|---------------------|-------------------|------------|---------------------|
| AVIF (8 or 10 bits) | ❌               | ❌         | ❌                 |
| BMP                 | ✅               | ✅         | ✅                 |
| GIF                 | ✅               | ✅         | ✅                 |
| HEIC                | ❌               | ✅         | ❌                 |
| ICO (Icon)          | ✅               | ✅         | ✅                 |
| JPEG                | ✅               | ✅         | ✅                 |
| ORA (Open Raster)   | ❌               | ❌         | ❌                 |
| PNG                 | ✅               | ✅         | ✅                 |
| PNG HDR             | ✅               | ✅         | ✅                 |
| APNG (Animated PNG) | ✅               | ❌         | ❌                 |
| PSD (Photoshop)     | ❌               | ❌         | ❌                 |
| RAW                 | ❌               | ❌         | ❌                 |
| SVG                 | ✅               | ❌         | ✅                 |
| TIF (TIFF)          | ❌               | ❌         | ❌ (Sort of)       |
| WEBP                | ✅               | ❌         | ❌                 |
| WEBP Animated       | ✅               | ❌         | ❌                 |
| KRA                 | ❌               | ❌         | ❌                 |

## Tech Stack

- Backend
  - Golang
    - Wails
- Frontend
  - Svelte
    - Typescript
    - Vanilla CSS
    - Vanilla HTML
