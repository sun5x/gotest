package gotest

import "fmt" 

// say Hi to someone
func SayHi(name string) string {
   fmt.Println("hello golang!!!")
   return fmt.Sprintf("Hi, %s", name)
}
