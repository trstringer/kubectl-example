package cmd

import (
	"embed"
	"fmt"
	"path/filepath"
	"sort"
	"strings"
)

//go:embed examples/*.yaml
//go:embed examples/istio/*.yaml
var examples embed.FS

const examplesDir string = "examples"

func resourceNameFromFile(filename string) string {
	fileParts := strings.Split(filename, "/")
	fileName := strings.Split(fileParts[len(fileParts)-1], ".")[0]
	return fileName
}

func getResourcesByDir(dir string) ([]string, error) {
	files, err := examples.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	resourceNames := []string{}

	for _, file := range files {
		if file.IsDir() {
			subDirFiles, err := getResourcesByDir(filepath.Join(dir, file.Name()))
			if err != nil {
				return nil, fmt.Errorf("error reading dir %s: %v", file.Name(), err)
			}
			resourceNames = append(resourceNames, subDirFiles...)
		} else {
			resourceNames = append(resourceNames, resourceNameFromFile(file.Name()))
		}
	}
	sort.Strings(resourceNames)

	return resourceNames, nil
}

func getResources() ([]string, error) {
	return getResourcesByDir(examplesDir)
}

func findFileByName(dir, filename string) (string, error) {
	files, err := examples.ReadDir(dir)
	if err != nil {
		return "", fmt.Errorf("error reading dir %s: %v", dir, err)
	}

	for _, file := range files {
		if file.IsDir() {
			subDirSearch, err := findFileByName(filepath.Join(dir, file.Name()), filename)
			if err != nil {
				return "", err
			}
			if subDirSearch != "" {
				return subDirSearch, nil
			}
		} else if resourceNameFromFile(file.Name()) == filename {
			return filepath.Join(dir, file.Name()), nil
		}
	}

	return "", nil
}

func getResourcesByName(names []string) ([]string, error) {
	definitions := []string{}

	for _, name := range names {
		filename, err := findFileByName(examplesDir, name)
		if err != nil {
			return nil, err
		}
		if filename == "" {
			return nil, fmt.Errorf("resource %s not found", name)
		}
		contents, err := examples.ReadFile(filename)
		if err != nil {
			return nil, fmt.Errorf("error getting content for %s: %w", name, err)
		}
		definitions = append(definitions, string(contents))
	}

	return definitions, nil
}
