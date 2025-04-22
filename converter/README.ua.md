# Fast-Go Converter
[![Go Version](https://img.shields.io/badge/Go-1.23%2B-blue?logo=go&logoColor=white)](https://go.dev/doc/install) [![Status](https://img.shields.io/badge/Status-Active-brightgreen)](#)

[**Return to the main menu**](https://github.com/raulbondarchuk/fast-go/tree/main)

**Fast-Go Converter** — пакет для конвертації зображень, відео та аудіо з підтримкою різних форматів.

🌐 **Select Language / Seleccione el idioma / Виберіть мову / Выберите язык:**
- [English (Default)](https://github.com/raulbondarchuk/fast-go/tree/main/converter)
- [Español](README.es.md)
- [Українська](README.ua.md)
- [Русский](README.ru.md)

---

## Швидкий доступ
- [Installation](#Встановлення)
- [Image](#Фото)
- [Logo](#Логотип)
- [Video](#Відео)
- [Audio](#Аудіо)

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

Конфігурація для конвертації зображень:

```go
type ImageConfig struct {
    FileName              string    // ім'я вхідного файлу
    File                  io.Reader // рідер з вмістом зображення
    Width                 int       // бажана ширина
    Height                int       // бажана висота
    FormatToConvert       string    // формат конвертації ("png", "jpg", "jpeg", "webp", "jfif")
    StretchThreshold      float64   // поріг розтягування (у %)
    Quality               int       // якість 1–5
    TransparentBackground bool      // прозорий фон замість розмитого
    DirToStorage          string    // директорія для збереження
}
```

**Методи**:

- `Convert() (string, error)` — валідує налаштування, зберігає тимчасовий файл, оброблює зображення та повертає шлях до кінцевого файлу.
- `Delete(...string) error` — видаляє зазначений файл або початковий за замовчуванням.

**Приклад використання**:

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
fmt.Println("Збережено в", newPath)
```

---

### Логотип

Конфігурація для обробки логотипів (завжди конвертує в WebP):

```go
type LogoConfig struct {
    FileName     string    // ім'я файлу логотипу
    File         io.Reader // рідер з вмістом логотипу
    DirToStorage string    // директорія для збереження
    MaxWidth     int       // максимальна ширина
    MaxHeight    int       // максимальна висота
    MinWidth     int       // мінімальна ширина
    MinHeight    int       // мінімальна висота
}
```

**Методи**:

- `Convert() (string, error)` — змінює розміри, налаштовує контраст і яскравість, та зберігає у форматі WebP.

**Приклад**:

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
fmt.Println("Логотип збережено в:", logoPath)
```

---

### Відео

Конфігурація для конвертації відео (MP4 ↔ WebM):

```go
type VideoConfig struct {
    FileName        string    // ім'я вхідного файлу (.mp4 або .webm)
    File            io.Reader // рідер з вмістом відео
    Width           int       // бажана ширина кадру
    Height          int       // бажана висота кадру
    FormatToConvert string    // формат виводу: "mp4" або "webm"
    Quality         int       // рівень якості 1–5 (CRF + preset)
    DirToStorage    string    // директорія для збереження
}
```

**Методи**:

- `Convert() (string, error)` — створює тимчасовий файл, викликає ffmpeg-go для перекодування та повертає кінцевий шлях.
- `Delete(...string) error` — видаляє кінцевий або початковий файл.

**Приклад**:

```go
vidCfg := &converter.VideoConfig{
    FileName:        "video",
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
fmt.Println("Відео конвертовано в:", outPath)
```

---

### Аудіо

Конфігурація для конвертації аудіо:

```go
type AudioConfig struct {
    FileName        string    // ім'я вхідного файлу
    File            io.Reader // рідер з вмістом аудіо
    Bitrate         int       // цільовий бітрейт (64-320 kbps)
    FormatToConvert string    // формат конвертації ("mp3", "m4a")
    DirToStorage    string    // директорія для збереження
}
```

**Методи**:

- `Convert() (string, error)` — валідує налаштування, оброблює аудіофайл за допомогою ffmpeg та повертає шлях до кінцевого файлу.
- `Delete(...string) error` — видаляє зазначений файл або початковий за замовчуванням.

**Приклад**:

```go
audioCfg := &converter.AudioConfig{
    FileName:        "track",
    File:            audioReader,
    Bitrate:         192,
    FormatToConvert: "m4a",
    DirToStorage:    "./audio",
}
audioPath, err := audioCfg.Convert()
if err != nil {
    log.Fatal(err)
}
fmt.Println("Аудіо конвертовано в:", audioPath)
```

## Залежності

- Go ≥ 1.21
- FFmpeg (має бути встановлений та доступний у PATH).

```bash
# Встановити imaging
go get github.com/disintegration/imaging
# Встановити ffmpeg-go
go get github.com/u2takey/ffmpeg-go
```

