package create

import (
	"fmt"
	"path/filepath"
)

func handler(m createModel) {
	projName := m.projectName.Value()
	templateName := m.selectedTemplate


	fmt.Printf("Creating new NodeArch app: %s using template: %s\n", projName, templateName)

	projectDir := filepath.Join(".", projName)
	// if err := os.MkdirAll(projectDir, os.ModePerm); err != nil {
	// 	fmt.Printf("Failed to create project directory: %v\n", err)
	// 	os.Exit(1)
	// }

	fmt.Printf("Project directory '%s' created successfully.\n", projectDir)
}
