package main

import (
	"fmt"
	"reflect"
)

type GoErr struct {
	code    int
	message string
}

const (
	valueOf = "reflect.ValueOf()."
	typeOf  = "reflect.TypeOf()."
)

func main() {

	err := &GoErr{
		code:    200,
		message: "Success",
	}
	t := reflect.TypeOf(err)
	v := reflect.ValueOf(err)
	fmt.Println(typeOf + "String(): " + t.String())
	fmt.Println("reflect.TypeOf().Name() does not resolve pointers, so it is empty")
	fmt.Println(typeOf + "Name(): " + t.Name())
	fmt.Println(typeOf + "Elem().Name(): " + t.Elem().Name())
	fmt.Println(valueOf + "String(): " + v.String())
	fmt.Println(valueOf + "Kind().String(): " + v.Kind().String())
	if v.Kind() == reflect.Ptr {
		fmt.Println("Kind() is reflect.Ptr so we need to use indirect to get reference to actual value before we go any further")
		v = reflect.Indirect(v)
	}
	fmt.Println(valueOf + "Kind().String(): " + v.Kind().String())
	fmt.Println(typeOf+"Field(0).Int(): ", v.Field(0).Int())
	fmt.Println(typeOf+"Field(1).String(): ", v.Field(1).String())

	//fmt.Println("reflect.ValueOf().NumField(): ", v.NumField())
}
