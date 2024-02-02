package main

import(
	"fmt"
	"time"
	"reflect"
)

func main(){
	fmt.Println("Test")
	a := time.Now().Weekday()
	// fmt.Print(a)
	fmt.Print(reflect.TypeOf(a))
}