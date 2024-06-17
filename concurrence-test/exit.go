package main
import (
	"net/http"
	"runtime"
)
func main() {
	go http.ListenAndServe(":8080", nil)
	runtime.Goexit()
}
