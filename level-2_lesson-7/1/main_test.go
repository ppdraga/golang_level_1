package main

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {

	v := struct {
		FieldString string `json:"field_string"`
		FieldInt    int
	}{
		FieldString: "stroka",
		FieldInt:    107,
	}
	newValues := map[string]interface{}{
		"FieldString": "stroka_updated",
		"FieldInt":    777,
		"unknown":     "unknoun",
	}
	UpdateStruct(&v, newValues)
	if !(v.FieldString == newValues["FieldString"]) {
		t.Error("v.FieldString expected", newValues["FieldString"], "Got", v.FieldString)
	}
	if !(v.FieldInt == newValues["FieldInt"]) {
		t.Error("v.FieldString expected", newValues["FieldInt"], "Got", v.FieldInt)
	}
}
