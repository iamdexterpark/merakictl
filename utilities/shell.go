package utilities

import (
	"encoding/json"
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"github.com/kr/pretty"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"

	"gopkg.in/d4l3k/messagediff.v1"
)

// displayYaml - Formats an interface into a YAML byte array
func displayYaml(input interface{}) []byte {
	var output, err = yaml.Marshal(&input)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return output
}

// displayJSON - Formats an interface into a JSON byte array
func DisplayJSON(input interface{}) []byte {
	var output, err = json.MarshalIndent(&input, "", " ")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return output
}

// Console - Determines what format to display to console
func Console(format string, input interface{}) []byte {
	switch {
	// Determine if JSON Flag Set
	case format == "json":
		results := DisplayJSON(input)
		return results

	// Determine if YAML Flag Set
	case format == "yaml":
		results := displayYaml(input)
		return results

	// Default If unable to determine Output
	default:
		log.Fatal("UnKnown output format: ", format)
		return nil
	}
}

// ExportToFile - Takes a byte array and prints it to file
func ExportToFile(input []byte, name, format string) {
	filename := fmt.Sprintf("%s.%s", name, format)
	ioutil.WriteFile(filename, input, 0644)
	return
}

// TestDisplay
func Display(metadatas []api.Results, name string, flags *pflag.FlagSet) {
	var results []byte
	var format string
	verbose, _ := flags.GetBool("verbose")
	export, _ := flags.GetBool("export")
	jsonflag, _ := flags.GetBool("json")
	diff, _ := flags.GetBool("diff")

	if jsonflag {
		format = "json"
	} else {
		format = "yaml"
	}

	// config file for diff
	var cfgFile interface{}
	if diff {
		// Read Config File
		RenderInput(&cfgFile)
	}


	// Concatenate []byte for single display
	for _, metadata := range metadatas {
		display := Console(format, metadata.Payload)
		if diff {
			diff, equal := messagediff.PrettyDiff(metadata.Payload, cfgFile)
			pretty.Println(diff, equal)

		} else {
			results = append(results, display...)
		}

	} // end for loop

	// Print Results to console
	pretty.Println(string(results))

	// return complete data to console
	if export {
		ExportToFile(results, name, format)
	}

	// Display Verbose HTTP Information
	if verbose {
		for _, metadata := range metadatas {
			RenderVerboseOutput(metadata)
		} // end for loop
	}

}

func RenderInput(input interface{}) interface{} {

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&input)
	if err != nil {
		log.Fatal(errors.Wrap(err, "unmarshal config file"))
	}
	return input
}

func RenderVerboseOutput(metadata api.Results) {
	_, err := pretty.Println(metadata.Request)
	if err != nil {
		log.Fatal(err)
	}

	_, err = pretty.Println(metadata.Response)
	if err != nil {
		log.Fatal(err)
	}

}
