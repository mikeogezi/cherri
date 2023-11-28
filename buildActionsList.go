package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Assuming the actionDefinition struct, parameterDefinition struct,
// tokenType type, and standardActions function are defined elsewhere
func populateActions() {
	actions = make(map[string]*actionDefinition)
	mediaActions()
	documentActions()
}

func main() {
	populateActions()

	var actionsList []map[string]interface{}
	for key, action := range actions {
		properties := make(map[string]map[string]interface{})
		var required []string
		for _, param := range action.parameters {
			paramMap := map[string]interface{}{
				"type": param.validType, // Use the actual type of validType, converting it to a string if necessary
			}
			if len(param.enum) > 0 {
				paramMap["enum"] = param.enum
			}
			if param.optional {
				paramMap["optional"] = param.optional
			}
			if param.infinite {
				paramMap["infinite"] = param.infinite
			}
			if param.defaultValue != nil {
				paramMap["defaultValue"] = param.defaultValue
			}
			if !param.optional {
				required = append(required, param.name)
			}
			properties[param.name] = paramMap
		}

		parametersObject := map[string]interface{}{
			"type":       "object",
			"properties": properties,
		}
		if len(required) > 0 {
			parametersObject["required"] = required
		}

		functionObject := map[string]interface{}{
			"name":       key,
			"parameters": parametersObject,
		}

		actionMap := map[string]interface{}{
			"type":     "function",
			"function": functionObject,
		}
		actionsList = append(actionsList, actionMap)
	}

	jsonData, err := json.MarshalIndent(actionsList, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error serializing actions to JSON: %v\n", err)
		return
	}

	outputFilePath := "../actionsList.json"
	if len(os.Args) > 1 {
		outputFilePath = os.Args[1]
	}

	if err := ioutil.WriteFile(outputFilePath, jsonData, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing JSON to file: %v\n", err)
		return
	}

	fmt.Printf("Actions JSON written to %s\n", outputFilePath)
}
