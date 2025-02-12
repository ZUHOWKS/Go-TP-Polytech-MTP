package main

import "fmt"

func main() {
	vec := vec2D32{1, 2}
	fmt.Println(vec.sub(vec2D32{3, 4}))
}
