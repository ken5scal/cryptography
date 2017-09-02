package main

import (
	"fmt"
	"reflect"
	"bytes"
)

func main() {
	ff := "hogehoge"
	fmt.Println(reflect.TypeOf(ff))
	fmt.Printf("%b", ff)

	var buffer bytes.Buffer
	for i := 0; i < len(ff);  i++ {
		fmt.Fprintf(&buffer, "%b", ff[i])
	}

	fmt.Println([]byte(ff))

}
