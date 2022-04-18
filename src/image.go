package src

import (
	"encoding/base64"
	"io"
	"io/ioutil"
	"net/http"
)

func ConvertToByte() {

}

func ToBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func ReadResponse(resp io.ReadCloser) ([]byte, string, error) {
	var (
		dataByte []byte
		mimeType string
		errData  error
	)

	dataByte, errData = ioutil.ReadAll(resp)

	if errData != nil {
		return dataByte, mimeType, errData
	}

	mimeType = http.DetectContentType(dataByte)

	return dataByte, mimeType, errData
}
