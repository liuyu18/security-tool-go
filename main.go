package main

import (
	"fmt"
	"security-go/c3"
	"security-go/c4"
)

func main() {
	fmt.Println("run main func")
	c4.FindRecentlyChangedFile()
	c3.CheckTheFile()
}
func test() {

}
