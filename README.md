# **Fast-Go Builder**

**Fast-Go Builder** is a simple and efficient Go library designed for quickly building Go projects. It allows you to compile and package your applications with minimal setup, supports multiple environments, and facilitates configuration management for Linux and Windows platforms.

---

## **Table of Contents**
- [Features](#features)
- [Example Usage](#example-usage)
- [Installation](#installation)
- [How to Use](#how-to-use)
- [Output](#output)

---

## **Features**
- **Quick Setup**: Minimal configuration required for a seamless build process.
- **Multi-Platform Support**: Build binaries for both Linux and Windows environments.
- **Flexible Configuration**: Automatically locates and updates configuration files (e.g., `toml`, `yaml`) during the build process.
- **Customizable**: Easily set source files, output directories, and filenames.
- **Easy Integration**: Embed `Fast-Go Builder` in your existing projects for streamlined builds.

---

## **Example Usage**

```go
package main

import (
	"github.com/raulbondarchuk/fast-go/builder"
)

func Build() {
	builderConfig := builder.BuildConfig{
		DefaultMode:      "dev", // For expample: local, dev, prod 
		OutputFilename:   "my-app-build",
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

You can integrate this build functionality in your `main.go`:

```go
package main

import (
	build "api/pkg/build"
	"api/internal/app"
	"flag"
)

func main() {
	// Define a flag for the build command
	buildFlag := flag.Bool("build", false, "Run the build process")
	flag.Parse()

	if *buildFlag {
		build.Build()
	} else {
		app.Run()
	}
}
```

---

## **Installation**

To install the library, ensure you have Go **1.19** or higher installed.

1. Initialize your Go module if not already done:
   ```bash
   go mod init <your_project_name>
   ```

2. Add the Fast-Go Builder library to your project:
   ```bash
   go get github.com/raulbondarchuk/fast-go/builder
   ```

3. Import the library in your code:
   ```go
   import "github.com/raulbondarchuk/fast-go/builder"
   ```

---

## **How to Use**

1. **Define Build Configuration:**
   Configure the `BuildConfig` struct with your project's details:
   - **DefaultMode**: Define the mode (`dev`, `prod`, or `local`).
   - **OutputFilename**: Set the output binary name.
   - **OutputDir**: Specify the output directory (default includes a timestamped "builds" folder).
   - **SourceFile**: Specify the main Go file for building.
   - **BuildLinux/BuildWindows**: Enable builds for Linux or Windows platforms.
   - **PossibleDirs/ConfigExtensions**: Define where to look for configuration files.

2. **Run the Build Process:**
   Call `Run()` to execute the build process.

3. **Optional Integration with CLI:**
   Add build functionality to your project CLI using flags, as shown in the example.

---

## **Output**
- Compiled binaries are stored in the `builds` directory within your specified `OutputDir`.
- Configuration files are updated with the current mode and copied alongside the binaries.

---

With **Fast-Go Builder**, simplify your Go project builds and focus on writing quality code. 🚀
