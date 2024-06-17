package main
import (
	"net/http"
	"os"
	"os/signal"
)
func main() {
	go http.ListenAndServe(":8080", nil)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
