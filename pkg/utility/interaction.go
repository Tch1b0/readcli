package utility

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

func RequestValueInput(requestedValue string, defaultValue interface{}) string {
	requestedValue = color.YellowString(requestedValue)
	var request string
	if defaultValue != nil {
		coloredDefaultValue := color.GreenString("%v", defaultValue)
		request = fmt.Sprintf("%s (%v): ", requestedValue, coloredDefaultValue)
	} else {
		request = fmt.Sprintf("%s: ", requestedValue)
	}

	answer := GetInput(request)
	if len(answer) == 0 {
		return fmt.Sprintf("%v", defaultValue)
	} else {
		return answer
	}
}

func RequestValueListInput(requestedValue string) []string {
	requestedValue = color.YellowString(requestedValue)
	answer := GetInput(fmt.Sprintf("%s [seperate with comma]: ", requestedValue))
	if answer == "" {
		return []string{}
	}
	s, _ := regexp.Compile(", ?")

	return s.Split(answer, -1)
}

func RequestDecisionInput(question string, preferTrue bool) bool {
	question = color.YellowString("%s?", question)

	var yes string
	var no string
	if preferTrue {
		yes = "Y"
		no = "n"
	} else {
		yes = "y"
		no = "N"
	}

	options := color.BlueString("(%s/%s)", yes, no)

	answer := GetInput(fmt.Sprintf("%s %s: ", question, options))
	answer = strings.ToLower(answer)
	if strings.HasPrefix(answer, "y") {
		return true
	} else if strings.HasPrefix(answer, "n") {
		return false
	} else {
		return preferTrue
	}
}

func GetInput(question string) string {
	fmt.Print(question)
	var input string
	reader := bufio.NewScanner(os.Stdin)
	if reader.Scan() {
		input = reader.Text()
	}
	return input
}
