package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Message struct {
	Token  string
	FileID string
	Data   string
}

type TokenValidator interface {
	IsValid(token string) bool
}

// SimpleTokenValidator implementation of TokenValidator
type SimpleTokenValidator struct {
	validTokens map[string]struct{}
}

func NewTokenValidator(tokens []string) *SimpleTokenValidator {
	validTokens := make(map[string]struct{}, len(tokens))
	for _, token := range tokens {
		validTokens[token] = struct{}{}
	}
	return &SimpleTokenValidator{validTokens: validTokens}
}

func (v *SimpleTokenValidator) IsValid(token string) bool {
	_, exists := v.validTokens[token]
	return exists
}

// MessageCache interface for message caching
type MessageCache interface {
	AddMessage(msg Message)
	FlushToFiles()
}

// SimpleMessageCache implementation of MessageCache
type SimpleMessageCache struct {
	mu        sync.RWMutex
	cache     map[string][]Message
	validator TokenValidator
}

func NewMessageCache(validator TokenValidator) *SimpleMessageCache {
	return &SimpleMessageCache{
		cache:     make(map[string][]Message),
		validator: validator,
	}
}

func (mc *SimpleMessageCache) AddMessage(msg Message) {
	if !mc.validator.IsValid(msg.Token) {
		return
	}

	mc.mu.Lock()
	defer mc.mu.Unlock()
	mc.cache[msg.FileID] = append(mc.cache[msg.FileID], msg)
}

func (mc *SimpleMessageCache) RetryWriteToFile(fileID string, messages []Message) error {
	const maxRetries = 3
	for attempt := 1; attempt <= maxRetries; attempt++ {
		file, err := os.OpenFile(fileID, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			fmt.Printf("Attempt %d: Error opening file %s: %v\n", attempt, fileID, err)
			if attempt == maxRetries {
				return err
			}
			time.Sleep(100 * time.Millisecond)
			continue
		}
		defer file.Close()

		for _, msg := range messages {
			if _, err := file.WriteString(msg.Data + "\n"); err != nil {
				fmt.Printf("Attempt %d: Error writing to file %s: %v\n", attempt, fileID, err)
				if attempt == maxRetries {
					return err
				}
				time.Sleep(100 * time.Millisecond)
				break
			}
		}
		return nil
	}
	return fmt.Errorf("failed to write to file %s after %d attempts", fileID, maxRetries)
}

func (mc *SimpleMessageCache) FlushToFiles() {
	mc.mu.Lock()
	defer mc.mu.Unlock()

	for fileID, messages := range mc.cache {
		if err := mc.RetryWriteToFile(fileID, messages); err != nil {
			log.Printf("Failed to flush messages to file %s: %v\n", fileID, err)
		} else {
			delete(mc.cache, fileID)
			log.Println("clear cache ", fileID)
		}
	}
}

// Worker struct for handling periodic cache flushing
type Worker struct {
	cache    MessageCache
	interval time.Duration
}

func NewWorker(cache MessageCache, interval time.Duration) *Worker {
	return &Worker{cache: cache, interval: interval}
}

func (w *Worker) Start(ctx context.Context, wg *sync.WaitGroup) {
	ticker := time.NewTicker(w.interval)
	defer ticker.Stop()
	defer wg.Done()

	for {
		select {
		case <-ticker.C:
			log.Println("start flushing to files")
			w.cache.FlushToFiles()
		case <-ctx.Done():
			// Graceful shutdown
			w.cache.FlushToFiles()
			return
		}
	}
}

// SimulateUsers simulates users sending messages
func SimulateUsers(channels []chan Message, tokens []string, numMessages int) {
	for i := 0; i < numMessages; i++ {
		for user, ch := range channels {
			msg := Message{
				Token:  tokens[user],
				FileID: fmt.Sprintf("file_%d.txt", user),
				Data:   fmt.Sprintf("Message %d from user %d", i, user),
			}
			ch <- msg
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// Configuration
	workerInterval := 1 * time.Second
	numUsers := 3
	numMessages := 10
	validTokens := []string{"token1", "token2", "token3"}

	validator := NewTokenValidator(validTokens)
	cache := NewMessageCache(validator)

	// Message channels for each user
	messageChannels := make([]chan Message, numUsers)
	for i := range messageChannels {
		messageChannels[i] = make(chan Message, 100)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup
	wg.Add(1)
	worker := NewWorker(cache, workerInterval)
	go worker.Start(ctx, &wg)

	for _, ch := range messageChannels {
		go func(mc chan Message) {
			for msg := range mc {
				cache.AddMessage(msg)
				log.Println("added to cache: ", msg)
			}
		}(ch)
	}

	// Simulate users sending messages
	go SimulateUsers(messageChannels, validTokens, numMessages)

	<-ctx.Done()

	wg.Wait()

	// Close all channels
	for _, ch := range messageChannels {
		close(ch)
	}

	fmt.Println("gracefully stutted down")
}
