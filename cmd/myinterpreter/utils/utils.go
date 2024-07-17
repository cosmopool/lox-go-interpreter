package utils

import (
	"fmt"
	"strings"
)

func VariableToString(variable interface{}, toNull bool) string {
	if variable == nil && toNull {
		return "null"
	}

  if variable == nil {
    return "nil"
  }

	_, isFloat := variable.(float64)
	if !isFloat {
		return fmt.Sprint(variable)
	}

	separated := strings.Split(fmt.Sprint(variable), ".")
	if len(separated) == 1 {
		return fmt.Sprintf("%.1f", variable)
	}

	decimalPart := separated[len(separated)-1]
	decimalPart = strings.ReplaceAll(decimalPart, "0", "")

	if decimalPart == "" {
		return fmt.Sprintf("%.1f", variable)
	}

	return fmt.Sprintf("%g", variable)
}
