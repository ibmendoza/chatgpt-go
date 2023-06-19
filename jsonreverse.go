//https://go.dev/play/p/sNnS1D66Gte
//https://medium.com/@pithomlabs/test-driven-development-with-prompt-engineering-2ba9efb9af7

package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func reverseObject(jsonObj json.RawMessage) json.RawMessage {
	var data map[string]json.RawMessage
	err := json.Unmarshal(jsonObj, &data)
	if err != nil {
		panic(err)
	}

	reversedData := make(map[string]json.RawMessage)

	for key, value := range data {
		var reversedValue json.RawMessage

		if isObject(value) {
			reversedValue = reverseObject(value)
		} else {

			if isArray(value) {
				//fmt.Println("array,", string(value))

				output, err := processArray(string(value))
				if err != nil {
					fmt.Println("Error:", err)
					panic(err)
				}

				//fmt.Println(output)
				reversedValue = output //reverseArray(value)
			} else {
				reversedValue = value
			}
		}

		reversedData[key] = reversedValue
	}

	reversedJSON, err := json.Marshal(reversedData)
	if err != nil {
		//fmt.Println("reverseObject", err)
		panic(err)
	}

	return reversedJSON
}

func reverseArray(arr []interface{}) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func isArray(value json.RawMessage) bool {
	var arr []json.RawMessage
	err := json.Unmarshal(value, &arr)
	//fmt.Println("isArray, ", err)
	return err == nil
}

func isObject(value json.RawMessage) bool {
	var obj map[string]json.RawMessage
	err := json.Unmarshal(value, &obj)
	//fmt.Println("isObject, ", err)
	return err == nil
}

//***

func processValue(value interface{}) {
	fmt.Println("processValue,", value)
	switch v := value.(type) {
	case []interface{}:
		reverseArray(v)
		for _, item := range v {
			if reflect.TypeOf(item).Kind() == reflect.Map {
				processValue(item)
			}
		}

	/*  //replace with reverseObject
	case map[string]interface{}:
	 reverseObject(v)
	 for _, item := range v {
	  if reflect.TypeOf(item).Kind() == reflect.Map || reflect.TypeOf(item).Kind() == reflect.Slice {
	   processValue(item)
	  }
	 }
	}
	*/

	//case map[string]interface{}:
	case map[string]json.RawMessage:
		//case json.RawMessage:
		//reverseObject(v)
		for _, item := range v {
			if reflect.TypeOf(item).Kind() == reflect.Map || reflect.TypeOf(item).Kind() == reflect.Slice {
				processValue(item)
			}
		}
	}
}

func processArray(jsonStr string) (json.RawMessage, error) {
	var data interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return nil, err
	}

	processValue(data)

	result, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return result, nil
	//return string(result), nil
}

func main() {
	// JSON input
	jsonStr := `{
  "i": {
   "k": {
    "m": "monkey",
    "l": "lion"
   },
   "j": "jupiter"
  },
  "f": {
   "h": "horse",
   "g": "grape",
   "a": [1,2,3, {"four": 4, "five": 5}]
  },
  "array": ["a", "b", 20, false, "c", {"six": 6, "seven": 7}]
 }`

	// Parse the JSON input
	var jsonObj json.RawMessage
	err := json.Unmarshal([]byte(jsonStr), &jsonObj)
	if err != nil {
		panic(err)
	}

	// Reverse the order of key-value pairs in the JSON object
	reversedJSON := reverseObject(jsonObj)

	// Print the reversed JSON string
	fmt.Println(string(reversedJSON))
}
