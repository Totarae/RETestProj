package config

import (
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"sync"
)

type Config struct {
	mu    sync.RWMutex
	Sizes []int
}

// LoadPackSizes loads pack sizes
func (c *Config) LoadPackSizesFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read pack sizes file: %w", err)
	}

	var sizes []int
	if err := json.Unmarshal(data, &sizes); err != nil {
		return fmt.Errorf("failed to parse pack sizes JSON: %w", err)
	}

	if len(sizes) == 0 {
		return fmt.Errorf("no pack sizes provided")
	}

	c.mu.Lock()
	c.Sizes = sizes
	c.mu.Unlock()

	fmt.Println("Config reloaded:", sizes)
	return nil
}

// GetPackSizes returns a copy of current config.
func (c *Config) GetPackSizes() []int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	// Return a copy to prevent external modification
	safeCopy := make([]int, len(c.Sizes))
	copy(safeCopy, c.Sizes)
	return safeCopy
}

// ChangeWatcher
func (c *Config) WatchConfigFile(filename string) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("watcher failed: %w", err)
	}
	err = watcher.Add(filename)
	if err != nil {
		return fmt.Errorf("watch file: %w", err)
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					if err := c.LoadPackSizesFromFile(filename); err != nil {
						log.Println("Failed to reload config:", err)
					}
				}
			case err := <-watcher.Errors:
				log.Println("Watcher error:", err)
			}
		}
	}()

	return nil
}
