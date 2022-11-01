package test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestName(t *testing.T) {
	var s []string

	ptr := &s
	fmt.Println(ptr == nil)
	value := reflect.ValueOf(ptr)
	fmt.Println(value.Kind())
	if value.Kind() != reflect.Ptr {
		fmt.Println("must pass a pointer, not a value, to StructScan destination")
	}
	if value.IsNil() {
		fmt.Println("nil pointer passed to StructScan destination")
	}
}
