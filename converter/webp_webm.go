package converter

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"os/exec"
)

// convertToWebp converts an image to WebP format
func convertToWebp(img any, quality uint8) ([]byte, error) {
	cmd := exec.Command("cwebp", "-q", fmt.Sprint(quality), "-o", "-", "--", "-")
	var out bytes.Buffer
	cmd.Stdout = &out
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	go func() {
		defer stdin.Close()
		_ = png.Encode(stdin, img.(image.Image))
	}()
	if err := cmd.Wait(); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

// convertToWebm converts a video to WebM format with specified dimensions and quality
func convertToWebm(inputVideoPath string, quality int, width int, height int) ([]byte, error) {
	var crf int
	var maxrate string

	// Define quality settings based on the quality parameter
	switch quality {
	case 1:
		crf = 32
		maxrate = "500k"
	case 2:
		crf = 28
		maxrate = "1M"
	case 3:
		crf = 24
		maxrate = "1.5M"
	case 4:
		crf = 20
		maxrate = "2M"
	case 5:
		crf = 16
		maxrate = "2.5M"
	default:
		crf = 28
		maxrate = "1M"
	}

	cmd := exec.Command("ffmpeg", "-i", inputVideoPath, "-vf", fmt.Sprintf("scale=%d:%d", width, height), "-c:v", "libvpx", "-crf", fmt.Sprint(crf), "-b:v", maxrate, "-c:a", "libvorbis", "-f", "webm", "-")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	if err := cmd.Wait(); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}
