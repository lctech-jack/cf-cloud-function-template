package cloudfunc

import (
	"context"
	"log"
)

func init() {
	// Write custom initialization here ...
}

func EntryHelloWorld(ctx context.Context, e FirestoreEvent) (err error) {
	log.Println("Hello World")
	return
}
