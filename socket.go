package socketserver

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("Hello mr , %s!", name)
}

func Username(name string) string {
	return fmt.Sprintf(name)
}

func Useremail(name string) string {
	return fmt.Sprintf("%s", name+"@email.com")
}

func Welcomemessage(name string) string {
	return fmt.Sprintf("welcome user %s", name)
}
