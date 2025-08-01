package jsonloader

import (
	"encoding/json"
	"fmt"
	"os"
)

type Loader struct {
	path string
}

func New(path string) *Loader{
	return &Loader{
		path: path,
	}
}

func (l *Loader) LoadAndParse() (map[string]any, error) {
	content, err := os.ReadFile(l.path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file : %w", err)
	}

	var parsedContent map[string]any
	err = json.Unmarshal(content,&parsedContent)
	if err != nil {
		return nil, fmt.Errorf("failed to pars json file : %w", err)
	}

	return parsedContent, nil
}
