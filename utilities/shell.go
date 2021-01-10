package utilities

import (
	"encoding/json"
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"github.com/kr/pretty"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"

	"gopkg.in/d4l3k/messagediff.v1"
)

// Formats an interface into a YAML byte array
func RenderYaml(input interface{}) []byte {
	var output, err = yaml.Marshal(&input)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return output
}

// Formats an interface into a JSON byte array
func RenderJSON(input interface{}) []byte {
	var output, err = json.MarshalIndent(&input, "", " ")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return output
}

// Determines what format to display to console
func Console(format string, input interface{}) []byte {
	switch {
	// Determine if JSON Flag Set
	case format == "json":
		results := RenderJSON(input)
		return results

	// Determine if YAML Flag Set
	case format == "yaml":
		results := RenderYaml(input)
		return results

	// Default If unable to determine Output
	default:
		log.Fatal("UnKnown output format: ", format)
		return nil
	}
}

// Takes a byte array and exports it to file
func RenderFile(input []byte, name, format string) {
	filename := fmt.Sprintf("%s.%s", name, format)
	ioutil.WriteFile(filename, input, 0644)
	return
}

// Prints to screen
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

	if diff {
		// Read Config File
		// RenderInput(&cfgFile)

	}
	input, _ := flags.GetString("input")

	// Concatenate []byte for single display
	for _, metadata := range metadatas {
		display := Console(format, metadata.Payload)
		if diff {
			diff, equal := messagediff.PrettyDiff(metadata.Payload, input)
			pretty.Println(diff, equal)

		} else {
			results = append(results, display...)
		}

	} // end for loop

	// Print Results to console
	pretty.Println(string(results))

	// return complete data to console
	if export {
		RenderFile(results, name, format)
	}

	// Display Verbose HTTP Information
	if verbose {
		for _, metadata := range metadatas {
			RenderVerboseOutput(metadata)
		} // end for loop
	}

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

func ReadConfigFile(cmd *cobra.Command, configFile interface{}) (interface{}, error) {
	v := viper.New()


	name, _ := cmd.Flags().GetString("input")
	// Set the base name of the config file, without the file extension.
	v.SetConfigName(name)
	v.SetConfigFile(name)
	// Set as many paths as you like where viper should look for the
	// config file. We are only looking in the current working directory.
	v.AddConfigPath(".")

	// Attempt to read the config file, gracefully ignoring errors
	// caused by a config file not being found. Return an error
	// if we cannot parse the config file.
	if err := v.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}

	// marshal config into interface
	err := v.Unmarshal(&configFile)
	if err != nil {
		log.Printf("Unable to decode into struct, %v", err)
	}

	return configFile, nil
}