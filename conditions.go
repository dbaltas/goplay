package main

import "fmt"

func modula(n int) {
	for i :=0; i <10; i++ {
		if i % n == 0 {
			fmt.Println(i)
		}
	}

}

func main() {
	modula(3)
}