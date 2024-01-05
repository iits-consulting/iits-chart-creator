package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

const bluePrintPath = "blueprints/templates"

func appendToFile(source, destination string) error {
	// Read the source file
	srcFile, err := os.ReadFile(source)
	if err != nil {
		return err
	}

	// Check if the source file is not empty
	if len(srcFile) == 0 {
		return fmt.Errorf("source file is empty")
	}

	// Open the destination file in append mode
	destFile, err := os.OpenFile(destination, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Write a newline to the destination file
	if _, err = destFile.Write([]byte("\n")); err != nil {
		return err
	}

	// Write the source file content to the destination file
	if _, err = destFile.Write(srcFile); err != nil {
		return err
	}

	fmt.Printf("Appended source file %s to destination file %s successfully.\n", source, destination)

	return nil
}

func mergeValuesFiles(cmd *cobra.Command, chartType string) error {
	folderName, _ := getCurrentFolderName()
	pluginDir := os.Getenv("HELM_PLUGIN_DIR")
	valuesSource := filepath.Join(pluginDir, bluePrintPath, chartType, cmd.Name()+"-values.yaml")

	if _, err := os.Stat(valuesSource); os.IsNotExist(err) {
		return err
	}

	tempPath := "./" + cmd.Name() + "-values.yaml"
	err := ReplaceInFile(valuesSource, "myService", folderName, tempPath)
	if err != nil {
		return err
	}

	err = appendToFile(tempPath, "./values.yaml")
	os.Remove(tempPath)

	if err != nil {
		return err
	}

	fmt.Println("values.yaml file updated successfully.")
	return nil
}

// checkAndCreateResource checks files and creates the required resources
func checkAndCreateResource(cmd *cobra.Command, kubernetesResourceName string, chartType string) {
	err := checkFileExist("./Chart.yaml")
	if err != nil {
		fmt.Println("You can only create a kubernetesResourceName if you are inside a helm chart.")
		return
	}
	folderName, err := getCurrentFolderName()
	if err != nil {
		fmt.Printf("Error getting current folder name: %v\n", err)
		return
	}
	pluginDir := os.Getenv("HELM_PLUGIN_DIR")
	sourcePath := filepath.Join(pluginDir, bluePrintPath, chartType, kubernetesResourceName)
	destinationPath := fmt.Sprintf("./templates/%s", kubernetesResourceName)
	err = ReplaceInFile(sourcePath, "myService", folderName, destinationPath)
	if err != nil {
		fmt.Printf("Error creating %s file: %v\n", kubernetesResourceName, err)
		return
	}
	fmt.Printf("%s file created successfully.\n", kubernetesResourceName)

	err2 := mergeValuesFiles(cmd, chartType)
	if err2 != nil {
		fmt.Printf("Error merging values files: %v\n", err2)
		return
	}

	fmt.Println("Please check the created file and replace all lines carrying 'REPLACE_ME' comment as needed.")
}
