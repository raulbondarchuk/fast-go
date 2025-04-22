# Fast-Go Converter
[![Go Version](https://img.shields.io/badge/Go-1.23%2B-blue?logo=go&logoColor=white)](https://go.dev/doc/install) [![Status](https://img.shields.io/badge/Status-Active-brightgreen)](#)

[**Return to the main menu**](https://github.com/raulbondarchuk/fast-go/tree/main)

**Fast-Go Converter** — пакет для конвертации изображений, видео и аудио с поддержкой различных форматов.

🌐 **Select Language / Seleccione el idioma / Виберіть мову / Выберите язык:**
- [English (Default)](https://github.com/raulbondarchuk/fast-go/tree/main/converter)
- [Español](README.es.md)
- [Українська](README.ua.md)
- [Русский](README.ru.md)

---

## Быстрый доступ
- [Установка](#установка)
- [Изображения](#изображения)
- [Логотипы](#логотипы)
- [Видео](#видео)
- [Аудио](#аудио)
- [Определение типа файла](#определение-типа-файла)

---

## Установка

```bash
go get github.com/raulbondarchuk/fast-go/converter
```

> Пакет использует:
>
> - [disintegration/imaging](https://github.com/disintegration/imaging) для обработки изображений.
> - [u2takey/ffmpeg-go](https://github.com/u2takey/ffmpeg-go) для конвертации видео и аудио через FFmpeg.

## Изображения

Конфигурация для конвертации изображений:

```go
type ImageConfig struct {
    FileName              string    // имя входного файла без расширения
    File                  io.Reader // ридер с содержимым изображения
    Width                 int       // целевая ширина
    Height                int       // целевая высота
    FormatToConvert       string    // формат конвертации ("png", "jpg", "jpeg", "webp", "jfif")
    StretchThreshold      float64   // порог растяжения (%)
    Quality               int       // уровень качества 1–5
    TransparentBackground bool      // прозрачный фон вместо размытого
    DirToStorage          string    // директория для сохранения
}
```

**Методы**:
- `Convert() (string, error)` — проверяет настройки, сохраняет временный файл, обрабатывает изображение и возвращает путь к результату.
- `Delete(...string) error` — удаляет указанный файл или исходный.

**Пример**:

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
fmt.Println("Сохранено в", newPath)
```

## Логотипы

Конфигурация для обработки логотипов (всегда конвертирует в WebP):

```go
type LogoConfig struct {
    FileName     string    // имя входного файла без расширения
    File         io.Reader // ридер с содержимым логотипа
    DirToStorage string    // директория для сохранения
    MaxWidth     int       // максимальная ширина
    MaxHeight    int       // максимальная высота
    MinWidth     int       // минимальная ширина
    MinHeight    int       // минимальная высота
}
```

**Методы**:
- `Convert() (string, error)` — изменяет размер, настраивает контраст и яркость, сохраняет в WebP.

**Пример**:

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
fmt.Println("Логотип сохранён в:", logoPath)
```

## Видео

Конфигурация для конвертации видео (MP4 ↔ WebM):

```go
type VideoConfig struct {
    FileName        string    // имя входного файла без расширения (.mp4 или .webm)
    File            io.Reader // ридер с содержимым видео
    Width           int       // целевая ширина кадра
    Height          int       // целевая высота кадра
    FormatToConvert string    // формат вывода: "mp4" или "webm"
    Quality         int       // уровень качества 1–5 (CRF + preset)
    DirToStorage    string    // директория для сохранения
}
```

**Методы**:
- `Convert() (string, error)` — создаёт временный файл, выполняет перекодирование через ffmpeg-go и возвращает путь к результату.
- `Delete(...string) error` — удаляет итоговый или исходный файл.

**Пример**:

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
fmt.Println("Видео конвертировано в:", outPath)
```

## Аудио

Конфигурация для конвертации аудио:

```go
type AudioConfig struct {
    FileName        string    // имя входного файла без расширения
    File            io.Reader // ридер с содержимым аудио
    Bitrate         int       // битрейт (64–320 kbps)
    FormatToConvert string    // формат конвертации ("mp3", "m4a", "opus", "wav")
    DirToStorage    string    // директория для сохранения
}
```

**Методы**:
- `Convert() (string, error)` — проверяет настройки, конвертирует аудио через ffmpeg и возвращает путь к результату.
- `Delete(...string) error` — удаляет указанный или исходный файл.

**Пример**:

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
fmt.Println("Аудио конвертировано в:", audioPath)
```

## Определение типа файла

Функция `DetermineFileType` определяет тип файла по расширению и возвращает `FileType`:
- `Image` — изображение
- `Video` — видео
- `Audio` — аудио
- `Json`  — JSON
- `Unknown` — неизвестный формат

**Пример**:

```go
fileType := converter.DetermineFileType("example.mp3")
switch fileType {
case converter.Image:
    fmt.Println("Это изображение.")
case converter.Video:
    fmt.Println("Это видеофайл.")
case converter.Audio:
    fmt.Println("Это аудиофайл.")
case converter.Json:
    fmt.Println("Это JSON файл.")
default:
    fmt.Println("Неизвестный тип файла.")
}
```

## Зависимости

- Go ≥ 1.21
- FFmpeg (должен быть установлен и доступен в PATH).

```bash
# Установить imaging
go get github.com/disintegration/imaging
# Установить ffmpeg-go
go get github.com/u2takey/ffmpeg-go
```
