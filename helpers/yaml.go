package helpers

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

func CreateYaml(conf *Config) {
	yamlFile, err := yaml.Marshal(&conf)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(yamlFile))

	f, err := os.Create("values.yaml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = io.Writer.Write(f, yamlFile)
	if err != nil {
		panic(err)
	}
}
