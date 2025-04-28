package converter

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"math"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
)

// ImageConfig holds configuration for image processing
type ImageConfig struct {
	FileName              string    // Name of the file
	File                  io.Reader // File reader for the image
	Width                 int       // Target width for the image
	Height                int       // Target height for the image
	FormatToConvert       string    // Desired format to convert the image to
	StretchThreshold      float64   // Threshold for stretching the image
	Quality               int       // Quality of the output image
	TransparentBackground bool      // Flag for transparent background
	DirToStorage          string    // Directory to store the processed image
}

// Checks if the desired format is supported
func (c *ImageConfig) isFormatSupported() bool {
	for _, format := range supportedFormatsImage {
		if format == c.FormatToConvert {
			return true
		}
	}
	return false
}

// Validates the configuration values
func (c *ImageConfig) validateValues() error {
	if c.Width <= 0 || c.Width > 8192 || c.Height <= 0 || c.Height > 8192 {
		return fmt.Errorf("width and height must be greater than 0 and less than 8192")
	}
	if c.StretchThreshold < 0 || c.StretchThreshold > 100 {
		return fmt.Errorf("stretch threshold must be between 0 and 100")
	}
	if c.Quality < 1 || c.Quality > 5 {
		return fmt.Errorf("quality must be between 1 and 5")
	}
	if !c.isFormatSupported() {
		return fmt.Errorf("unsupported format: %s", c.FormatToConvert)
	}
	if c.DirToStorage == "" {
		return fmt.Errorf("dir to storage is required")
	}
	if c.FileName == "" {
		return fmt.Errorf("file name is required")
	}
	if c.File == nil {
		return fmt.Errorf("file is required")
	}
	return nil
}

// ConvertImage converts the image to the desired format and saves it to the directory
func (c *ImageConfig) convertImage() (string, error) {
	if err := c.validateValues(); err != nil {
		return "", err
	}

	// Create the directory if it doesn't exist
	if err := os.MkdirAll(c.DirToStorage, 0755); err != nil {
		return "", fmt.Errorf("failed to create result dir: %w", err)
	}

	// Save original image temporarily in dirResult
	tempPath := filepath.Join(c.DirToStorage, c.FileName)
	outFile, err := os.Create(tempPath)
	if err != nil {
		return "", fmt.Errorf("failed to save original: %w", err)
	}
	if _, err := io.Copy(outFile, c.File); err != nil {
		outFile.Close()
		return "", fmt.Errorf("failed to write original: %w", err)
	}
	outFile.Close()

	// Process image
	processedPath, err := c.processImage(tempPath)
	if err != nil {
		_ = os.Remove(tempPath)
		return "", err
	}

	// Remove original
	_ = os.Remove(tempPath)

	// Final file path
	finalPath := filepath.Join(c.DirToStorage, filepath.Base(processedPath))
	if err := os.Rename(processedPath, finalPath); err != nil {
		return "", fmt.Errorf("failed to move processed image: %w", err)
	}

	return finalPath, nil
}

// Processes the image by resizing and converting it to the desired format
func (c *ImageConfig) processImage(sourcePath string) (string, error) {
	src, err := imaging.Open(sourcePath, imaging.AutoOrientation(true))
	if err != nil {
		return "", fmt.Errorf("error opening image: %w", err)
	}

	width := src.Bounds().Dx()
	height := src.Bounds().Dy()
	srcRatio := float64(width) / float64(height)
	targetRatio := float64(c.Width) / float64(c.Height)
	ratioDiff := math.Abs((srcRatio - targetRatio) / targetRatio * 100)

	var resized *image.NRGBA
	if ratioDiff <= c.StretchThreshold {
		// Resize image to target dimensions
		resized = imaging.Resize(src, c.Width, c.Height, imaging.Lanczos)
	} else if srcRatio > targetRatio {
		// Adjust height to maintain aspect ratio
		newHeight := int(float64(c.Width) / srcRatio)
		resized = imaging.Resize(src, c.Width, newHeight, imaging.Lanczos)
	} else {
		// Adjust width to maintain aspect ratio
		newWidth := int(float64(c.Height) * srcRatio)
		resized = imaging.Resize(src, newWidth, c.Height, imaging.Lanczos)
	}

	var background *image.NRGBA
	if c.TransparentBackground {
		// Create a transparent background
		background = imaging.New(c.Width, c.Height, image.Transparent)
	} else {
		// Create a blurred background
		background = c.createBlurredBackground(src, c.Width, c.Height)
	}

	// Center the resized image on the background
	posX := (c.Width - resized.Bounds().Dx()) / 2
	posY := (c.Height - resized.Bounds().Dy()) / 2
	final := imaging.Overlay(background, resized, image.Pt(posX, posY), 1.0)

	// Apply sharpening and contrast adjustments
	final = imaging.Sharpen(final, 0.5)
	final = imaging.AdjustContrast(final, 2)

	base := strings.TrimSuffix(filepath.Base(sourcePath), filepath.Ext(sourcePath))
	destPath := filepath.Join(c.DirToStorage, "processed_"+base+"."+c.FormatToConvert)

	// Determine quality based on configuration
	quality := map[int]int{1: 30, 2: 50, 3: 70, 4: 85, 5: 95}[c.Quality]
	if quality == 0 {
		quality = 80
	}

	// Save the final image in the desired format
	switch c.FormatToConvert {
	case PNG:
		err = imaging.Save(final, destPath, imaging.PNGCompressionLevel(pngCompressionLevel(c.Quality)))
	case JPEG, JPG:
		err = imaging.Save(final, destPath, imaging.JPEGQuality(quality))
	case WEBP:
		webpData, err := convertToWebp(final, uint8(quality))
		if err != nil {
			return "", fmt.Errorf("webp encode error: %w", err)
		}
		err = os.WriteFile(destPath, webpData, 0644)
		if err != nil {
			return "", fmt.Errorf("error saving webp file: %w", err)
		}
	default:
		return "", fmt.Errorf("unsupported format: %s", c.FormatToConvert)
	}
	if err != nil {
		return "", fmt.Errorf("error saving processed file: %w", err)
	}

	return destPath, nil
}

// Creates a blurred background for the image
func (c *ImageConfig) createBlurredBackground(src image.Image, targetWidth, targetHeight int) *image.NRGBA {
	srcRatio := float64(src.Bounds().Dx()) / float64(src.Bounds().Dy())
	targetRatio := float64(targetWidth) / float64(targetHeight)

	var bgWidth, bgHeight int
	if srcRatio > targetRatio {
		bgWidth = int(float64(targetHeight) * srcRatio)
		bgHeight = targetHeight
	} else {
		bgWidth = targetWidth
		bgHeight = int(float64(targetWidth) / srcRatio)
	}

	bg := imaging.Resize(src, bgWidth, bgHeight, imaging.Lanczos)
	bg = imaging.Blur(bg, 10.0)
	bg = imaging.AdjustBrightness(bg, 0)
	return imaging.CropCenter(bg, targetWidth, targetHeight)
}

// Determines the PNG compression level based on quality
func pngCompressionLevel(quality int) png.CompressionLevel {
	switch quality {
	case 1:
		return png.BestSpeed
	case 2, 3:
		return png.DefaultCompression
	case 4, 5:
		return png.BestCompression
	default:
		return png.DefaultCompression
	}
}

// deleteImage deletes the image from the directory
func (c *ImageConfig) deleteImage(reqFilePath ...string) error {
	var filePath string
	if len(reqFilePath) > 0 {
		filePath = reqFilePath[0]
	} else {
		filePath = c.DirToStorage + c.FileName
	}

	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("failed to delete image: %w", err)
	}
	return nil
}

// ---------------------------------------------------------------------
// ---------------------------- LOGOTYPE -------------------------------
// ---------------------------------------------------------------------

// LogoConfig holds configuration for logo processing
type LogoConfig struct {
	FileName        string    // Name of the logo file
	File            io.Reader // File reader for the logo
	FormatToConvert string    // Format to convert the logo to
	DirToStorage    string    // Directory to store the processed logo
	MaxWidth        int       // Maximum width for the logo
	MaxHeight       int       // Maximum height for the logo
	MinWidth        int       // Minimum width for the logo
	MinHeight       int       // Minimum height for the logo
}

// ProcessLogo handles logo upload, resizing with quality strategies, and saves it in the specified format.
func (cfg *LogoConfig) processLogo() (string, error) {
	if cfg.File == nil {
		return "", fmt.Errorf("logo file is required")
	}
	if cfg.DirToStorage == "" {
		return "", fmt.Errorf("DirToStorage is required")
	}

	// Set default format to webp if not specified
	if cfg.FormatToConvert == "" {
		cfg.FormatToConvert = "webp"
	}

	// Check if the format is supported
	supportedFormats := []string{"png", "webp", "jpg"}
	isSupported := false
	for _, format := range supportedFormats {
		if cfg.FormatToConvert == format {
			isSupported = true
			break
		}
	}
	if !isSupported {
		return "", fmt.Errorf("unsupported format: %s", cfg.FormatToConvert)
	}

	if err := os.MkdirAll(cfg.DirToStorage, 0755); err != nil {
		return "", err
	}

	tempPath := filepath.Join(cfg.DirToStorage, "original_"+cfg.FileName)
	outFile, err := os.Create(tempPath)
	if err != nil {
		return "", fmt.Errorf("failed to create logo temp file: %w", err)
	}
	if _, err := io.Copy(outFile, cfg.File); err != nil {
		outFile.Close()
		return "", fmt.Errorf("failed to write logo file: %w", err)
	}
	outFile.Close()

	src, err := imaging.Open(tempPath, imaging.AutoOrientation(true))
	if err != nil {
		os.Remove(tempPath)
		return "", fmt.Errorf("error opening uploaded logo: %w", err)
	}

	width := src.Bounds().Dx()
	height := src.Bounds().Dy()
	newW, newH := calculateDimensions(width, height, cfg.MaxWidth, cfg.MaxHeight, cfg.MinWidth, cfg.MinHeight)

	var resized *image.NRGBA
	if width > newW || height > newH {
		// Resize with intermediate step for better quality
		intermediate := imaging.Resize(src, (width+newW)/2, (height+newH)/2, imaging.Lanczos)
		resized = imaging.Resize(intermediate, newW, newH, imaging.Lanczos)
		resized = imaging.Sharpen(resized, 0.5)
	} else if width < newW || height < newH {
		// Resize using Mitchell-Netravali filter
		resized = imaging.Resize(src, newW, newH, imaging.MitchellNetravali)
	} else {
		// Clone the image if no resizing is needed
		resized = imaging.Clone(src)
	}

	// Adjust contrast and brightness
	resized = imaging.AdjustContrast(resized, 2)
	resized = imaging.AdjustBrightness(resized, 2)

	outputName := strings.TrimSuffix(cfg.FileName, filepath.Ext(cfg.FileName)) + "." + cfg.FormatToConvert
	outputPath := filepath.Join(cfg.DirToStorage, outputName)

	// Convert and save the logo in the specified format
	switch cfg.FormatToConvert {
	case "png":
		err = imaging.Save(resized, outputPath, imaging.PNGCompressionLevel(png.BestCompression))
	case "jpg":
		err = imaging.Save(resized, outputPath, imaging.JPEGQuality(95))
	case "webp":
		webpData, err := convertToWebp(resized, 95)
		if err != nil {
			os.Remove(tempPath)
			return "", err
		}
		err = os.WriteFile(outputPath, webpData, 0644)
		if err != nil {
			return "", fmt.Errorf("error saving webp file: %w", err)
		}
	default:
		return "", fmt.Errorf("unsupported format: %s", cfg.FormatToConvert)
	}
	if err != nil {
		os.Remove(tempPath)
		return "", fmt.Errorf("error saving processed file: %w", err)
	}

	os.Remove(tempPath)
	return outputPath, nil
}

// calculateDimensions resizes keeping proportions, with max/min thresholds
func calculateDimensions(width, height, maxW, maxH, minW, minH int) (int, int) {
	ratio := float64(width) / float64(height)

	if width > maxW || height > maxH {
		if width > maxW {
			newW := maxW
			newH := int(float64(newW) / ratio)
			if newH > maxH {
				newH = maxH
				newW = int(float64(newH) * ratio)
			}
			return newW, newH
		}
		newH := maxH
		newW := int(float64(newH) * ratio)
		return newW, newH
	} else if width < minW || height < minH {
		if width < minW {
			newW := minW
			newH := int(float64(newW) / ratio)
			if newH < minH {
				newH = minH
				newW = int(float64(newH) * ratio)
			}
			return newW, newH
		}
		newH := minH
		newW := int(float64(newH) * ratio)
		return newW, newH
	}

	return width, height
}
