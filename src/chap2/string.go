package main

import "fmt"

func main() {
    var main_string = "Hello, world!"   /* var declares variable */
    fmt.Println(main_string[1])         /* prints 101 */
    fmt.Println(main_string[0:5])       /* prints "Hello" */
    fmt.Println("Hello, world!"[0:5])   /* also prints "Hello" */
}
