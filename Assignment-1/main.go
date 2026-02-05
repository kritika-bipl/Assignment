package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func find(value any, indent string) {
	v := reflect.ValueOf(value)
	t := reflect.TypeOf(value)

	switch v.Kind() {

	case reflect.Map:
		fmt.Println(indent, "Type:", t, "(map)")
		for _, key := range v.MapKeys() {
			fmt.Println(indent, "Key:", key.Interface())
			find(v.MapIndex(key).Interface(), indent+"  ")
		}

	case reflect.Slice, reflect.Array:
		fmt.Println(indent, "Type:", t, "(slice)")
		for i := 0; i < v.Len(); i++ {
			fmt.Printf("%sIndex %d:\n", indent, i)
			find(v.Index(i).Interface(), indent+"  ")
		}
	default:
		fmt.Printf("%sType: %v | Value: %v\n", indent, t, v.Interface())
	}
}

func main() {

	jsonStr := `{
		"name" : "Tolexo Online Pvt. Ltd",
		"age_in_years" : 8.5,
		"origin" : "Noida",
		"head_office" : "Noida, Uttar Pradesh",
		"address" : [
			{
				"street" : "91 Springboard",
				"landmark" : "Axis Bank",
				"city" : "Noida",
				"pincode" : 201301,
				"state" : "Uttar Pradesh"
			},
			{
				"street" : "91 Springboard",
				"landmark" : "Axis Bank",
				"city" : "Noida",
				"pincode" : 201301,
				"state" : "Uttar Pradesh"
			}
		],
		"sponsers" : {
			"name" : "One"
		},
		"revenue" : "19.8 million$",
		"no_of_employee" : 630,
		"str_text" : ["one","two"],
		"int_text" : [1,3,4]
	}`

	var data any

	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		panic(err)
	}

	find(data, "")
}
