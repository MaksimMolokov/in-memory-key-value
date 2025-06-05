package compute

// Домашнее задание: комментарий 1
// Command представляет собой интерфейс для всех команд
type Command interface {
	Execute() error
}

// Домашнее задание: комментарий 2
// Parser представляет собой интерфейс для парсера команд
type Parser interface {
	Parse(input string) (Command, error)
}

// CommandType представляет тип команды
type CommandType string

const (
	SetCommandType CommandType = "SET"
	GetCommandType CommandType = "GET"
	DelCommandType CommandType = "DEL"
)

// SetCommand представляет команду SET
type SetCommand struct {
	Key   string
	Value string
}

func (c *SetCommand) Execute() error {
	return nil
}

// GetCommand представляет команду GET
type GetCommand struct {
	Key string
}

func (c *GetCommand) Execute() error {
	return nil
}

// DelCommand представляет команду DEL
type DelCommand struct {
	Key string
}

func (c *DelCommand) Execute() error {
	return nil
}
