package main

func main() {
	waitForever := make(chan interface{})
	go func() {
		waitForever<-"abc"
		//panic("Test Panic")
	}()
	<-waitForever
}