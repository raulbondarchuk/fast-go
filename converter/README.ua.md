# Fast-Go Converter
[![Go Version](https://img.shields.io/badge/Go-1.23%2B-blue?logo=go&logoColor=white)](https://go.dev/doc/install) [![Status](https://img.shields.io/badge/Status-Active-brightgreen)](#)

[**Return to the main menu**](https://github.com/raulbondarchuk/fast-go/tree/main)

**Fast-Go Converter** — пакет для конвертації зображень, відео та аудіо з підтримкою різних форматів.

🌐 **Select Language / Seleccione el idioma / Виберіть мову / Выберите язык:**
- [English (Default)](https://github.com/raulbondarchuk/fast-go/tree/main/converter)
- [Español](README.es.md)
- Українська <---
- [Русский](README.ru.md)

---

## Швидкий Доступ
- [Встановлення](#встановлення)
- [Фото](#фото)
- [Логотип](#логотип)
- [Відео](#відео)
- [Аудіо](#аудіо)
- [Визначення Типу Файлу](#визначення-типу-файлу)

---

## Встановлення

```bash
go get github.com/raulbondarchuk/fast-go/converter
```

> Пакет використовує:
>
> - [disintegration/imaging](https://github.com/disintegration/imaging) для обробки зображень.
> - [u2takey/ffmpeg-go](https://github.com/u2takey/ffmpeg-go) для конвертації відео через FFmpeg.

## Структури та методи

### Фото

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
    FileName:              "avatar",
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

### Логотип

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
    FileName:     "logo",
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

### Відео

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
    FileName:        "input",
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

### Аудио

Configuration for converting audio:

```go
type AudioConfig struct {
    FileName        string    // name of the input file
    File            io.Reader // reader with the audio content
    Bitrate         int       // target bitrate (64-320 kbps)
    FormatToConvert string    // conversion format ("mp3", "m4a", "opus", "wav")
    DirToStorage    string    // directory to save the output
}
```

**Methods**:

- `Convert() (string, error)` — validates settings, processes the audio file using ffmpeg, and returns the path to the final file.
- `Delete(...string) error` — deletes the specified file or the original by default.

**Example**:

```go
audioCfg := &converter.AudioConfig{
    FileName:        "track",
    File:            audioReader,
    Bitrate:         192,
    FormatToConvert: "opus",
    DirToStorage:    "./audio",
}
audioPath, err := audioCfg.Convert()
if err != nil {
    log.Fatal(err)
}
fmt.Println("Audio converted to:", audioPath)
```

### Визначення Типу Файлу

Функція `DetermineFileType` дозволяє визначити тип файлу на основі його розширення. Вона повертає `FileType`, який може бути одним з наступних:

- `Image` для зображень
- `Video` для відео
- `Audio` для аудіо
- `Json` для JSON файлів
- `Unknown` для непідтримуваних або невідомих типів файлів

**Приклад**:

```go
fileType := converter.DetermineFileType("example.mp3")
switch fileType {
case converter.Image:
    fmt.Println("Це файл зображення.")
case converter.Video:
    fmt.Println("Це відеофайл.")
case converter.Audio:
    fmt.Println("Це аудіофайл.")
case converter.Json:
    fmt.Println("Це JSON файл.")
default:
    fmt.Println("Невідомий тип файлу.")
}
```

## Залежності

- Go ≥ 1.21
- FFmpeg (must be installed and available in PATH).

```bash
# Install imaging
go get github.com/disintegration/imaging
# Install ffmpeg-go
go get github.com/u2takey/ffmpeg-go
```

