package converter

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	MP3  = "mp3"
	M4A  = "m4a"
	OPUS = "opus"
	WAV  = "wav"
)

var supportedFormatsAudio = []string{MP3, M4A, OPUS, WAV, MP4}

type AudioConfig struct {
	FileName        string
	File            io.Reader
	Bitrate         int
	FormatToConvert string // mp3, m4a, opus, wav, mp4
	DirToStorage    string
}

func (c *AudioConfig) validateValues() error {
	if c.Bitrate < 64 || c.Bitrate > 320 {
		return fmt.Errorf("bitrate must be between 64 and 320 kbps")
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

func (c *AudioConfig) isFileExtensionSupported() bool {
	ext := strings.ToLower(filepath.Ext(c.FileName))
	for _, format := range supportedFormatsAudio {
		if strings.Contains(ext, format) {
			if format == MP4 {
				c.FormatToConvert = MP3
			}
			return true
		}
	}
	return false
}

func (c *AudioConfig) processAudio() (string, error) {
	if err := c.validateValues(); err != nil {
		return "", err
	}

	if err := os.MkdirAll(c.DirToStorage, 0755); err != nil {
		return "", fmt.Errorf("failed to create result dir: %w", err)
	}

	// Save original audio temporarily in dirResult
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

	// Prepare ffmpeg args based on format
	var args []string
	switch c.FormatToConvert {
	case "mp3":
		args = []string{"-i", tempPath, "-c:a", "libmp3lame", "-b:a", fmt.Sprintf("%dk", c.Bitrate), destPath}
	case "m4a":
		args = []string{"-i", tempPath, "-c:a", "aac", "-b:a", fmt.Sprintf("%dk", c.Bitrate), destPath}
	case "opus":
		args = []string{"-i", tempPath, "-c:a", "libopus", "-b:a", fmt.Sprintf("%dk", c.Bitrate), destPath}
	case "wav":
		// uncompressed WAV
		args = []string{"-i", tempPath, "-c:a", "pcm_s16le", destPath}
	default:
		return "", fmt.Errorf("unsupported conversion format: %s", c.FormatToConvert)
	}

	// Execute ffmpeg
	cmd := exec.Command("ffmpeg", args...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("error processing audio: %s, ffmpeg error: %s", err, stderr.String())
	}

	// Check if the file was created
	if _, err := os.Stat(destPath); os.IsNotExist(err) {
		return "", fmt.Errorf("file was not created: %s", destPath)
	}

	return destPath, nil
}

func (c *AudioConfig) deleteAudio(reqFilePath ...string) error {
	var filePath string
	if len(reqFilePath) > 0 {
		filePath = reqFilePath[0]
	} else {
		filePath = c.DirToStorage + c.FileName
	}

	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("failed to delete audio: %w", err)
	}
	return nil
}
