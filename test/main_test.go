package test

import (
	"cloudfunc"
	"context"
	yaml "gopkg.in/yaml.v2"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestMain(t *testing.M) {
	SetupSample()
}

func SetupSample() {
	var err error
	_, file, _, _ := runtime.Caller(0)
	filePath := filepath.Dir(file) + "/samples"
	samplePath := filePath + "/example.yml"
	data, err := os.ReadFile(samplePath)
	if err != nil {
		panic(err)
	}
	var samples []Sample
	err = yaml.Unmarshal(data, &samples)
	if err != nil {
		log.Println(data)
		panic(err)
	}

	ctx := context.Background()
	for _, s := range samples {
		log.Println(s)
		_, err = cloudfunc.FirestoreCli.Collection("MyTesting").Doc("Author").Create(ctx, map[string]interface{}{
			"Name": "123",
		})
		if err != nil {
			panic(err)
		}
	}
}

type Sample struct {
	Path   string                 `yaml:"path"`
	Fields map[string]interface{} `yaml:"fields"`
}
