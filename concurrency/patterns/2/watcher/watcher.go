package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Config struct {
	ServerPort string `json:"server_port"`
	LogLevel   string `json:"log_level"`
	// Add other configuration fields
}

var config Config

func loadConfig(filePath string) error {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}
	if err := json.Unmarshal(file, &config); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}
	return nil
}

func watchConfig(ctx context.Context, filePath string) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Printf("Error stating config file: %v\n", err)
		return
	}
	lastModTime := fileInfo.ModTime()

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fileInfo, err := os.Stat(filePath)
			if err != nil {
				fmt.Printf("Error stating config file: %v\n", err)
				continue
			}

			if fileInfo.ModTime().After(lastModTime) {
				lastModTime = fileInfo.ModTime()
				fmt.Println("Config file changed, reloading...")

				if err := loadConfig(filePath); err != nil {
					fmt.Printf("Error reloading config: %v\n", err)
					continue
				}

				fmt.Printf("Reloaded successfully. New Config: %+v\n", config)
			}
		case <-ctx.Done():
			log.Println("Stopping periodic task:", ctx.Err())
			return
		}
	}
}

func main() {
	configFilePath := "config.json"

	if err := loadConfig(configFilePath); err != nil {
		fmt.Printf("Error loading initial config: %v\n", err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go watchConfig(ctx, configFilePath)

	<-ctx.Done()

	fmt.Println("Shutting down...")
}
