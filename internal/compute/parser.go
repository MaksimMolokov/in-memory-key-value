package compute

import (
	"errors"
	"regexp"
	"strings"

	"db/pkg/compute"
	"go.uber.org/zap"
)

var (
	ErrInvalidCommand = errors.New("неверная команда")
	ErrInvalidArgs    = errors.New("неверные аргументы")
)

// parser реализует интерфейс Parser
type parser struct {
	logger *zap.Logger
}

// NewParser создает новый экземпляр парсера
func NewParser(logger *zap.Logger) *parser {
	return &parser{
		logger: logger,
	}
}

// Parse разбирает входную строку и возвращает соответствующую команду
func (p *parser) Parse(input string) (compute.Command, error) {
	p.logger.Info("разбор команды",
		zap.String("input", input))

	parts := strings.Fields(input)
	if len(parts) == 0 {
		return nil, ErrInvalidCommand
	}

	cmdType := compute.CommandType(strings.ToUpper(parts[0]))
	switch cmdType {
	case compute.SetCommandType:
		if len(parts) != 3 {
			return nil, ErrInvalidArgs
		}
		return &compute.SetCommand{
			Key:   parts[1],
			Value: parts[2],
		}, nil
	case compute.GetCommandType:
		if len(parts) != 2 {
			return nil, ErrInvalidArgs
		}
		return &compute.GetCommand{
			Key: parts[1],
		}, nil
	case compute.DelCommandType:
		if len(parts) != 2 {
			return nil, ErrInvalidArgs
		}
		return &compute.DelCommand{
			Key: parts[1],
		}, nil
	default:
		return nil, ErrInvalidCommand
	}
}

// isValidArgument проверяет, соответствует ли аргумент требованиям
func isValidArgument(arg string) bool {
	matched, _ := regexp.MatchString(`^[\w\-\.]+$`, arg)
	return matched
} 