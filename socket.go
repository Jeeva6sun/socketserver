package socketserver

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("Hello mr , %s!", name)
}

func username(name string) string {
	return fmt.Sprintf(name)
}
