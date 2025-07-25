package jsonloader

import (
	"encoding/json"
	"fmt"
	"os"
)

func Load(path string) (string, error){
	content, err := os.ReadFile(path)
	if err != nil {
		return "" , fmt.Errorf("faild to read file from path : %w", err)
	}

	return string(content), nil
}

func Parse(content string) (map[string]any, error) {
	var parsedJson map[string]any
	err := json.Unmarshal([]byte(content), &parsedJson)
	if err != nil {
		return nil, fmt.Errorf("faild to pars json file : %w", err)
	}

	return parsedJson, nil
}

