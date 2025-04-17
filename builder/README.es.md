
# **Fast-Go**

🌐 Select Language / Seleccione el idioma / Виберіть мову / Выберите язык:
- [English (Default)](README.md)
- [Español](README.es.md)
- [Українська](README.ua.md)
- [Русский](README.ru.md)

---

# **Fast-Go Builder**
[![Versión de Go](https://img.shields.io/badge/Go-1.23%2B-blue?logo=go&logoColor=white)](https://go.dev/doc/install) [![Estado](https://img.shields.io/badge/Estado-Activo-brightgreen)](#)

**Fast-Go Builder** es una biblioteca simple y eficiente de Go diseñada para compilar proyectos de manera rápida. Permite compilar y empaquetar aplicaciones con una configuración mínima, admite múltiples entornos y facilita la gestión de configuraciones para plataformas Linux y Windows.

---

## **Tabla de Contenidos**
- [Características](#características)
- [Ejemplo de Uso](#ejemplo-de-uso)
- [Instalación](#instalación)
- [Cómo Usar](#cómo-usar)
- [Salida](#salida)

---

## **Instalación**

Para instalar la biblioteca, asegúrate de tener Go **1.19** o una versión superior.

1. Inicializa tu módulo de Go si aún no lo has hecho:
   ```bash
   go mod init <nombre_de_tu_proyecto>
   ```

2. Agrega la biblioteca Fast-Go Builder a tu proyecto:
   ```bash
   go get github.com/raulbondarchuk/fast-go/builder
   ```

3. Importa la biblioteca en tu código:
   ```go
   import "github.com/raulbondarchuk/fast-go/builder"
   ```

---

## **Características**
- **Configuración Rápida**: Configuración mínima necesaria para un proceso de compilación sin problemas.
- **Soporte Multi-Plataforma**: Compila binarios para entornos Linux y Windows.
- **Configuración Flexible**: Localiza y actualiza automáticamente archivos de configuración (por ejemplo, `toml`, `yaml`) durante el proceso de compilación.
- **Personalizable**: Permite definir fácilmente archivos fuente, directorios de salida y nombres de archivos.
- **Integración Sencilla**: Integra `Fast-Go Builder` en tus proyectos existentes para simplificar los procesos de compilación.

---

## **Ejemplo de Uso**

```go
package main

import (
	"github.com/raulbondarchuk/fast-go/builder"
)

func Build() {
	builderConfig := builder.BuildConfig{
		DefaultMode:      "dev", // Por ejemplo: local, dev, prod
		OutputFilename:   "mi-aplicacion-compilada",
		OutputDir:        "./",
		SourceFile:       "./cmd/main.go",
		BuildLinux:       true,
		BuildWindows:     false,
		PossibleDirs:     []string{"", "configs", "cfg", "config", "internal/config"},
		ConfigExtensions: []string{"toml", "yaml"},
		AddAppOnConfig:   false,
	}
	builderConfig.Run()
}
```

Puedes integrar esta funcionalidad de compilación en tu `main.go`:

```go
package main

import (
	build "api/pkg/build"
	"api/internal/app"
	"flag"
)

func main() {
	// Define una bandera para el comando de compilación
	buildFlag := flag.Bool("build", false, "Ejecutar el proceso de compilación")
	flag.Parse()

	if *buildFlag {
		build.Build()
	} else {
		app.Run()
	}
}
```

---

## **Cómo Usar**

1. **Define la Configuración de Compilación:**
   Configura la estructura `BuildConfig` con los detalles de tu proyecto:
   - **DefaultMode**: Define el modo (`dev`, `prod`, o `local`).
   - **OutputFilename**: Establece el nombre del binario de salida.
   - **OutputDir**: Especifica el directorio de salida (por defecto incluye una carpeta `builds` con marcas de tiempo).
   - **SourceFile**: Especifica el archivo principal de Go para compilar.
   - **BuildLinux/BuildWindows**: Habilita la compilación para las plataformas Linux o Windows.
   - **PossibleDirs/ConfigExtensions**: Define dónde buscar los archivos de configuración.

2. **Ejecuta el Proceso de Compilación:**
   Llama a `Run()` para ejecutar el proceso de compilación.

3. **Integración Opcional con CLI:**
   Agrega funcionalidad de compilación a tu CLI del proyecto utilizando banderas, como se muestra en el ejemplo.

---

## **Salida**
- Los binarios compilados se almacenan en el directorio `builds` dentro de tu `OutputDir` especificado.
- Los archivos de configuración se actualizan con el modo actual y se copian junto con los binarios.

---

Con **Fast-Go Builder**, simplifica las compilaciones de tus proyectos en Go y enfócate en escribir código de calidad. 🚀
