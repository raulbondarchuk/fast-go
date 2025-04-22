# **Fast-Go Builder**
[![Go Version](https://img.shields.io/badge/Go-1.23%2B-blue?logo=go&logoColor=white)](https://go.dev/doc/install) [![Status](https://img.shields.io/badge/Status-Active-brightgreen)](#)

[**Вернуться на начальное меню**](https://github.com/raulbondarchuk/fast-go/tree/main)

**Fast-Go Builder** — это простая и эффективная библиотека для языка Go, разработанная для быстрого создания проектов. Она позволяет компилировать и упаковывать приложения с минимальными настройками, поддерживает несколько сред и упрощает управление конфигурациями для платформ Linux и Windows.

🌐 **Select Language / Seleccione el idioma / Виберіть мову / Выберите язык:**
- [English (Default)](https://github.com/raulbondarchuk/fast-go/tree/main/builder)
- [Español](README.es.md)
- [Українська](README.ua.md)
- Русский <---

---

## Быстрый доступ
- [Особенности](#особенности)
- [Пример использования](#пример-использования)
- [Установка](#установка)
- [Как использовать](#как-использовать)
- [Результат](#результат)

---

## **Установка**

Чтобы установить библиотеку, убедитесь, что у вас установлена версия Go **1.19** или выше.

1. Инициализируйте ваш Go-модуль, если это еще не сделано:
   ```bash
   go mod init <название_вашего_проекта>
   ```

2. Добавьте библиотеку Fast-Go Builder в ваш проект:
   ```bash
   go get github.com/raulbondarchuk/fast-go/builder
   ```

3. Импортируйте библиотеку в вашем коде:
   ```go
   import "github.com/raulbondarchuk/fast-go/builder"
   ```

---

## **Особенности**
- **Быстрая настройка**: Минимальные требования для беспроблемного процесса сборки.
- **Поддержка нескольких платформ**: Компиляция бинарных файлов для сред Linux и Windows.
- **Гибкая конфигурация**: Автоматическое нахождение и обновление файлов конфигурации (например, `toml`, `yaml`) во время сборки.
- **Настраиваемость**: Легкая настройка исходных файлов, выходных директорий и имен файлов.
- **Простая интеграция**: Интеграция `Fast-Go Builder` в существующие проекты для упрощения процесса сборки.

---

## **Пример использования**

```go
package main

import (
	"github.com/raulbondarchuk/fast-go/builder"
)

func Build() {
	builderConfig := builder.BuildConfig{
		DefaultMode:      "dev", // Например: local, dev, prod
		OutputFilename:   "мой-приложение-сборка",
		OutputDir:        "./",
		SourceFile:       "./cmd/main.go",
		BuildLinux:       true,
		BuildWindows:     false,
		PossibleDirs:     []string{"", "configs", "cfg", "config", "internal/config"},
		ConfigExtensions: []string{"toml", "yaml"},
	}
	builderConfig.Run()
}
```

Вы можете интегрировать эту функциональность сборки в ваш `main.go`:

```go
package main

import (
	build "api/pkg/build"
	"api/internal/app"
	"flag"
)

func main() {
	// Определение флага для команды сборки
	buildFlag := flag.Bool("build", false, "Запустить процесс сборки")
	flag.Parse()

	if *buildFlag {
		build.Build()
	} else {
		app.Run()
	}
}
```

---

## **Как использовать**

1. **Определите конфигурацию сборки:**
   Настройте структуру `BuildConfig` с учетом деталей вашего проекта:
   - **DefaultMode**: Укажите режим (`dev`, `prod` или `local`).
   - **OutputFilename**: Задайте имя выходного бинарного файла.
   - **OutputDir**: Укажите выходной каталог (по умолчанию включает папку `builds` с отметкой времени).
   - **SourceFile**: Укажите основной Go-файл для компиляции.
   - **BuildLinux/BuildWindows**: Включите сборку для платформ Linux или Windows.
   - **PossibleDirs/ConfigExtensions**: Укажите, где искать файлы конфигурации.

2. **Запустите процесс сборки:**
   Вызовите `Run()`, чтобы выполнить процесс сборки.

3. **Опциональная интеграция с CLI:**
   Добавьте функциональность сборки в CLI вашего проекта, используя флаги, как показано в примере.

---

## **Результат**
- Скомпилированные бинарные файлы сохраняются в каталоге `builds` внутри указанного вами `OutputDir`.
- Файлы конфигурации обновляются в соответствии с текущим режимом и копируются вместе с бинарными файлами.

---

С помощью **Fast-Go Builder** упрощайте сборку ваших Go-проектов и сосредотачивайтесь на написании качественного кода. 🚀
