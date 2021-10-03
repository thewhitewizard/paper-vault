package util

import (
	"bytes"
	"fmt"
	"image"
	"io/ioutil"

	"github.com/liyue201/goqr"
	qrcode "github.com/yeqown/go-qrcode"
)

func GenerateQR(fileName string, data string) error {

	qrc, err := qrcode.New(data)
	if err != nil {
		return err
	}
	return qrc.Save(fileName + ".jpeg")
}

func ReadDataFromQRFile(fileName string) (string, error) {
	imgdata, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	img, _, err := image.Decode(bytes.NewReader(imgdata))
	if err != nil {
		return "", err
	}
	qrCodes, err := goqr.Recognize(img)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", qrCodes[0].Payload), nil

}
