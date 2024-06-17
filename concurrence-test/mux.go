package main
import (
	"net/http"
	"sync"
)
func main() {
	go http.ListenAndServe(":8080", nil)
	var m sync.Mutex
	m.Lock()
	m.Lock()
}
