package dataloader

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func LoadData(file string) string {

	// Try loading from cache first
	if content, err := loadFromCache(file); err == nil {
		return string(content)
	}

	if content, err := loadFileFromHTTP(file); err == nil {
		// Save to cache
		saveToCache(file, content)
		return string(content)
	}

	return ""
}

func loadFileFromHTTP(filePath string) ([]byte, error) {
	client := &http.Client{Timeout: 10}
	url := fmt.Sprintf("https://raw.githubusercontent.com/Ainaras/gostuff/main/data/%s", filePath)

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, resp.Status)
	}

	content, err := io.ReadAll(io.LimitReader(resp.Body, 1024*1024)) // Limit to 1MB
	if err != nil {
		return nil, err
	}

	return content, nil
}

func saveToCache(filePath string, content []byte) error {
	// Ensure directory exists
	if err := os.MkdirAll(getCacheDirPath(), 0755); err != nil {
		return fmt.Errorf("failed to create cache directory: %v", err)
	}

	cachePath := getCachePath(filePath)
	if err := os.WriteFile(cachePath, content, 0644); err != nil {
		return fmt.Errorf("failed to write cache file: %v", err)
	}

	return nil
}

func loadFromCache(file string) ([]byte, error) {
	cachePath := getCachePath(file)

	// Check if cache file exists
	if _, err := os.Stat(cachePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("cache file does not exist")
	}

	content, err := os.ReadFile(cachePath)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func ClearCache() error {
	return os.RemoveAll(getCacheDirPath())
}

func getCacheDirPath() string {
	return filepath.Join(os.TempDir(), "_gostuff_cache")
}

func getCachePath(file string) string {
	sum := md5.Sum([]byte(file))
	filenameEncoded := hex.EncodeToString(sum[:])
	return filepath.Join(getCacheDirPath(), filenameEncoded)
}
