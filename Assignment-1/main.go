package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func find(value any, indent string) {
	v := reflect.ValueOf(value)
	t := reflect.TypeOf(value)
	if !v.IsValid() {
		fmt.Println("null value")
		return
	}
	

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
		
    "moq": 0.01,
    "pack_quantity": 1,
    "item_quantity": 3,
    "pack_measurement_unit": {
        "id": 2,
        "name": "Bale"
    },
    "measurement_unit": {
        "id": 1,
        "name": "Bag"
    },
    "hsn": {
        "code": "0400",
        "id": 3343
    },
    "tax_rate": 18,
    "lp": 100,
    "mrp": 1000,
    "companies": "9",
    "reorder_value": "",
    "hsn_id": 3343,
    "tax_group": {
        "id": 3293,
        "name": "18%"
    },
    "sale": false,
    "purchase": false,
    "intermediate": false,
    "is_bundle": false,
    "users": "5"
}`

	var data any

	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		panic(err)
	}

	find(data, "")
}
