package main

import (
	"fmt"
	"reflect"
)

type orang struct{
	nama string
	umur int8
}
func (o *orang) getPropertyinfo(){
	var reflectpeople=reflect.ValueOf(o)
	if reflectpeople.Kind()== reflect.Ptr{
		reflectpeople=reflectpeople.Elem()
	}
		var reflectype=reflectpeople.Type()

		for i:=0;i<reflectpeople.NumField();i++{
			fmt.Println("nama field",reflectype.Field(i).Name)
			fmt.Println("tipe data",reflectype.Field(i).Type)
		}	
	}
func main(){
	a:=55
	reflectA:=reflect.ValueOf(a)
	fmt.Println("tipe data dari variabel tersebut",reflectA.Type())

	if reflectA.Kind() == reflect.Int{
		fmt.Println("nilai variable =",reflectA.Int())
	}

	fmt.Println("=====================================")

	var kosong interface{}
	kosong="Bryan"
	reflectionksg:=reflect.ValueOf(kosong)

	fmt.Println(reflectionksg.Interface())
	var1:=reflectionksg.Interface().(string)
	fmt.Println(var1)
	fmt.Println("=======================================================")
	orang1:=&orang{nama:"bryan",umur:26}
	orang1.getPropertyinfo();
}
