package main

import "fmt"

func nums() {
	i := 1
	j := 2
	fmt.Println(i+j)	
}

func strings(k string) {
	var prefix = "do "
	fmt.Println(prefix + k)
}

func main() {
	nums()
	strings("boo")
}
