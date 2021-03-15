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
		StructField struct {
			subField1 string
			subField2 int
		}
	}{
		FieldString: "stroka",
		FieldInt:    107,
		StructField: struct {
			subField1 string
			subField2 int
		}{
			subField1: "subField1value",
			subField2: 123,
		},
	}
	PrintStruct(v)
	newValues := map[string]interface{}{
		"FieldString": "stroka_updated",
		"FieldInt":    777,
		"unknown":     "unknoun",
		"StructField": map[string]interface{}{
			"subField1": "newSubFieldValue",
			"subField2": 888,
		},
	}
	UpdateStruct(&v, newValues)
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

func UpdateStruct(in interface{}, values map[string]interface{}) {
	if in == nil {
		return
	}

	inVal := reflect.ValueOf(in)
	if inVal.Kind() == reflect.Ptr {
		inVal = inVal.Elem()
	} else {
		return
	}

	if inVal.Kind() != reflect.Struct {
		return
	}

	for fieldName, newValue := range values {
		newValueReflect := reflect.ValueOf(newValue)
		fieldValue := inVal.FieldByName(fieldName)
		fmt.Println(fieldValue)
		newValueReflectKind := newValueReflect.Kind()
		if newValueReflectKind != reflect.Map {
			fieldValue := inVal.FieldByName(fieldName)
			if fieldValue.CanSet() {
				fieldValue.Set(newValueReflect)
			}

		}
	}
}
