package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/himaatluri/tagion/pkg/cfn"
	"github.com/himaatluri/tagion/pkg/utils"
)

func main() {
	tagsFile := flag.String("tags", "", "Path to JSON file containing tags")
	cfnPath := flag.String("path", "", "Path to CloudFormation template or directory")
	flag.Parse()

	if *tagsFile == "" || *cfnPath == "" {
		flag.Usage()
		os.Exit(1)
	}

	config, err := utils.LoadTagsConfig(*tagsFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading tags config: %v\n", err)
		os.Exit(1)
	}

	if utils.IsDirectory(*cfnPath) {
		err = cfn.ProcessDirectory(*cfnPath, config)
	} else {
		status, err := cfn.AnalyzeTemplate(*cfnPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error analyzing template: %v\n", err)
			os.Exit(1)
		}

		utils.DisplayTemplateChanges([]utils.TemplateStatus{status})

		if !utils.ConfirmChanges() {
			fmt.Println("Operation cancelled")
			return
		}

		if status.Modifiable {
			err = cfn.ProcessTemplate(*cfnPath, config)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error processing template: %v\n", err)
			}
		}
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error processing templates: %v\n", err)
		os.Exit(1)
	}
}
