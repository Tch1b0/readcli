package utility

import "fmt"

func RequestValueInput(requestedValue string, defaultValue interface{}) interface{} {
	var request string
	if defaultValue != nil {
		request = fmt.Sprintf("%s (%s): ", requestedValue, defaultValue)
	} else {
		request = fmt.Sprintf("%s: ", requestedValue)
	}

	answer := GetInput(request)
	if len(request) == 0 {
		return defaultValue
	} else {
		return answer
	}
}

func GetInput(question string) string {
	fmt.Print(question)
	var input string
	fmt.Scanln(&input)
	return input
}
