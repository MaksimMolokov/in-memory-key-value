package storage

import (
	"errors"
	"sync"

	"go.uber.org/zap"
)

var (
	ErrKeyNotFound = errors.New("ключ не найден")
)

// Домашнее задание: комментарий 1
// engine реализует интерфейс Engine
type engine struct {
	data   map[string]string
	logger *zap.Logger
	mu     sync.RWMutex
}

// NewEngine создает новый экземпляр движка
func NewEngine(logger *zap.Logger) *engine {
	return &engine{
		data:   make(map[string]string),
		logger: logger,
	}
}

// Set сохраняет значение по ключу
func (e *engine) Set(key, value string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.logger.Info("установка значения",
		zap.String("key", key),
		zap.String("value", value))

	e.data[key] = value
	return nil
}

// Get возвращает значение по ключу
func (e *engine) Get(key string) (string, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()

	e.logger.Info("получение значения",
		zap.String("key", key))

	value, exists := e.data[key]
	if !exists {
		return "", ErrKeyNotFound
	}

	return value, nil
}

// Delete удаляет значение по ключу
func (e *engine) Delete(key string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.logger.Info("удаление значения",
		zap.String("key", key))

	if _, exists := e.data[key]; !exists {
		return ErrKeyNotFound
	}

	delete(e.data, key)
	return nil
}

// Домашнее задание: комментарий 2
