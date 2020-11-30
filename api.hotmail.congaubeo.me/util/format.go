package util

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

// FormatInterfaceToInt ...
func FormatInterfaceToInt(value interface{}) int {
	result, _ := strconv.Atoi(value.(string))
	return result
}
func BytesToString(data []byte) string {
	return string(data[:])
}

func XMLEncoder(data interface{}) []byte {
	output, err := xml.MarshalIndent(data, "", "")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	return output
}
