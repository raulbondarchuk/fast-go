package converter

import (
	"path/filepath"
	"strings"

	_ "golang.org/x/image/webp"
)

const (
	// List of supported image formats
	PNG  = "png"
	JPEG = "jpeg"
	JPG  = "jpg"
	WEBP = "webp" // VP8X is not supported, only VP8 and VP8L are allowed
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

	// List of supported JSON formats
	JSON = "json"
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
	Json                    // 4
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
	} else if ext == JSON {
		return Json
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
