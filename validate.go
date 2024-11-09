package validate

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Errors map[string]string

func Validate(item any) (bool, Errors) {
	v := reflect.ValueOf(item)

	validationErrors := make(Errors)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag, ok := v.Type().Field(i).Tag.Lookup("validate")
		if !ok {
			continue
		}

		rules := strings.Split(tag, ",")
		for _, rule := range rules {
			ruleKey := strings.Split(rule, "=")[0]

			switch strings.TrimSpace(ruleKey) {
			case "required":
				err := handleRequired(field)
				if err != nil {
					validationErrors[v.Type().Field(i).Name] = err.Error()
				}
			case "min":
				err := handleMin(field, rule)
				if err != nil {
					validationErrors[v.Type().Field(i).Name] = err.Error()

				}
			case "max":
				err := handleMax(field, rule)
				if err != nil {
					validationErrors[v.Type().Field(i).Name] = err.Error()
				}
			case "email":
				// handle email
			}
		}
	}

	ok := len(validationErrors) == 0
	return ok, validationErrors
}

func handleRequired(field reflect.Value) error {
	if field.Type().Kind() != reflect.String {
		return fmt.Errorf("%s is not of type string", field.Type().Name())
	}

	if field.String() == "" {
		return fmt.Errorf("%s is required and cannot be empty", field.Type().Name())
	}

	return nil
}

func handleMin(field reflect.Value, rule string) error {
	k := field.Type().Kind()
	if k != reflect.String && k != reflect.Int {
		return fmt.Errorf("must be string or int")
	}

	minAsString := strings.Split(rule, "=")[1]
	min, err := strconv.ParseInt(minAsString, 10, 64)
	if err != nil {
		return fmt.Errorf("min is not a number")
	}

	switch k {
	case reflect.String:
		if int64(len(field.String())) < min {
			return fmt.Errorf("must be minimum length of %d", min)
		}
	case reflect.Int:
		if field.Int() < min {
			return fmt.Errorf("must be minimum length of %d", min)

		}
	}

	return nil
}

func handleMax(field reflect.Value, rule string) error {
	k := field.Type().Kind()
	if k != reflect.String && k != reflect.Int {
		return fmt.Errorf("must be string or int")
	}

	maxAsString := strings.Split(rule, "=")[1]
	max, err := strconv.ParseInt(maxAsString, 10, 64)
	if err != nil {
		return fmt.Errorf("min is not a number")
	}

	switch k {
	case reflect.String:
		if int64(len(field.String())) > max {
			return fmt.Errorf("must be maximum length of %d", max)
		}
	case reflect.Int:
		if field.Int() > max {
			return fmt.Errorf("must be maximum length of %d", max)

		}
	}

	return nil
}
