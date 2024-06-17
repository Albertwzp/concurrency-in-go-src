package main
import (
	"fmt"
	"net/http"
)
func main() {
	go http.ListenAndServe(":8080", nil)
	fmt.Scanln()
}
