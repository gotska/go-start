package main

import (
	"fmt"
	"time"
	/*"os"
	"strconv"*///"strings"
	//"bytes"
)

func main() {

	startT := time.Now()
	s := ""
	s1 := "1"
	//var buffer = bytes.Buffer{}
	for i := 1; i < 100000; i++ {
		//buffer.WriteString(s1)

		s = s + s1

	}
	secT := time.Since(startT).Seconds()

	fmt.Println(secT)

	//fmt.Println(s)
}
