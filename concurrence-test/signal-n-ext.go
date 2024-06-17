package main
import (
	"context"
	"net/http"
	"os"
	"os/signal"
)
func main() {
	go http.ListenAndServe(":8080", nil)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	<-ctx.Done()
}
