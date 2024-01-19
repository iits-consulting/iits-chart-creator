package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

const bluePrintPath = "blueprints/templates"

type DataInputForCreation struct {
	ManifestName string
	ChartType    string
	Version      string
	ReleaseName  string
}

func (d *DataInputForCreation) GenerateReplacements() map[string]string {
	return map[string]string{
		"myService":             d.ReleaseName,
		"CHART_CREATOR_VERSION": d.Version,
	}
}

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

func mergeValuesFiles(cmd *cobra.Command, dataInputForCreation DataInputForCreation) error {
	pluginDir := os.Getenv("HELM_PLUGIN_DIR")
	valuesSource := filepath.Join(pluginDir, bluePrintPath, dataInputForCreation.ChartType, cmd.Name()+"-values.yaml")

	if _, err := os.Stat(valuesSource); os.IsNotExist(err) {
		return err
	}

	tempPath := "./" + cmd.Name() + "-values.yaml"
	err := ReplaceInFile(valuesSource, dataInputForCreation.GenerateReplacements(), tempPath)
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
func checkAndCreateResource(cmd *cobra.Command, dataInputForCreation DataInputForCreation) {
	err := checkFileExist("./Chart.yaml")
	if err != nil {
		fmt.Println("You can only create a kubernetesResourceName if you are inside a helm chart.")
		return
	}
	pluginDir := os.Getenv("HELM_PLUGIN_DIR")
	manifestName := dataInputForCreation.ManifestName
	chartType := dataInputForCreation.ChartType
	sourcePath := filepath.Join(pluginDir, bluePrintPath, chartType, manifestName)
	destinationPath := fmt.Sprintf("./templates/%s", manifestName)

	err = ReplaceInFile(sourcePath, dataInputForCreation.GenerateReplacements(), destinationPath)
	if err != nil {
		fmt.Printf("Error creating %s file: %v\n", manifestName, err)
		return
	}
	fmt.Printf("%s file created successfully.\n", manifestName)

	err2 := mergeValuesFiles(cmd, dataInputForCreation)
	if err2 != nil {
		fmt.Printf("Error merging values files: %v\n", err2)
		return
	}

	fmt.Println("Please check the created file and replace all lines carrying 'REPLACE_ME' comment as needed.")
}
