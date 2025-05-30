package storage

// Engine представляет собой интерфейс для хранилища данных
type Engine interface {
	// Set сохраняет значение по ключу
	Set(key, value string) error
	// Get возвращает значение по ключу
	Get(key string) (string, error)
	// Delete удаляет значение по ключу
	Delete(key string) error
} 