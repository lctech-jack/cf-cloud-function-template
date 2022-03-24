package cloudfunc

import (
	"context"
	"log"
)

func init() {
	// Write custom initialization here ...
}

func EntryByeWorld(ctx context.Context, e FirestoreEvent) (err error) {
	log.Println("Bye World")
	return
}
