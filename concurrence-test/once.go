package main
import (
	"fmt"
	"sync"
)
func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)

       /*randvalue := func() int {
       	return rand.Int()
       }
       bar := sync.OnceValue(randvalue)
       for i := 0; i < 10; i++ {
       	fmt.Println(bar())
       }*/

	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
               fmt.Println(i)
		<-done
	}
}
