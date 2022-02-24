package graph

import (
	"bufio"
	"fmt"
	"os"
	"pluralith/pkg/ux"

	"github.com/spf13/pflag"
)

func GetDiagramValues(flags *pflag.FlagSet) (map[string]string, error) {
	functionName := "GetDiagramValues"

	diagramValues := make(map[string]string)

	// Get variable values from flags if present
	diagramValues["Title"], _ = flags.GetString("title")
	diagramValues["Author"], _ = flags.GetString("author")
	diagramValues["Version"], _ = flags.GetString("version")

	// Print UX head
	ux.PrintFormatted("⠿", []string{"blue", "bold"})
	if diagramValues["Title"] == "" && diagramValues["Author"] == "" && diagramValues["Version"] == "" {
		fmt.Println(" Exporting Diagram ⇢ Specify details below")
	} else {
		fmt.Println(" Exporting Diagram ⇢ Details taken from flags")
	}

	// Read all missing diagram values from stdin
	for key, _ := range diagramValues {
		if diagramValues[key] == "" {
			ux.PrintFormatted("  →", []string{"blue", "bold"})
			fmt.Printf(" %s: ", key)

			// Create scanner
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				diagramValues[key] = scanner.Text()
			}
			if scanErr := scanner.Err(); scanErr != nil {
				return diagramValues, fmt.Errorf("scanning input failed -> %v: %w", functionName, scanErr)
			}
		}
	}

	return diagramValues, nil
}
