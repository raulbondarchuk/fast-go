# Fast-Go Converter
[![Go Version](https://img.shields.io/badge/Go-1.23%2B-blue?logo=go&logoColor=white)](https://go.dev/doc/install) [![Status](https://img.shields.io/badge/Status-Active-brightgreen)](#)

[**Return to the main menu**](https://github.com/raulbondarchuk/fast-go/tree/main)

🌐 **Select Language / Seleccione el idioma / Виберіть мову / Выберите язык:**
- [English (Default)](README.md)
- [Español](README.es.md)
- [Українська](README.ua.md)
- [Русский](README.ru.md)

---

`converter` — пакет для конвертации изображений и видео с поддержкой форматов MP4 и WebM, а также конвертации изображений (PNG, JPEG, WebP, JFIF).

## Установка

```bash
go get github.com/raulbondarchuk/fast-go/converter
```

> Пакет использует:
>
> - [disintegration/imaging](https://github.com/disintegration/imaging) для обработки изображений.
> - [u2takey/ffmpeg-go](https://github.com/u2takey/ffmpeg-go) для конвертации видео через FFmpeg.

## Структуры и методы

### ImageConfig

Конфигурация для конвертации изображений:

```go
type ImageConfig struct {
    FileName              string    // имя входного файла
    File                  io.Reader // ридер с содержимым изображения
    Width                 int       // целевая ширина
    Height                int       // целевая высота
    FormatToConvert       string    // желаемый формат ("png", "jpg", "jpeg", "webp")
    StretchThreshold      float64   // порог растяжения (в %)
    Quality               int       // качество 1–5
    TransparentBackground bool      // прозрачный фон вместо размытого
    DirToStorage          string    // директория для сохранения
}
```

**Методы**:

- `Convert() (string, error)` — выполняет валидацию, сохраняет временный файл, обрабатывает изображение и возвращает путь к итоговому файлу.
- `Delete(...string) error` — удаляет указанный файл или исходный по умолчанию.

**Пример использования**:

```go
cfg := &converter.ImageConfig{
    FileName:        "avatar.png",
    File:            fileReader,
    Width:           800,
    Height:          600,
    FormatToConvert: "webp",
    Quality:         4,
    TransparentBackground: false,
    DirToStorage:    "./out",
}
newPath, err := cfg.Convert()
if err != nil {
    log.Fatal(err)
}
fmt.Println("Saved to", newPath)
```

---

### LogoConfig

Конфигурация для обработки логотипов (всегда конвертирует в WebP):

```go
type LogoConfig struct {
    FileName     string    // имя входного файла
    File         io.Reader // ридер с содержимым логотипа
    DirToStorage string    // директория для сохранения
    MaxWidth     int       // макс. ширина
    MaxHeight    int       // макс. высота
    MinWidth     int       // мин. ширина
    MinHeight    int       // мин. высота
}
```

**Методы**:

- `Convert() (string, error)` — ресайз, контраст/яркость, сохраняет в WebP.

**Пример**:

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
fmt.Println("Logo saved:", logoPath)
```

---

### VideoConfig

Конфигурация для конвертации видео (MP4 ↔ WebM):

```go
type VideoConfig struct {
    FileName        string    // имя входного файла (.mp4 или .webm)
    File            io.Reader // ридер с содержимым видео
    Width           int       // целевая ширина кадра
    Height          int       // целевая высота кадра
    FormatToConvert string    // целевой формат: "mp4" или "webm"
    Quality         int       // уровень качества 1–5 (CRF + preset)
    DirToStorage    string    // директория для сохранения
}
```

**Методы**:

- `Convert() (string, error)` — создаёт временный файл, вызывает ffmpeg-go для перекодирования и возвращает итоговый путь.
- `Delete(...string) error` — удаляет итоговый или исходный файл.

**Пример**:

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
fmt.Println("Video converted:", outPath)
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


