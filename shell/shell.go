package shell

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
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
func displayJSON(input interface{}) []byte {
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
		results := displayJSON(input)
		fmt.Printf(string(results))
		return results

	// Determine if YAML Flag Set
	case format == "yaml":
		results := displayYaml(input)
		fmt.Printf(string(results))
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

func Display(payload, traceback interface{}, name string, flags  *pflag.FlagSet) {
	var results []byte
	var format string
	verbose, _ := flags.GetBool("verbose")
	export, _ := flags.GetBool("export")
	jsonflag, _ := flags.GetBool("json")

	if jsonflag {
		format = "json"
		results = Console(format, payload)
	} else {
		format = "yaml"
		results = Console(format, payload)
	}
	if verbose {
		fmt.Println(traceback)
	}
	if export {
			ExportToFile(results, name, format)
	}
}


func RenderInput(input interface{}) interface{} {

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err != nil{
		log.Fatal(err)
	}

	err = viper.Unmarshal(&input)
	if err != nil {
		log.Fatal(errors.Wrap(err, "unmarshal config file"))
	}
	return input
}