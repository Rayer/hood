package hood

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func PrintConfidentialData(binding interface{}) (string, error) {
	ift := reflect.TypeOf(binding)
	ifv := reflect.ValueOf(binding)
	var kvString []string

	for i := 0; i < ifv.NumField(); i++ {
		field := ift.Field(i)
		if tagContent, exists := field.Tag.Lookup("confidential"); exists {
			valueField := ifv.Field(i)
			if valueField.Kind() != reflect.String {
				return "", fmt.Errorf("confidental tag can be only applied on string")
			}
			values := strings.Split(tagContent, ",")

			keepFirst := 0
			var err error
			if len(values) > 0 && strings.Trim(values[0], " ") != "" {
				keepFirst, err = strconv.Atoi(strings.Trim(values[0], " "))
				if err != nil {
					return "", fmt.Errorf("confidental value can only be integer")
				}
			}
			keepTail := 0
			if len(values) > 1 {
				keepTail, err = strconv.Atoi(strings.Trim(values[1], " "))
				if err != nil {
					return "", fmt.Errorf("confidental value can only be integer")
				}
			}

			value := valueField.String()
			ret := HoodString(value, keepFirst, keepTail)

			kvString = append(kvString, fmt.Sprintf("%v:%v", field.Name, ret))

		} else {
			t := field.Type
			if t.Kind() == reflect.Struct {
				inner, err := PrintConfidentialData(ifv.Field(i).Interface())
				if err != nil {
					return "", err
				}
				kvString = append(kvString, inner)
			} else {
				kvString = append(kvString, fmt.Sprintf("%v:%v", field.Name, ifv.Field(i).Interface()))
			}
		}
	}

	return "{" + strings.Join(kvString, " ") + "}", nil
}

func HoodString(target string, keepFirst, keepTail int) string {
	ret := ""
	if len(target) < keepFirst && len(target) < keepTail {
		ret = target
	} else {
		for i := 0; i < len(target); i++ {
			if i < keepFirst || i >= len(target)-keepTail {
				ret += string(target[i])
			} else {
				ret += "*"
			}
		}
	}
	return ret
}
