package compute

import (
	"testing"

	"go.uber.org/zap"
)

func TestParser(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	parser := NewParser(logger)

	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "valid SET command",
			input:   "SET key value",
			wantErr: false,
		},
		{
			name:    "valid GET command",
			input:   "GET key",
			wantErr: false,
		},
		{
			name:    "valid DEL command",
			input:   "DEL key",
			wantErr: false,
		},
		{
			name:    "invalid command",
			input:   "INVALID",
			wantErr: true,
		},
		{
			name:    "SET with invalid args",
			input:   "SET key",
			wantErr: true,
		},
		{
			name:    "GET with invalid args",
			input:   "GET",
			wantErr: true,
		},
		{
			name:    "empty input",
			input:   "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd, err := parser.Parse(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && cmd == nil {
				t.Error("Parse() returned nil command when no error expected")
			}
		})
	}
} 