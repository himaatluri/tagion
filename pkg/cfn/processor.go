package cfn

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/himaatluri/tagion/pkg/types"
	"github.com/himaatluri/tagion/pkg/utils"
	"gopkg.in/yaml.v3"
)

func ProcessDirectory(dirPath string, config types.Config) error {
	var templates []utils.TemplateStatus
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".yml") || strings.HasSuffix(path, ".json") {
			status, err := AnalyzeTemplate(path)
			if err != nil {
				return err
			}
			templates = append(templates, status)
		}
		return nil
	})

	if err != nil {
		return err
	}

	// Display changes and get confirmation
	utils.DisplayTemplateChanges(templates)

	if !utils.ConfirmChanges() {
		fmt.Println("Operation cancelled")
		return nil
	}

	for _, template := range templates {
		if template.Modifiable {
			if err := ProcessTemplate(template.Path, config); err != nil {
				return err
			}
		}
	}

	return nil
}

func AnalyzeTemplate(templatePath string) (utils.TemplateStatus, error) {
	status := utils.TemplateStatus{Path: templatePath}

	data, err := os.ReadFile(templatePath)
	if err != nil {
		return status, err
	}

	var template types.CloudFormationTemplate
	if strings.HasSuffix(templatePath, ".json") {
		err = json.Unmarshal(data, &template)
	} else {
		err = yaml.Unmarshal(data, &template)
	}
	if err != nil {
		return status, err
	}

	status.Resources = len(template.Resources)
	status.Modifiable = false
	status.HasTags = false

	for _, resource := range template.Resources {
		if utils.SupportsTags(resource.Type) {
			if _, hasTags := resource.Properties["Tags"]; hasTags {
				status.HasTags = true
			} else {
				status.Modifiable = true
			}
		}
	}

	return status, nil
}

func ProcessTemplate(templatePath string, config types.Config) error {
	data, err := os.ReadFile(templatePath)
	if err != nil {
		return err
	}

	var template types.CloudFormationTemplate

	if strings.HasSuffix(templatePath, ".json") {
		err = json.Unmarshal(data, &template)
	} else {
		err = yaml.Unmarshal(data, &template)
	}
	if err != nil {
		return err
	}

	modified := processResources(&template, config)

	if modified {
		if err := utils.WriteTemplateToFile(templatePath, template); err != nil {
			return err
		}
		fmt.Printf("Updated template: %s\n", templatePath)
	}

	return nil
}

func processResources(template *types.CloudFormationTemplate, config types.Config) bool {
	modified := false
	for resourceName, resource := range template.Resources {
		if resource.Properties == nil {
			resource.Properties = make(map[string]interface{})
		}

		if utils.SupportsTags(resource.Type) {
			tags, hasTags := resource.Properties["Tags"]
			if !hasTags {
				newTags := utils.CreateNewTags(config.Tags)
				resource.Properties["Tags"] = newTags
				template.Resources[resourceName] = resource
				modified = true
			} else {
				existingTags := utils.ConvertToTagsArray(tags)
				if updatedTags := utils.MergeTags(existingTags, config.Tags); updatedTags != nil {
					resource.Properties["Tags"] = updatedTags
					template.Resources[resourceName] = resource
					modified = true
				}
			}
		}
	}
	return modified
}
