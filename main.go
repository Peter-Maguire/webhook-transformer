package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"webhook-transformer/config"
	"webhook-transformer/io"
)

func main() {

	inputHandlers := map[string]io.InputHandler{
		"http": &io.InputHTTP{},
	}

	outputHandlers := map[string]io.OutputHandler{
		"http": &io.OutputHTTP{},
	}

	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	yamlConfig := config.Config{}
	err = yaml.Unmarshal(yamlFile, &yamlConfig)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, wh := range yamlConfig.Webhooks {
		outputFuncs := make([]io.OutputFunc, len(wh.Outputs))
		for i, output := range wh.Outputs {
			fmt.Println("Setting up output ", output.Type)
			oh := outputHandlers[output.Type]
			oh.Initialise()
			outputFuncs[i] = oh.SetupOutput(output)
		}

		for _, input := range wh.Inputs {
			fmt.Println("Setting up input ", input.Type)
			ih := inputHandlers[input.Type]
			ih.Initialise()
			ih.SetupInput(input, outputFuncs)
		}

	}

	forever := make(chan bool)
	<-forever

}
