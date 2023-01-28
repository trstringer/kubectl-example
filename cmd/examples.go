package cmd

import (
	"embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed examples/*.yaml
var examples embed.FS

const examplesDir string = "examples"

func resourceNameFromFile(filename string) string {
	fileParts := strings.Split(filename, "/")
	fileName := strings.Split(fileParts[len(fileParts)-1], ".")[0]
	return fileName
}

func getResources() ([]string, error) {
	files, err := examples.ReadDir(examplesDir)
	if err != nil {
		return nil, err
	}
	resourceNames := []string{}

	for _, file := range files {
		resourceNames = append(resourceNames, resourceNameFromFile(file.Name()))
	}
	sort.Strings(resourceNames)

	return resourceNames, nil
}

func getResourcesByName(names []string) ([]string, error) {
	definitions := []string{}

	for _, name := range names {
		contents, err := examples.ReadFile(fmt.Sprintf("%s/%s.yaml", examplesDir, name))
		if err != nil {
			return nil, fmt.Errorf("error getting content for %s: %w", name, err)
		}
		definitions = append(definitions, string(contents))
	}

	return definitions, nil
}
