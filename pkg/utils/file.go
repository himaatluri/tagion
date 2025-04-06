package utils

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/himaatluri/tagion/pkg/types"
	"gopkg.in/yaml.v3"
)

func LoadTagsConfig(filepath string) (types.Config, error) {
	var config types.Config
	data, err := os.ReadFile(filepath)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(data, &config)
	return config, err
}

func IsDirectory(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func WriteTemplateToFile(templatePath string, template types.CloudFormationTemplate) error {
	var (
		output []byte
		err    error
	)

	if strings.HasSuffix(templatePath, ".json") {
		output, err = json.MarshalIndent(template, "", "  ")
	} else {
		output, err = yaml.Marshal(template)
	}
	if err != nil {
		return err
	}

	return os.WriteFile(templatePath, output, 0644)
}
