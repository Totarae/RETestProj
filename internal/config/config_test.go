package config

import (
	"os"
	"path/filepath"
	"testing"
)

func writeTempFile(t *testing.T, content string) string {
	t.Helper()
	tmpDir := t.TempDir()
	file := filepath.Join(tmpDir, "packs.json")
	if err := os.WriteFile(file, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}
	return file
}

func TestLoadFromFile_Success(t *testing.T) {
	file := writeTempFile(t, `[250, 500, 1000]`)
	conf := &Config{}
	err := conf.LoadPackSizesFromFile(file)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	expected := []int{250, 500, 1000}
	actual := conf.GetPackSizes()
	for i, v := range expected {
		if actual[i] != v {
			t.Errorf("expected %d at index %d, got %d", v, i, actual[i])
		}
	}
}

func TestLoadFromFile_FileNotFound(t *testing.T) {
	conf := &Config{}
	err := conf.LoadPackSizesFromFile("nonexistent.json")
	if err == nil {
		t.Fatal("expected error for nonexistent file, got nil")
	}
}

func TestLoadFromFile_InvalidJSON(t *testing.T) {
	file := writeTempFile(t, `not a json`)
	conf := &Config{}
	err := conf.LoadPackSizesFromFile(file)
	if err == nil {
		t.Fatal("expected JSON parsing error, got nil")
	}
}

func TestLoadFromFile_EmptyArray(t *testing.T) {
	file := writeTempFile(t, `[]`)
	conf := &Config{}
	err := conf.LoadPackSizesFromFile(file)
	if err == nil {
		t.Fatal("expected error for empty pack sizes, got nil")
	}
}

func TestLoadFromFile_WrongType(t *testing.T) {
	file := writeTempFile(t, `["a", "b", "c"]`)
	conf := &Config{}
	err := conf.LoadPackSizesFromFile(file)
	if err == nil {
		t.Fatal("expected error for wrong data types, got nil")
	}
}
