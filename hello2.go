package hello_modules

import "fmt"

func Hello2(name string) string {
	message := fmt.Sprintf("Hello %s from Go!", name)
	return message
}
