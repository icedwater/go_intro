# Introducing Go
ISBN 978-1-491-94195-9

First edition by Caleb Doxsey in 2016.

## Prerequisites

Installing go from the [official website][gosite] was recommended, and the book
used go1.5 windows/amd64. However, that was 7 years ago. I work on Ubuntu 22.04
so the version is go 1.18~ubuntu2.

The command I used to set things up on Ubuntu was:

    apt install golang-go

An alternative would be to use `snap` or some other sort of package manager.

## Notes

- While `$GOPATH` is introduced early on in the book, the official repo package
from Ubuntu sets up `$GOROOT`. This means there was no extra step needed to get
`go run main.go` to work as on page 4.

- The syntax for `godoc` has also changed:

      go doc fmt Println

  On Ubuntu, `godoc` is now a utility program, not built-in documentation as in
  the book.

[gosite]: https://golang.org/dl
