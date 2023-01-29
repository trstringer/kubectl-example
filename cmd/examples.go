package cmd

import (
	"embed"
	"fmt"
	"path/filepath"
	"sort"
	"strings"
)

//go:embed examples/*
var examples embed.FS

const examplesDir string = "examples"

func resourceNameFromFile(filename string) string {
	fileParts := strings.Split(filename, "/")
	fileName := strings.Split(fileParts[len(fileParts)-1], ".")[0]
	return fileName
}

func resourcePathFromDir(dir string) string {
	dirParts := strings.Split(dir, "/")
	if dirParts[0] != examplesDir {
		return dir
	}
	return strings.Join(dirParts[1:], "/")
}

func getResourcesByDir(dir string) (map[string][]string, error) {
	resourceMap := make(map[string][]string)
	files, err := examples.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			subResourceMap, err := getResourcesByDir(filepath.Join(dir, file.Name()))
			if err != nil {
				return nil, fmt.Errorf("error reading dir %s: %v", file.Name(), err)
			}
			for resource, dirs := range subResourceMap {
				resourceMap[resource] = append(resourceMap[resource], dirs...)
			}
		} else {
			resourceName := resourceNameFromFile(file.Name())
			resourceMap[resourceName] = append(resourceMap[resourceName], resourcePathFromDir(dir))
		}
	}

	return resourceMap, nil
}

func getResources() ([]string, error) {
	resourceMap, err := getResourcesByDir(examplesDir)
	if err != nil {
		return nil, err
	}

	resourceList := []string{}
	for resource, dirs := range resourceMap {
		if len(dirs) == 1 {
			resourceList = append(resourceList, resource)
			continue
		}

		sort.Strings(dirs)
		resourceList = append(resourceList, fmt.Sprintf("%s %v", resource, dirs))
	}

	sort.Strings(resourceList)
	return resourceList, nil
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
	resourceMap, err := getResourcesByDir(examplesDir)
	if err != nil {
		return nil, fmt.Errorf("error getting resources by dir: %w", err)
	}

	for _, name := range names {
		nameParts := strings.Split(name, "/")
		kind := nameParts[len(nameParts)-1]
		groupVersion := strings.Join(nameParts[:len(nameParts)-1], "/")

		kindGroupVersions, found := resourceMap[kind]
		if !found {
			return nil, fmt.Errorf("kind %s not found", name)
		}
		if len(kindGroupVersions) > 1 && groupVersion == "" {
			return nil, fmt.Errorf("ambigious kind %s, specify a group version prefix: %v", kind, kindGroupVersions)
		}

		if groupVersion == "" && len(kindGroupVersions) == 1 {
			groupVersion = kindGroupVersions[0]
		}

		definitionPath := filepath.Join(examplesDir, groupVersion, kind)
		definitionPath += ".yaml"
		contents, err := examples.ReadFile(definitionPath)
		if err != nil {
			return nil, fmt.Errorf("error reading path %s: %w", definitionPath, err)
		}
		definitions = append(definitions, string(contents))
	}

	return definitions, nil
}
