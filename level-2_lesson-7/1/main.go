package main

import (
	"fmt"
	"log"
	"reflect"
)

func main() {
	v := struct {
		FieldString string `json:"field_string"`
		FieldInt    int
	}{
		FieldString: "stroka",
		FieldInt:    107,
	}
	PrintStruct(v)
	newValues := map[string]interface{}{
		"FieldString": "stroka_updated",
		"FieldInt":    777,
		"unknown":     "unknoun",
	}
	UpdateStruct(&v, newValues)
	fmt.Println("")
	PrintStruct(v)
}

func PrintStruct(in interface{}) {
	if in == nil {
		return
	}

	val := reflect.ValueOf(in)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)

		if typeField.Type.Kind() == reflect.Struct {
			log.Printf("nested field: %v", typeField.Name)
			PrintStruct(val.Field(i).Interface())
			continue
		}
		log.Printf("\tname=%s, type=%s, value=%v, tag=`%s`\n",
			typeField.Name,
			typeField.Type,
			val.Field(i),
			typeField.Tag,
		)
	}
}

func UpdateStruct(in interface{}, values map[string]interface{}) error {
	if in == nil {
		return fmt.Errorf("empty data, nothing to update")
	}

	inVal := reflect.ValueOf(in)
	if inVal.Kind() == reflect.Ptr {
		inVal = inVal.Elem()
	} else {
		return fmt.Errorf("in data should be a pointer to a struct")
	}

	if inVal.Kind() != reflect.Struct {
		return fmt.Errorf("in data should be a pointer to a struct")
	}

	for fieldName, newValue := range values {
		newValueReflect := reflect.ValueOf(newValue)
		newValueReflectKind := newValueReflect.Kind()
		if newValueReflectKind != reflect.Map {
			fieldValue := inVal.FieldByName(fieldName)
			if fieldValue.CanSet() {
				fieldValue.Set(newValueReflect)
			}

		}
	}
	return nil
}
