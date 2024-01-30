package main //package called main which creates exe files

import "fmt" //imports functionalities of 'fmt' library into the main package, its like a sub-package

func main() {
	fmt.Println("Hi there!")
}

//1. How do we run code in our project?
//2. What is 'package'?
//3. What is 'import'?
//4. How is the HelloWorld.go file organised?
//5. executable packages vs. reusable packages

//go build creates executable files by compiling the code but 'go run' compiles and then executes the code file.
