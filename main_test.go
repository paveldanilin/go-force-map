package main

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type usr struct {
	Name string
	Cars []string
}

func TestForceMap_Scalar(t *testing.T) {
	str := "Test"
	num := 3310

	assert.Equal(t, str, ForceMap(str))
	assert.Equal(t, num, ForceMap(num))
}

func TestForceMap_Array(t *testing.T) {
	arr := make([]int, 0)
	arr = append(arr, 1)
	arr = append(arr, 2)
	arr = append(arr, 3)

	converted := ForceMap(arr)

	assert.True(t, reflect.ValueOf(converted).Kind() == reflect.Map)

	js, err := json.Marshal(converted)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(js))
}

func TestForceMap_Map(t *testing.T) {
	nums := []int{9,8,7,6,5,4,3,2,1}
	users := make(map[string][]int)
	users["test"] = nums

	converted := ForceMap(users)

	assert.True(t, reflect.ValueOf(converted).Kind() == reflect.Map)

	js, err := json.Marshal(converted)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(js))
}

func TestForceMap_ArrayOfStruct(t *testing.T) {
	usrs := make([]usr, 0)
	usrs = append(usrs, usr{Name: "Pasha", Cars: []string{"KIA"}})
	usrs = append(usrs, usr{Name: "Lesha"})
	usrs = append(usrs, usr{Name: "Dima", Cars: []string{"Ford", "Tesla"}})

	converted := ForceMap(usrs)

	js, err := json.Marshal(converted)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(js))
}

func TestForceMap_Struct(t *testing.T) {
	pasha := usr{Name: "Pasha", Cars: []string{"KIA"}}

	converted := ForceMap(pasha)

	js, err := json.Marshal(converted)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(js))
}
