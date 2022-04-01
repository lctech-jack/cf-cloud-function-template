package test

import (
	"cloudfunc"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
)

type ExampleSuite struct {
	suite.Suite
}

func (t *ExampleSuite) SetupTest() {

}

func TestExampleSuiteInit(t *testing.T) {
	suite.Run(t, new(ExampleSuite))
}

func (t *ExampleSuite) TestExample() {
	log.Println("Hello World!")
	doc, err := cloudfunc.FirestoreCli.Collection("Example").Doc("Author").Get(context.Background())
	assert.NoError(t.T(), err)
	log.Println(doc.Data())
}
