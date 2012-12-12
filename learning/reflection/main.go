package main

import (
	"fmt"
	"reflect"
)

// http://golang.org/doc/articles/laws_of_reflection.html
func main() {
	var x float64 = 3.4

	fmt.Println("1. Reflection goes from interface value to reflection object.")
	{

		fmt.Println("type:", reflect.TypeOf(x))
		fmt.Println("value:", reflect.ValueOf(x))

		v := reflect.ValueOf(x)
		fmt.Println("type:", v.Type())
		fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
		fmt.Println("value:", v.Float())
	}
	{
		var x uint8 = 'x'
		v := reflect.ValueOf(x)
		fmt.Println("type:", v.Type())                            // uint8.
		fmt.Println("kind is uint8: ", v.Kind() == reflect.Uint8) // true.
		x = uint8(v.Uint())                                       // v.Uint returns a uint64.
	}
	{
		type MyInt int
		var x MyInt = 7
		v := reflect.ValueOf(x)
		fmt.Println("type:", v.Type())
	}

	fmt.Println("2. Reflection goes from reflection object to interface value.")
	{
		v := reflect.ValueOf(x)
		y := v.Interface().(float64) // y will have type float64.
		fmt.Println(y)
		fmt.Println(v.Interface())
		fmt.Printf("value is %7.1e\n", v.Interface())
	}

	fmt.Println("3. To modify a reflection object, the value must be settable.")
	{
		v := reflect.ValueOf(x)
		//v.SetFloat(7.1) // Error: will panic.

		fmt.Println("settability of v:", v.CanSet())
	}
	{
		p := reflect.ValueOf(&x) // Note: take the address of x.
		fmt.Println("type of p:", p.Type())
		fmt.Println("settability of p:", p.CanSet())

		v := p.Elem()
		fmt.Println("settability of v:", v.CanSet()) // yes!!!!

		v.SetFloat(7.1)
		fmt.Println(v.Interface())
		fmt.Println(x)
	}
	{
		type T struct {
			A int
			B string
		}
		t := T{23, "skidoo"}
		s := reflect.ValueOf(&t).Elem()
		typeOfT := s.Type()
		for i := 0; i < s.NumField(); i++ {
			f := s.Field(i)
			fmt.Printf("%d: %s %s = %v\n", i,
				typeOfT.Field(i).Name, f.Type(), f.Interface())
		}

		s.Field(0).SetInt(77)
		s.Field(1).SetString("Sunset Strip")
		fmt.Println("t is now", t)
	}
}

