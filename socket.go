package socketserver

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("Hello mr , %s!", name)
}

func username(name string) string {
	return fmt.Sprintf(name)
}

func useremail(name string) string {
	return fmt.Sprintf("%s", name+"@email.com")
}

func welcomemessage(name string) string {
	return fmt.Sprintf("welcome user %s", name)
}
