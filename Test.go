package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	timeConversion("07:05:45PM")
}

func timeConversion(twelveHourFormat string) {
	stringArr := strings.Split(twelveHourFormat, ":")
	hourInt, _ := strconv.Atoi(stringArr[0])
	minuteStr, _ := strconv.Atoi(stringArr[1])
	// minuteStr := strconv.ParseInt(stringArr[1])
	secondsString := stringArr[2]
	var numberOfSeconds string
	var amOrPm string
	if strings.Contains(secondsString, "PM") { //PM
		numberOfSeconds = strings.Split(secondsString, "PM")[0]
		amOrPm = "PM"
	} else {
		numberOfSeconds = strings.Split(secondsString, "AM")[0]
		amOrPm = "AM"
	}
	fmt.Printf("%v for first\n", hourInt)
	fmt.Printf("%v for second\n", minuteStr)
	fmt.Printf("%v for %v third\n", string(amOrPm), numberOfSeconds)
}
