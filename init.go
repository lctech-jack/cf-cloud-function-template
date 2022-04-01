package cloudfunc

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	firebase "firebase.google.com/go/v4"
	"log"
	"os"
)

// custom logger
var logger *log.Logger

// read shared env (optional)
var (
	// Default run mode is test
	RUN_MODE     = "test"
	PROJECT_ID   string
	FirebaseApp  *firebase.App
	FirestoreCli *firestore.Client
)

// write shared initialization here
func init() {
	var err error
	logger = log.New(os.Stdout, "", log.LstdFlags)
	RUN_MODE = os.Getenv("runmode")
	log.Println(RUN_MODE)
	switch RUN_MODE {
	case "test":
	default:
		log.Println("RUN_MODE = test")
		PROJECT_ID = "testing"
		// export env variable
		//cmd := exec.Command("export FIRESTORE_EMULATOR_HOST=\"localhost:8080\"")
		//if err = cmd.Run(); err != nil {
		//	panic(err)
		//}
		err = os.Setenv("FIRESTORE_EMULATOR_HOST", "localhost:8080")
		if err != nil {
			panic(err)
		}
		log.Println("emulator host:", os.Getenv("FIRESTORE_EMULATOR_HOST"))
	}
	ctx := context.Background()
	if FirebaseApp, err = NewFirebase(ctx); err != nil {
		panic(err)
	}
	if FirestoreCli, err = NewFirestore(ctx, FirebaseApp); err != nil {
		panic(err)
	}
}

func NewFirebase(ctx context.Context) (app *firebase.App, err error) {
	log.Println("new firestore:", PROJECT_ID)
	app, err = firebase.NewApp(ctx, &firebase.Config{
		ProjectID: PROJECT_ID,
	})
	if err != nil {
		err = errors.New("failed to initialize firebase app")
		return
	}
	return
}

func NewFirestore(ctx context.Context, firebaseApp *firebase.App) (app *firestore.Client, err error) {
	if firebaseApp == nil {
		err = errors.New("need firebase app to initialize firestore client")
		return
	}
	app, err = firebaseApp.Firestore(ctx)
	if err != nil {
		err = errors.New("failed to initialize firestore client")
		return
	}
	return
}
