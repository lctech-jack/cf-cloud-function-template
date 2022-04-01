package cloudfunc

import (
	"cloud.google.com/go/functions/metadata"
	"context"
	"fmt"
	"log"
	"time"
)

func EntryExampleFirestoreOnCreated(ctx context.Context, e ExampleFirestoreEvent) (err error) {
	meta, err := metadata.FromContext(ctx)
	if err != nil {
		return fmt.Errorf("metadata.FromContext: %v", err)
	}
	log.Println(meta.EventID)
	log.Println(meta.EventType)
	log.Println(meta.Resource)
	_, err = FirestoreCli.Collection("Example").Doc("Author").Create(ctx, map[string]interface{}{
		"Name":      "Jack",
		"CreatedAt": time.Now().UTC(),
	})
	return
}
