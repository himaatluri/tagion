package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

// ANSI color codes
const (
	colorReset  = "\033[0m"
	colorYellow = "\033[33m"
	colorGreen  = "\033[32m"
)

type TemplateStatus struct {
	Path       string
	HasTags    bool
	Modifiable bool
	Resources  int
}

func DisplayTemplateChanges(templates []TemplateStatus) {
	if len(templates) == 0 {
		fmt.Println("No templates found to process")
		return
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	// Set table style
	t.SetStyle(table.StyleRounded)

	// Columns
	t.AppendHeader(table.Row{"Template", "Resources", "Status"})

	// Rows
	var toModify int
	for _, tmpl := range templates {
		displayPath := tmpl.Path
		if cwd, err := os.Getwd(); err == nil {
			if rel, err := filepath.Rel(cwd, tmpl.Path); err == nil {
				displayPath = rel
			}
		}

		status := "No changes needed"
		if tmpl.Modifiable {
			status = colorYellow + "Will be modified" + colorReset
			toModify++
		} else if tmpl.HasTags {
			status = colorGreen + "Has tags" + colorReset
		}

		t.AppendRow(table.Row{displayPath, tmpl.Resources, status})
	}

	// footer
	t.AppendFooter(table.Row{"Total", len(templates),
		fmt.Sprintf("%sTemplates to modify: %d%s", colorYellow, toModify, colorReset)})

	t.SetColumnConfigs([]table.ColumnConfig{
		{
			Name:     "Template",
			WidthMax: 60,
		},
		{
			Name:     "Resources",
			Align:    text.AlignRight,
			WidthMax: 10,
		},
		{
			Name:     "Status",
			WidthMax: 20,
		},
	})

	t.Render()

	fmt.Println()
}

func ConfirmChanges() bool {
	fmt.Print("Do you want to proceed with these changes? [y/N]: ")
	var response string
	fmt.Scanln(&response)
	return strings.ToLower(response) == "y"
}
