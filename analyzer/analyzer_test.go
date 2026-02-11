package analyzer

import (
	"testing"
)

func TestParseStringLiteral(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "simple quoted string",
			input:   `"hello world"`,
			want:    "hello world",
			wantErr: false,
		},
		{
			name:    "string with spaces",
			input:   `"database connection established"`,
			want:    "database connection established",
			wantErr: false,
		},
		{
			name:    "string with escaped quotes",
			input:   `"error: \"connection failed\""`,
			want:    `error: "connection failed"`,
			wantErr: false,
		},
		{
			name:    "empty string",
			input:   `""`,
			want:    "",
			wantErr: false,
		},
		{
			name:    "single character",
			input:   `"a"`,
			want:    "a",
			wantErr: false,
		},
		{
			name:    "string with special chars",
			input:   `"process-timeout: 30s"`,
			want:    "process-timeout: 30s",
			wantErr: false,
		},
		{
			name:    "too short string",
			input:   `"`,
			want:    "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseStringLiteral(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseStringLiteral() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseStringLiteral() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestGetPackagePath(t *testing.T) {
	tests := []struct {
		name string

	}{
		{
			name: "basic existence check",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {})
	}
}

func TestAnalyzerExists(t *testing.T) {
	if Analyzer == nil {
		t.Error("Analyzer should not be nil")
	}

	if Analyzer.Name != "loglint" {
		t.Errorf("Analyzer.Name = %q, want %q", Analyzer.Name, "loglint")
	}

	if Analyzer.Run == nil {
		t.Error("Analyzer.Run should not be nil")
	}
}

func TestNewAnalyzerFunction(t *testing.T) {
	analyzer := New()
	if analyzer == nil {
		t.Error("New() should return an Analyzer")
	}

	if analyzer.Name != "loglint" {
		t.Errorf("New() Analyzer.Name = %q, want %q", analyzer.Name, "loglint")
	}
}
