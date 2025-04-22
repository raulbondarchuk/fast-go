package converter

import (
	"path/filepath"
	"strings"
)

const (
	// List of supported image formats
	PNG  = "png"
	JPEG = "jpg"
	JPG  = "jpeg"
	WEBP = "webp"
	JFIF = "jfif"
	// AVIF = "avif"
	// HEIC = "heic"

	// List of supported audio formats
	MP3  = "mp3"
	M4A  = "m4a"
	OPUS = "opus"
	WAV  = "wav"

	// List of supported video formats
	MP4  = "mp4"
	WEBM = "webm"
)

var supportedFormatsImage = []string{PNG, JPEG, JPG, WEBP, JFIF}
var supportedFormatsAudio = []string{MP3, M4A, OPUS, WAV}
var supportedFormatsVideo = []string{MP4, WEBM}

// FileType represents the type of file
type FileType int

const (
	Unknown FileType = iota // 5
	Image                   // 1
	Video                   // 2
	Audio                   // 3
	JSON                    // 4
)

// DetermineFileType determines the type of a file based on its extension
func DetermineFileType(fileName string) FileType {
	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(fileName), "."))

	if contains(supportedFormatsImage, ext) {
		return Image
	} else if contains(supportedFormatsVideo, ext) {
		return Video
	} else if contains(supportedFormatsAudio, ext) {
		return Audio
	} else if ext == "json" {
		return JSON
	} else {
		return Unknown
	}
}

// contains checks if a slice contains a specific string
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
