package tools

import (
	"flag"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// YamlConfig struct
type YamlConfig struct {
	API struct {
		Port int `yaml:"port"`
	}
}

func getConfig() YamlConfig {
	fmt.Println("Parsing Config YAML file")

	var fileName string

	flag.StringVar(&fileName, "f", "", "YAML file to parse")
	flag.Parse()

	if fileName == "" {
		fmt.Println("Please provide yaml file by using -f option")
	}

	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error parsing YAML file: %s\n", err)
		panic('a')
	}

	var yamlConfig YamlConfig
	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %S\n", err)
	}

	fmt.Printf("Result: %v\n", yamlConfig)

	return yamlConfig
}

// Config static value
var Config YamlConfig = getConfig()
