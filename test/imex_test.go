package test

import (
	"fmt"
	"io"
	"testing"

	"github.com/ianyulistios/imex"
	"gopkg.in/go-playground/assert.v1"
)

var (
	dummyURL   = "https://images.unsplash.com/photo-1496181133206-80ce9b88a853?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=751&q=80"
	dummiesURL = []string{
		"https://images.unsplash.com/photo-1496181133206-80ce9b88a853?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=751&q=80",
		"https://error_url",
		"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSYscfUBUbqwGd_DHVhG-ZjCOD7MUpxp4uhNe7toUg4ug&s",
		"https://img.freepik.com/free-photo/autumn-leaf-falling-revealing-intricate-leaf-vein-generated-by-ai_188544-9869.jpg",
	}
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
	for _, url := range dummiesURL {
		instance := imex.InitImax(url)
		response := instance.DownloadFile()
		if response.ErrorData != nil {
			fmt.Println(response.ErrorData.Error())
			assert.Equal(t, response.RawFile, typeResponse)
		} else {
			assert.NotEqual(t, response.RawFile, typeResponse)
		}

	}
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

func TestDownloadImageCustomMime(t *testing.T) {
	instance := imex.InitImax(dummyURL)

	image, err := instance.DownloadFile().ToImage("jpeg")

	if err != nil {
		t.Errorf(err.Error())
	}

	if image == "" {
		t.Errorf("Image is not downloaded successfully")
	}
}

func TestDownloadAsByte(t *testing.T) {
	instance := imex.InitImax(dummyURL)

	image, mimeType, err := instance.DownloadFile().ToByte()

	if err != nil {
		t.Errorf(err.Error())
	}

	if mimeType == "" {
		t.Errorf("Image is not downloaded successfully")
	}

	if image == nil {
		t.Errorf("Image is not downloaded successfully")
	}
}

func TestDownloadToByteAndClose(t *testing.T) {
	for _, url := range dummiesURL {
		instance := imex.InitImax(url)
		image, mimeType, err := instance.DownloadFile().ToByte()

		if err == nil {
			assert.NotEqual(t, image, nil)
			assert.NotEqual(t, mimeType, "")
		} else {
			assert.Equal(t, image, nil)
			assert.Equal(t, mimeType, "")
		}

		instance.Close()
	}
}

func TestFailonClosedFile(t *testing.T) {
	for _, url := range dummiesURL {
		instance := imex.InitImax(url)
		instance.DownloadFile().Close()
		_, mimeType, err := instance.ToByte()

		assert.NotEqual(t, err, nil)
		assert.Equal(t, mimeType, "")
	}
}
