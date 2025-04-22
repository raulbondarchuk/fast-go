package converter

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

type VideoConfig struct {
	FileName              string
	File                  io.Reader
	Width                 int
	Height                int
	FormatToConvert       string
	Quality               int
	TransparentBackground bool
	DirToStorage          string
}

func (c *VideoConfig) isFormatSupported() bool {
	for _, format := range supportedFormatsVideo {
		if format == c.FormatToConvert {
			return true
		}
	}
	return false
}

func (c *VideoConfig) validateValues() error {
	if c.Width <= 0 || c.Height <= 0 {
		return fmt.Errorf("width and height must be greater than 0")
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
	if !c.isFileExtensionSupported() {
		return fmt.Errorf("unsupported file extension: %s", filepath.Ext(c.FileName))
	}
	return nil
}

func (c *VideoConfig) isFileExtensionSupported() bool {
	ext := strings.ToLower(filepath.Ext(c.FileName))
	return ext == ".mp4" || ext == ".webm"
}

func (c *VideoConfig) processVideo() (string, error) {
	if err := c.validateValues(); err != nil {
		return "", err
	}

	if err := os.MkdirAll(c.DirToStorage, 0755); err != nil {
		return "", fmt.Errorf("failed to create result dir: %w", err)
	}

	// Save original video temporarily in dirResult
	tempPath := filepath.Join(c.DirToStorage, c.FileName)
	outFile, err := os.Create(tempPath)
	if err != nil {
		return "", fmt.Errorf("failed to save original: %w", err)
	}
	defer os.Remove(tempPath) // Ensure the file is removed after processing
	if _, err := io.Copy(outFile, c.File); err != nil {
		outFile.Close()
		return "", fmt.Errorf("failed to write original: %w", err)
	}
	outFile.Close()

	// Create a path for the processed file
	filename := strings.TrimSuffix(filepath.Base(tempPath), filepath.Ext(tempPath))
	destPath := filepath.Join(c.DirToStorage, "processed_"+filename+"."+c.FormatToConvert)

	// Define quality settings
	var crf int
	var preset string
	var maxrate string

	switch c.Quality {
	case 1:
		crf = 28
		preset = "slow"
		maxrate = "1M"
	case 2:
		crf = 26
		preset = "medium"
		maxrate = "1.5M"
	case 3:
		crf = 23
		preset = "medium"
		maxrate = "2M"
	case 4:
		crf = 20
		preset = "fast"
		maxrate = "2.5M"
	case 5:
		crf = 18
		preset = "fast"
		maxrate = "3M"
	default:
		return "", fmt.Errorf("unsupported quality setting: %d", c.Quality)
	}

	// Check if the format is webm and use the convertToWebm function
	if c.FormatToConvert == "webm" {
		webmData, err := convertToWebm(tempPath, crf, c.Width, c.Height)
		if err != nil {
			return "", fmt.Errorf("error converting to webm: %w", err)
		}
		if err := os.WriteFile(destPath, webmData, 0644); err != nil {
			return "", fmt.Errorf("error writing webm file: %w", err)
		}
	} else {
		// Use ffmpeg to resize, crop, and convert the video
		err = ffmpeg.Input(tempPath).
			Filter("scale", ffmpeg.Args{fmt.Sprintf("iw*min(%d/iw\\,%d/ih):ih*min(%d/iw\\,%d/ih)", c.Width, c.Height, c.Width, c.Height)}).
			Filter("pad", ffmpeg.Args{fmt.Sprintf("%d:%d:(%d-iw)/2:(%d-ih)/2", c.Width, c.Height, c.Width, c.Height)}).
			Output(destPath, ffmpeg.KwArgs{
				"c:v":     "libx264",
				"crf":     crf,
				"preset":  preset,
				"maxrate": maxrate,
				"bufsize": "2M",
				"b:v":     "1M",
			}).
			OverWriteOutput().
			Run()

		if err != nil {
			return "", fmt.Errorf("error processing video: %w", err)
		}
	}

	return destPath, nil
}

// deleteVideo deletes the video from the directory
func (c *VideoConfig) deleteVideo(reqFilePath ...string) error {
	var filePath string
	if len(reqFilePath) > 0 {
		filePath = reqFilePath[0]
	} else {
		filePath = c.DirToStorage + c.FileName
	}

	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("failed to delete video: %w", err)
	}
	return nil
}
