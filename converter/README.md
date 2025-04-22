# Fast-Go Converter
[![Go Version](https://img.shields.io/badge/Go-1.23%2B-blue?logo=go&logoColor=white)](https://go.dev/doc/install) [![Status](https://img.shields.io/badge/Status-Active-brightgreen)](#)

[**Return to the main menu**](https://github.com/raulbondarchuk/fast-go/tree/main)

🌐 **Select Language / Seleccione el idioma / Виберіть мову / Выберите язык:**
- [English (Default)](README.md)
- [Español](README.es.md)
- [Українська](README.ua.md)
- [Русский](README.ru.md)

---

`converter` — a package for converting images and video with support for MP4 and WebM formats, as well as image conversion (PNG, JPEG, WebP, JFIF).

## Installation

```bash
go get github.com/raulbondarchuk/fast-go/converter
```

> The package uses:
>
> - [disintegration/imaging](https://github.com/disintegration/imaging) for image processing.
> - [u2takey/ffmpeg-go](https://github.com/u2takey/ffmpeg-go) for video conversion using FFmpeg.

## Structures and Methods

### ImageConfig

Configuration for converting images:

```go
type ImageConfig struct {
    FileName              string    // name of the input file
    File                  io.Reader // reader with the image content
    Width                 int       // target width
    Height                int       // target height
    FormatToConvert       string    // conversion format ("png", "jpg", "jpeg", "webp", "jfif")
    StretchThreshold      float64   // stretch threshold (in %)
    Quality               int       // quality level, 1–5
    TransparentBackground bool      // transparent background instead of blurred
    DirToStorage          string    // directory to save the output
}
```

**Methods**:

- `Convert() (string, error)` — validates settings, saves a temporary file, processes the image, and returns the path to the final file.
- `Delete(...string) error` — deletes the specified file or the original by default.

**Usage Example**:

```go
cfg := &converter.ImageConfig{
    FileName:              "avatar.png",
    File:                  fileReader,
    Width:                 800,
    Height:                600,
    FormatToConvert:       "webp",
    Quality:               4,
    TransparentBackground: false,
    DirToStorage:          "./out",
}
newPath, err := cfg.Convert()
if err != nil {
    log.Fatal(err)
}
fmt.Println("Saved to", newPath)
```

---

### LogoConfig

Configuration for processing logos (always converts to WebP):

```go
type LogoConfig struct {
    FileName     string    // name of the logo file
    File         io.Reader // reader with the logo content
    DirToStorage string    // directory to save the output
    MaxWidth     int       // maximum width
    MaxHeight    int       // maximum height
    MinWidth     int       // minimum width
    MinHeight    int       // minimum height
}
```

**Methods**:

- `Convert() (string, error)` — resizes, adjusts contrast and brightness, and saves in WebP format.

**Example**:

```go
logoCfg := &converter.LogoConfig{
    FileName:     "logo.png",
    File:         logoReader,
    DirToStorage: "./logos",
    MaxWidth:     400,
    MaxHeight:    200,
    MinWidth:     100,
    MinHeight:    50,
}
logoPath, err := logoCfg.Convert()
if err != nil {
    log.Fatal(err)
}
fmt.Println("Logo saved to:", logoPath)
```

---

### VideoConfig

Configuration for converting video (MP4 ↔ WebM):

```go
type VideoConfig struct {
    FileName        string    // name of the input file (.mp4 or .webm)
    File            io.Reader // reader with the video content
    Width           int       // target frame width
    Height          int       // target frame height
    FormatToConvert string    // output format: "mp4" or "webm"
    Quality         int       // quality level, 1–5 (CRF + preset)
    DirToStorage    string    // directory to save the output
}
```

**Methods**:

- `Convert() (string, error)` — creates a temporary file, calls ffmpeg-go for re-encoding, and returns the final path.
- `Delete(...string) error` — deletes the final or original file.

**Example**:

```go
vidCfg := &converter.VideoConfig{
    FileName:        "input.mp4",
    File:            videoReader,
    Width:           1280,
    Height:          720,
    FormatToConvert: "webm",
    Quality:         3,
    DirToStorage:    "./videos",
}
outPath, err := vidCfg.Convert()
if err != nil {
    log.Fatal(err)
}
fmt.Println("Video converted to:", outPath)
```

## Dependencies

- Go ≥ 1.21
- FFmpeg (must be installed and available in PATH).

```bash
# Install imaging
go get github.com/disintegration/imaging
# Install ffmpeg-go
go get github.com/u2takey/ffmpeg-go
```

