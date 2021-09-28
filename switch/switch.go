package main

import (
	"fmt"
	"runtime"
)

func main() {
	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Windows OS")
	case "darwin":
		fmt.Println("MAC OS")
	case "linux":
		fmt.Println("Linux OS")
	default:
		fmt.Println(os)
	}
}
