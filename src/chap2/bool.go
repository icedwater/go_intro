package main

import "fmt"

func main() {
    var output = "The result is:"
    /* string and int constants are not booleans. whew. */
    fmt.Println(output, true && false)  /* expect false */
    fmt.Println(output, false || true)  /* expect true */
    fmt.Println(output, !false == true) /* expect true */
}
