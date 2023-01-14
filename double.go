// Package path implements utility routines for manipulating slash-separated
// paths.
//
// The path package should only be used for paths separated by forward
// slashes, such as the paths in URLs. This package does not deal with
// Windows paths with drive letters or backslashes; to manipulate
// operating system paths, use the [path/filepath] package.
package double

// # This is heading
//
// Nested list:
//
//  1. On Read error or close, the stop func is called.
//
//     - Subitem 1.
//
//     - Subitem 2.
//
//  2. On Read failure, if reqDidTimeout is true, the error is wrapped and
//     marked as net.Error that hit its timeout.
//
// Link “[JSON and Go].”
//
// Golang cmd:
//
//	go run generate_cert.go --rsa-bits 1024 --host 127.0.0.1,::1,example.com \
//	    --ca --start-date "Jan 1 00:00:00 1970" --duration=1000000h
//
// JSON response:
//
//	{
//	    "kind":"MyAPIObject",
//	    "apiVersion":"v1",
//	    "myPlugin": {
//	        "kind":"PluginA",
//	        "aOption":"foo",
//	    },
//	}
//
// Pattern:
//
//	pattern:
//	    { term }
//	term:
//	    '*'         matches any sequence of non-/ characters
//	    '?'         matches any single non-/ character
//	    '[' [ '^' ] { character-range } ']'
//	                character class (must be non-empty)
//	    c           matches character c (c != '*', '?', '\\', '[')
//	    '\\' c      matches character c
//
// Code examples:
//
//	func GuessingGame() {
//	    var s string
//	    fmt.Printf("Pick an integer from 0 to 100.\n")
//	    answer := sort.Search(100, func(i int) bool {
//	        fmt.Printf("Is your number <= %d? ", i)
//	        fmt.Scanf("%s", &s)
//	        return s != "" && s[0] == 'y'
//	    })
//	    fmt.Printf("Your number is %d.\n", answer)
//	}
//
// Link to package [io.EOF].
//
// List examples:
//
//   - Item 1.
//   - Item 2.
//   - Item 3.
//
// Code examples 2:
//
//	Sqrt(+Inf) = +Inf
//	Sqrt(±0) = ±0
//	Sqrt(x < 0) = NaN
//	Sqrt(NaN) = NaN
//
// [JSON and Go]: https://golang.org/doc/articles/json_and_go.html
func Double(n int) int {
	return n * 2
}
