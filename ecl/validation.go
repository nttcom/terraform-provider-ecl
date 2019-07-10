package ecl

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
)

// IntInSlice returns a SchemaValidateFunc which tests if the provided value
// is of type int and matches the value of an element in the valid slice
func IntInSlice(valid []int) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (s []string, es []error) {
		v, ok := i.(int)
		if !ok {
			es = append(es, fmt.Errorf("expected type of %s to be an integer", k))
			return
		}

		for _, validInt := range valid {
			if v == validInt {
				return
			}
		}

		es = append(es, fmt.Errorf("expected %s to be one of %v, got %d", k, valid, v))
		return
	}
}

// ValidateVRID returns a SchemaValidateFunc which tests if the provided value
// is "null" or integer corresponding value in the range from 0 to 255
func ValidateVRID() schema.SchemaValidateFunc {
	return func(i interface{}, k string) (s []string, es []error) {

		v, ok := i.(string)
		if !ok {
			es = append(es, fmt.Errorf("Failed in converting value into string %s", k))
			return
		}

		if v == "null" {
			return
		}

		iv, err := strconv.Atoi(v)
		if err != nil {
			es = append(es, fmt.Errorf("Failed in converting value into int %s", err))
		}

		if iv < 0 || iv > 255 {
			es = append(es, fmt.Errorf("expected %s to be in the range from 1 to 255, got %d", k, iv))
			return
		}

		return
	}
}
