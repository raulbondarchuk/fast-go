# **Fast-Go Builder**
[![Go Version](https://img.shields.io/badge/Go-1.23%2B-blue?logo=go&logoColor=white)](https://go.dev/doc/install) [![Status](https://img.shields.io/badge/Status-Active-brightgreen)](#)

[**Повернутися до початкового меню**](https://github.com/raulbondarchuk/fast-go/tree/main)

**Fast-Go Builder** — це проста та ефективна бібліотека для мови Go, створена для швидкого компілювання проєктів. Вона дозволяє компілювати та пакувати застосунки з мінімальними налаштуваннями, підтримує декілька середовищ і спрощує керування конфігураціями для платформ Linux і Windows.

🌐 **Select Language / Seleccione el idioma / Виберіть мову / Выберите язык:**
- [English (Default)](https://github.com/raulbondarchuk/fast-go/tree/main/builder)
- [Español](README.es.md)
- Українська <---
- [Русский](README.ru.md)

---

## Швидкий Доступ
- [Особливості](#особливості)
- [Приклад використання](#приклад-використання)
- [Встановлення](#встановлення)
- [Як використовувати](#як-використовувати)
- [Результат](#результат)

---

## **Встановлення**

Щоб встановити бібліотеку, переконайтеся, що у вас встановлена версія Go **1.19** або новіша.

1. Ініціалізуйте свій Go-модуль, якщо це ще не зроблено:
   ```bash
   go mod init <назва_вашого_проєкту>
   ```

2. Додайте бібліотеку Fast-Go Builder до вашого проєкту:
   ```bash
   go get github.com/raulbondarchuk/fast-go/builder
   ```

3. Імпортуйте бібліотеку у вашому коді:
   ```go
   import "github.com/raulbondarchuk/fast-go/builder"
   ```

---

## **Особливості**
- **Швидке налаштування**: Мінімум налаштувань для безпроблемного процесу компіляції.
- **Підтримка декількох платформ**: Компіляція бінарних файлів для середовищ Linux та Windows.
- **Гнучка конфігурація**: Автоматичне знаходження та оновлення конфігураційних файлів (наприклад, `toml`, `yaml`) під час компіляції.
- **Налаштовуваність**: Легко встановлюйте файли джерел, каталоги виходу та імена файлів.
- **Легка інтеграція**: Включіть `Fast-Go Builder` у ваші існуючі проєкти для спрощення процесу компіляції.

---

## **Приклад використання**

```go
package main

import (
	"github.com/raulbondarchuk/fast-go/builder"
)

func Build() {
	builderConfig := builder.BuildConfig{
		DefaultMode:      "dev", // Наприклад: local, dev, prod
		OutputFilename:   "мій-застосунок-збірка",
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

Ви можете інтегрувати цю функціональність компіляції у ваш `main.go`:

```go
package main

import (
	build "api/pkg/build"
	"api/internal/app"
	"flag"
)

func main() {
	// Визначення прапорця для команди компіляції
	buildFlag := flag.Bool("build", false, "Запустити процес компіляції")
	flag.Parse()

	if *buildFlag {
		build.Build()
	} else {
		app.Run()
	}
}
```

---

## **Як використовувати**

1. **Визначте конфігурацію збірки:**
   Налаштуйте структуру `BuildConfig` відповідно до деталей вашого проєкту:
   - **DefaultMode**: Вкажіть режим (`dev`, `prod`, або `local`).
   - **OutputFilename**: Вкажіть назву вихідного бінарного файлу.
   - **OutputDir**: Вкажіть каталог виходу (за замовчуванням включає папку `builds` із позначкою часу).
   - **SourceFile**: Вкажіть основний Go-файл для компіляції.
   - **BuildLinux/BuildWindows**: Увімкніть компіляцію для платформ Linux або Windows.
   - **PossibleDirs/ConfigExtensions**: Визначте, де шукати конфігураційні файли.

2. **Запустіть процес збірки:**
   Викличте `Run()`, щоб виконати процес компіляції.

3. **Опціональна інтеграція з CLI:**
   Додайте функціональність компіляції до CLI вашого проєкту, використовуючи прапорці, як показано в прикладі.

---

## **Результат**
- Скомпільовані бінарні файли зберігаються в каталозі `builds` у вашому вказаному `OutputDir`.
- Конфігураційні файли оновлюються згідно з поточним режимом і копіюються разом із бінарними файлами.

---

**Fast-Go Builder** спрощуйте компіляції ваших Go-проєктів і зосередьтеся на написанні якісного коду. 🚀
