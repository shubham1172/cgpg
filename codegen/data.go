package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// readJsonFile reads a JSON file and returns the result as a map.
func readJsonFile(filePath string) (map[string]interface{}, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(file, &jsonMap)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return jsonMap, nil
}

// readJsonFiles reads all .json files in the given directory,
// combines them into a single JSON object,
// and returns the result as a map.
func readJsonFiles(path string) (map[string]interface{}, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	jsonMap := make(map[string]interface{})
	for _, file := range files {
		// Skip non-JSON files.
		if !strings.HasSuffix(file.Name(), ".json") {
			continue
		}

		// Read the file.
		filePath := path + "/" + file.Name()
		fileJson, err := readJsonFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read file: %w", err)
		}

		// Merge the file's JSON into the result.
		for key, value := range fileJson {
			if _, ok := jsonMap[key]; ok {
				return nil, fmt.Errorf("key %s already exists, cannot merge %s", key, filePath)
			}

			jsonMap[key] = value
		}
	}

	return jsonMap, nil
}
