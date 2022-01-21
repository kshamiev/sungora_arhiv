package log

import "github.com/sirupsen/logrus"

type Config struct {
	Title     string       `yaml:"title" json:"title" toml:"title"`             // no field "title" if value is empty
	Output    string       `yaml:"output" json:"output" toml:"output"`          // enum (stdout | filePathRelative)
	Formatter string       `yaml:"formatter" json:"formatter" toml:"formatter"` // enum (json|text)
	Level     logrus.Level `yaml:"level" json:"level" toml:"level"`             // enum (error|warning|info|debug|trace)
}

const (
	titleField = "title"

	stdout = "stdout"

	formatterJSON = "json"
)
