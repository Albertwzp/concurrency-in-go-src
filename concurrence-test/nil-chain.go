package main
import "net/http"
func main() {
	go http.ListenAndServe(":8080", nil)
	var c chan struct{}
	// <-c
	c <- struct{}{}
}
