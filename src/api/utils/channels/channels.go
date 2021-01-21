package channels

import "fmt"

func OK(done <-chan bool) bool {

	select {
	case ok := <-done:
		if ok {
			fmt.Println("ok true")
			return true
		}
	}
	return false
}
