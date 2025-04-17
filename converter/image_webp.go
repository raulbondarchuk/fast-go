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
