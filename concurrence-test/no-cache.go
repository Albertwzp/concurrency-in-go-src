package main
import "net/http"
func main() {
	go http.ListenAndServe(":8080", nil)
	c := make(chan struct{})
	<-c
	//c <- struct{}{}
}
