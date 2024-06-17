package main
import (
	"net/http"
	"sync"
)
func main() {
	go http.ListenAndServe(":8080", nil)
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
