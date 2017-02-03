/*
Package greeting provides simple functionality that will create a greeting.
It will use a default value if input is not provided.
*/
package greeting

import "fmt"

const testVersion = 3

// HelloWorld returns a greeting with a default value if no input is provided
func HelloWorld(s string) string {
	// Write some code here to pass the test suite.
	if s == "" {
		s = "World"
	}
	return fmt.Sprintf("Hello, %s!", s)
}
