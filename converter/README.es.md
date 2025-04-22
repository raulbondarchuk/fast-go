# Fast-Go Converter
[![Go Version](https://img.shields.io/badge/Go-1.23%2B-blue?logo=go&logoColor=white)](https://go.dev/doc/install) [![Status](https://img.shields.io/badge/Status-Active-brightgreen)](#)

[**Return to the main menu**](https://github.com/raulbondarchuk/fast-go/tree/main)

🌐 **Select Language / Seleccione el idioma / Виберіть мову / Выберите язык:**
- [English (Default)](README.md)
- [Español](README.es.md)
- [Українська](README.ua.md)
- [Русский](README.ru.md)

---

`converter` — paquete para convertir imágenes y vídeo con soporte para formatos MP4 y WebM, así como conversión de imágenes (PNG, JPEG, WebP, JFIF).

## Instalación

```bash
go get github.com/raulbondarchuk/fast-go/converter
```

> El paquete utiliza:
>
> - [disintegration/imaging](https://github.com/disintegration/imaging) para procesar imágenes.
> - [u2takey/ffmpeg-go](https://github.com/u2takey/ffmpeg-go) para convertir vídeo con FFmpeg.

## Estructuras y métodos

### ImageConfig

Configuración para convertir imágenes:

```go
type ImageConfig struct {
    FileName              string    // nombre del archivo de entrada
    File                  io.Reader // lector con el contenido de la imagen
    Width                 int       // ancho objetivo
    Height                int       // alto objetivo
    FormatToConvert       string    // formato deseado ("png", "jpg", "jpeg", "webp")
    StretchThreshold      float64   // umbral de estiramiento (en %)
    Quality               int       // calidad 1–5
    TransparentBackground bool      // fondo transparente en lugar de difuminado
    DirToStorage          string    // directorio de destino
}
```

**Métodos**:

- `Convert() (string, error)` — valida los parámetros, guarda un archivo temporal, procesa la imagen y devuelve la ruta del archivo final.
- `Delete(...string) error` — elimina el archivo especificado o el original por defecto.

**Ejemplo de uso**:

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
fmt.Println("Guardado en", newPath)
```

---

### LogoConfig

Configuración para procesar logotipos (siempre convierte a WebP):

```go
type LogoConfig struct {
    FileName     string    // nombre del archivo de logotipo
    File         io.Reader // lector con el contenido del logotipo
    DirToStorage string    // directorio de destino
    MaxWidth     int       // ancho máximo
    MaxHeight    int       // alto máximo
    MinWidth     int       // ancho mínimo
    MinHeight    int       // alto mínimo
}
```

**Métodos**:

- `Convert() (string, error)` — redimensiona, ajusta contraste y brillo, y guarda en formato WebP.

**Ejemplo**:

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
fmt.Println("Logotipo guardado en:", logoPath)
```

---

### VideoConfig

Configuración para convertir vídeo (MP4 ↔ WebM):

```go
type VideoConfig struct {
    FileName        string    // nombre del archivo de entrada (.mp4 o .webm)
    File            io.Reader // lector con el contenido del vídeo
    Width           int       // ancho objetivo del fotograma
    Height          int       // alto objetivo del fotograma
    FormatToConvert string    // formato de salida: "mp4" o "webm"
    Quality         int       // nivel de calidad 1–5 (CRF + preset)
    DirToStorage    string    // directorio de destino
}
```

**Métodos**:

- `Convert() (string, error)` — crea un archivo temporal, llama a ffmpeg-go para recodificar y devuelve la ruta final.
- `Delete(...string) error` — elimina el archivo final o el original.

**Ejemplo**:

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
fmt.Println("Vídeo convertido en:", outPath)
```

## Dependencias

- Go ≥ 1.21
- FFmpeg (debe estar instalado y disponible en PATH).

```bash
# Instalar imaging
go get github.com/disintegration/imaging
# Instalar ffmpeg-go
go get github.com/u2takey/ffmpeg-go
```


