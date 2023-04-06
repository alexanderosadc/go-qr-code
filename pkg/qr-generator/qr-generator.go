package qrgenerator

import (
	"errors"
	"image"
	"image/png"
	"io"
	"log"
)

func CreateQRCode(w io.Writer, data string, version Version) error {
	if len(data) == 0 {
		return errors.New("data parameter has size 0")
	}
	size := version.CalcSizeByVersion()

	img := image.NewNRGBA(image.Rect(0, 0, size, size))
	err := png.Encode(w, img)
	if err != nil {
		log.Printf("data %s is not encoded: %s", data, err)
		return err
	}

	return nil
}

type Version int8

func (v Version) CalcSizeByVersion() int {
	return 4*int(v) + 17
}
