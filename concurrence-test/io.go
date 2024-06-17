package main
import (
	"net/http"
	"os"
)
func main() {
	go http.ListenAndServe(":8080", nil)
	os.Stdin.Read(make([]byte, 1))
}
