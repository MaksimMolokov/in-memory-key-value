package storage

import (
	"testing"

	"go.uber.org/zap"
)

func TestEngine(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	engine := NewEngine(logger)

	t.Run("Set and Get", func(t *testing.T) {
		key := "test_key"
		value := "test_value"

		// Проверяем, что ключ не существует
		_, err := engine.Get(key)
		if err != ErrKeyNotFound {
			t.Errorf("Get() error = %v, want %v", err, ErrKeyNotFound)
		}

		// Устанавливаем значение
		err = engine.Set(key, value)
		if err != nil {
			t.Errorf("Set() error = %v", err)
		}

		// Проверяем, что значение установлено
		got, err := engine.Get(key)
		if err != nil {
			t.Errorf("Get() error = %v", err)
		}
		if got != value {
			t.Errorf("Get() = %v, want %v", got, value)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		key := "delete_key"
		value := "delete_value"

		// Устанавливаем значение
		err := engine.Set(key, value)
		if err != nil {
			t.Errorf("Set() error = %v", err)
		}

		// Удаляем значение
		err = engine.Delete(key)
		if err != nil {
			t.Errorf("Delete() error = %v", err)
		}

		// Проверяем, что значение удалено
		_, err = engine.Get(key)
		if err != ErrKeyNotFound {
			t.Errorf("Get() error = %v, want %v", err, ErrKeyNotFound)
		}

		// Проверяем удаление несуществующего ключа
		err = engine.Delete("non_existent_key")
		if err != ErrKeyNotFound {
			t.Errorf("Delete() error = %v, want %v", err, ErrKeyNotFound)
		}
	})
} 