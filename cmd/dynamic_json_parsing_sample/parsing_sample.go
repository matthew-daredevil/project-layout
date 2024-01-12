package dynamic_json_parsing_sample

import (
	"encoding/json"
	"fmt"
)

func Dynamic_json_parsing_sample() {
	// JSON data as a byte slice
	jsonData := []byte(`{
  "name": "John Doe",
  "age": 30,
  "city": "New York",
  "hasCar": true,
  "languages": ["Go", "JavaScript"]
 }`)

	// Parse JSON into an empty interface
	var result interface{}
	err := json.Unmarshal(jsonData, &result)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Accessing dynamic JSON fields
	dataMap, ok := result.(map[string]interface{})
	if !ok {
		fmt.Println("Invalid JSON structure")
		return
	}

	// Accessing specific fields
	name, nameExists := dataMap["name"].(string)
	age, ageExists := dataMap["age"].(float64)
	city, cityExists := dataMap["city"].(string)
	hasCar, hasCarExists := dataMap["hasCar"].(bool)
	languages, languagesExists := dataMap["languages"].([]interface{})

	// Displaying parsed data
	if nameExists {
		fmt.Println("Name:", name)
	}

	if ageExists {
		fmt.Println("Age:", int(age))
	}

	if cityExists {
		fmt.Println("City:", city)
	}

	if hasCarExists {
		fmt.Println("Has Car:", hasCar)
	}

	if languagesExists {
		fmt.Println("Languages:")
		for _, lang := range languages {
			fmt.Println(" -", lang)
		}
	}
}
