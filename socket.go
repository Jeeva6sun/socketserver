package socketserver

import "fmt"

type userDetail struct {
	id    int
	name  string
	email string
}

func Hello(name string) string {
	return fmt.Sprintf("Hello mr , %s!", name)
}

func username(name string) string {
	return fmt.Sprintf(name)
}

func useremail(name string) string {
	return fmt.Sprintf(name + "@email.com")
}

func welcomemessage() string {
	return fmt.Sprintf("welcome user")
}
