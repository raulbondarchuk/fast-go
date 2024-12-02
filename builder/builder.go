package builder

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"golang.org/x/mod/modfile"
)

// BuildConfig is the configuration for the build process
type BuildConfig struct {
	DefaultMode    string   // dev, prod, local
	OutputFilename string   // name of the output file
	OutputDir      string   // path to the output directory
	SourceFile     string   // path to the source file
	BuildLinux     bool     // true if is necessary build for Linux
	BuildWindows   bool     // true if is necessary build for Windows
	PossibleDirs   []string // possible directories to find the config file
}

func finalization() {
	fmt.Printf("\n-- -- -- -- Build process completed. -- -- -- --\n\n")
}

// validate validates the build configuration
func (config *BuildConfig) validate() error {
	if config.DefaultMode == "" {
		return fmt.Errorf("DefaultMode is required")
	}
	if config.OutputFilename == "" {
		return fmt.Errorf("OutputFilename is required")
	}
	if config.OutputDir == "" {
		return fmt.Errorf("OutputDir is required")
	}
	if config.SourceFile == "" {
		return fmt.Errorf("SourceFile is required")
	}
	if len(config.PossibleDirs) == 0 {
		return fmt.Errorf("PossibleDirs is required")
	}
	return nil
}

// Run runs the build process
func (config *BuildConfig) Run() {
	if err := config.validate(); err != nil {
		log.Fatalf("Validation error: %v", err)
	}

	defer finalization()

	fmt.Printf("\n-- -- -- -- Build process initialized. -- -- -- --\n\n")
	// Get current working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	log.Println("Current working directory:", wd)

	// Read file go.mod
	data, err := os.ReadFile("go.mod")
	if err != nil {
		log.Fatalf("Error reading go.mod file: %v", err)
	}

	// Parsing go.mod
	modFile, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		log.Fatalf("Error parsing go.mod file: %v", err)
	}

	// Create output directory
	timestamp := time.Now().Format("2006-01-02-15-04-05")
	outputDir := filepath.Join(config.OutputDir, "builds", "build-"+timestamp)
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		log.Fatalf("Error creating build directory: %v", err)
	}

	// Build for Linux
	if config.BuildLinux {
		linuxOutput := config.OutputFilename + ".linux"
		if err := buildForOS("linux", linuxOutput, config.SourceFile, outputDir); err != nil {
			log.Fatal(err)
		}
	}

	// Build for Windows
	if config.BuildWindows {
		windowsOutput := modFile.Module.Mod.Path + ".exe"
		if err := buildForOS("windows", windowsOutput, config.SourceFile, outputDir); err != nil {
			log.Fatal(err)
		}
	}

	// Update and copy config file
	configFile, err := findConfigFile(wd, config.PossibleDirs)
	if err != nil {
		log.Fatalf("Error finding config file: %v", err)
	}
	destConfigFilePath := filepath.Join(outputDir, "config.toml")
	if err := updateAndCopyConfigFile(configFile, destConfigFilePath, config.DefaultMode); err != nil {
		log.Fatalf("Error updating and copying config file: %v", err)
	} else {
		log.Println("Successfully updated and copied config file to:", destConfigFilePath)
	}

}

// updateAndCopyConfigFile updates and copies the config file
func updateAndCopyConfigFile(src, dst, defaultMode string) error {
	// Read the config file
	input, err := os.ReadFile(src)
	if err != nil {
		return fmt.Errorf("error reading config file: %v", err)
	}

	// Convert the file content to a string
	content := string(input)

	// Check if [app] section exists, if not, add it
	appSectionRegex := regexp.MustCompile(`(?m)^\[app\]`)
	if !appSectionRegex.MatchString(content) {
		content = "[app]\n" + content
	}

	// Regular expression to find the "mode" parameter in the config file
	modeRegex := regexp.MustCompile(`(?m)^(\s*mode\s*=\s*)".*?"`)
	if modeRegex.MatchString(content) {
		// If the "mode" parameter is found, replace it with DEFAULT_MODE
		content = modeRegex.ReplaceAllString(content, fmt.Sprintf(`${1}"%s"`, defaultMode))
	} else {
		// If the "mode" parameter is not found, add it under [app] section
		newModeEntry := fmt.Sprintf("mode = \"%s\"\n", defaultMode)
		content = appSectionRegex.ReplaceAllString(content, fmt.Sprintf("[app]\n%s", newModeEntry))
	}

	// Get the current date and time
	buildDate := time.Now().Format("2006-01-02 15:04:05")

	// Add a comment line with the build creation date to the beginning of the file
	commentLine := fmt.Sprintf("# build creation date: %s\n", buildDate)
	content = commentLine + content

	// Save updated content to destination file
	if err := os.WriteFile(dst, []byte(content), 0644); err != nil {
		return fmt.Errorf("error writing updated config file: %v", err)
	}

	log.Printf("Config file updated with mode = \"%s\"\n", defaultMode)
	return nil
}

// buildForOS builds the project for the given OS
func buildForOS(goos, outputFile, sourceFile, outputDir string) error {
	cmd := exec.Command("go", "build", "-o", filepath.Join(outputDir, outputFile), sourceFile)
	cmd.Env = append(os.Environ(), fmt.Sprintf("GOOS=%s", goos), "GOARCH=amd64")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error building for %s: %v\nOutput: %s", goos, err, string(output))
	}
	log.Printf("Successfully built for %s: %s\n", goos, filepath.Join(outputDir, outputFile))
	return nil
}

// findConfigFile finds the config file in the possible directories
func findConfigFile(wd string, possibleDirs []string) (string, error) {
	for _, dir := range possibleDirs {
		files, err := os.ReadDir(filepath.Join(wd, dir))
		if err != nil {
			continue
		}
		for _, file := range files {
			if !file.IsDir() && strings.HasSuffix(file.Name(), ".toml") {
				return filepath.Join(wd, dir, file.Name()), nil
			}
		}
	}
	return "", fmt.Errorf("no .toml config file found")
}
