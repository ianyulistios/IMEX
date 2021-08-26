package test

import (
	"fmt"
	"io"
	"testing"

	"github.com/ianyulistios/imex"
	"gopkg.in/go-playground/assert.v1"
)

var (
	dummyURL = "https://images.unsplash.com/photo-1496181133206-80ce9b88a853?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=751&q=80"
)

func TestInit(t *testing.T) {

	initImex := imex.InitImax(dummyURL)

	data := initImex.GetData()

	assert.Equal(t, data.FileURL, dummyURL)
}

func TestDownloadFile(t *testing.T) {
	var (
		typeResponse io.ReadCloser
	)

	instance := imex.InitImax(dummyURL)
	response := instance.DownloadFile()
	fmt.Print(response.RawFile)
	assert.NotEqual(t, response.RawFile, typeResponse)
}

func TestDownloadImage(t *testing.T) {
	instance := imex.InitImax(dummyURL)

	image, err := instance.DownloadFile().ToImage()

	if err != nil {
		t.Errorf(err.Error())
	}

	if image == "" {
		t.Errorf("Image is not downloaded successfully")
	}
}
