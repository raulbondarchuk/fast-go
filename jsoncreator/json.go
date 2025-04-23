package jsoncreator

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tidwall/sjson"
)

// CreateJSONFile creates a JSON file.
//
//	tempDir     – base directory (e.g., "/tmp").
//	fileName    – final file name (with ".json" extension).
//	data        – content to write (key-value pair).
//	subPath     – unlimited number of intermediate directories
//	              (companyID, requestID, any nesting).
//
// Example:
//
//	path, _ := CreateJSONFile("data.json", map[string]interface{}{"key": "value"}, "/tmp", "123456", "req42")
func CreateJSONFile(fileName string, data map[string]interface{},
	tempDir string, subPath ...string) (string, error) {

	// --- Format JSON ---------------------------

	var jsonBytes []byte
	var err error

	if len(data) == 0 {
		jsonBytes = []byte("{}")
	} else {
		// fast-path: regular marshal
		jsonBytes, err = json.Marshal(data)
		if err != nil {
			return "", fmt.Errorf("marshal json: %w", err)
		}
	}

	// --- Final path ----------------------------

	fullDir := filepath.Join(append([]string{tempDir}, subPath...)...)
	filePath := filepath.Join(fullDir, fileName)

	if err := os.MkdirAll(fullDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("mkdir %q: %w", fullDir, err)
	}

	if err := os.WriteFile(filePath, jsonBytes, 0o644); err != nil {
		return "", fmt.Errorf("write %q: %w", filePath, err)
	}
	return filePath, nil
}

// SetJSONField changes/adds a field in an existing JSON file
// using github.com/tidwall/sjson (if you need the path "a.b.c").
//
// Example:
//
//	_ = SetJSONField(path, "metadata.size", 123)
func SetJSONField(filePath, path string, value interface{}) error {
	raw, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("read %q: %w", filePath, err)
	}

	modified, err := sjson.SetBytes(raw, path, value)
	if err != nil {
		return fmt.Errorf("sjson set: %w", err)
	}

	if err := os.WriteFile(filePath, modified, 0o644); err != nil {
		return fmt.Errorf("rewrite %q: %w", filePath, err)
	}
	return nil
}

// RemoveJSONFile deletes a specific file.
func RemoveJSONFile(filePath string) error {
	if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("remove %q: %w", filePath, err)
	}
	return nil
}

// RemoveTempTree deletes the entire subdirectory (companyID, requestID…).
//
//	subPath must be the same as the one passed to CreateJSONFile.
func RemoveTempTree(tempDir string, subPath ...string) error {
	dir := filepath.Join(append([]string{tempDir}, subPath...)...)
	if err := os.RemoveAll(dir); err != nil {
		return fmt.Errorf("remove dir %q: %w", dir, err)
	}
	return nil
}

// GetFileSize returns the size of the file in bytes.
func GetFileSize(filePath string) (int64, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return 0, fmt.Errorf("stat %q: %w", filePath, err)
	}
	return info.Size(), nil
}
